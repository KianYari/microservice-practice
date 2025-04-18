package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	pb "github.com/kianyari/microservice-practice/common/api"
	"github.com/kianyari/microservice-practice/task-service/internal/dto"
	"github.com/kianyari/microservice-practice/task-service/internal/model"
	repository "github.com/kianyari/microservice-practice/task-service/internal/repository"
	"github.com/streadway/amqp"
)

type TaskServiceInterface interface {
	CreateTask(createTaskRequest dto.CreateTaskRequest) error
	GetTasks(ownerID uint) (dto.TaskList, error)
	CompleteTask(completeTaskRequest dto.CompleteTaskRequest) error
	DeleteTask(deleteTaskRequest dto.DeleteTaskRequest) error
	StartDeadlineChecker()
}

type TaskService struct {
	userClient     pb.UserServiceClient
	taskRepository repository.TaskRepository
	rabbitMQConn   *amqp.Connection
}

func NewTaskService(
	userClient pb.UserServiceClient,
	taskRepository repository.TaskRepository,
	rabbitMQConn *amqp.Connection,
) *TaskService {
	return &TaskService{
		userClient:     userClient,
		taskRepository: taskRepository,
		rabbitMQConn:   rabbitMQConn,
	}
}

func (taskService *TaskService) CreateTask(createTaskRequest dto.CreateTaskRequest) error {
	getUserByIdRequest := &pb.GetUserByIdRequest{
		Id: int32(createTaskRequest.OwnerID),
	}
	ctx := context.Background()
	_, err := taskService.userClient.GetUserByID(ctx, getUserByIdRequest)
	if err != nil {
		return errors.New("user not found")
	}

	err = taskService.taskRepository.CreateTask(createTaskRequest)

	if err != nil {
		return errors.New("failed to create task")
	}

	taskService.publishTaskCreatedNotification(createTaskRequest)

	return nil
}

func (taskService *TaskService) GetTasks(ownerID uint) (dto.TaskList, error) {
	ctx := context.Background()
	getUserByIdRequest := &pb.GetUserByIdRequest{
		Id: int32(ownerID),
	}
	_, err := taskService.userClient.GetUserByID(ctx, getUserByIdRequest)
	if err != nil {
		return dto.TaskList{}, errors.New("user not found")
	}

	tasks, err := taskService.taskRepository.GetTasks(ownerID)
	var taskList dto.TaskList
	for _, task := range tasks {
		taskList.Tasks = append(taskList.Tasks, dto.Task{
			ID:       task.ID,
			OwnerID:  task.OwnerID,
			Title:    task.Title,
			Deadline: task.Deadline,
			Status:   task.Status,
		})
	}
	if err != nil {
		return dto.TaskList{}, errors.New("failed to get tasks")
	}
	return taskList, nil
}

func (taskService *TaskService) CompleteTask(completeTaskRequest dto.CompleteTaskRequest) error {
	user, err := taskService.userClient.GetUserByID(context.Background(), &pb.GetUserByIdRequest{
		Id: int32(completeTaskRequest.OwnerID),
	})
	if err != nil {
		return errors.New("user not found")
	}
	task, exist := taskService.taskRepository.GetTaskByID(completeTaskRequest.TaskID)
	if exist != nil {
		return errors.New("task not found")
	}
	if task.OwnerID != uint(user.Id) {
		return errors.New("task does not belong to user")
	}
	task.Status = "completed"
	err = taskService.taskRepository.UpdateTask(task)
	if err != nil {
		return errors.New("failed to update task")
	}
	return nil
}

func (taskService *TaskService) DeleteTask(deleteTaskRequest dto.DeleteTaskRequest) error {
	user, err := taskService.userClient.GetUserByID(context.Background(), &pb.GetUserByIdRequest{
		Id: int32(deleteTaskRequest.OwnerID),
	})
	if err != nil {
		return errors.New("user not found")
	}
	task, exist := taskService.taskRepository.GetTaskByID(deleteTaskRequest.TaskID)
	if exist != nil {
		return errors.New("task not found")
	}
	if task.OwnerID != uint(user.Id) {
		return errors.New("task does not belong to user")
	}
	err = taskService.taskRepository.DeleteTask(task.ID)
	if err != nil {
		return errors.New("failed to delete task")
	}
	return nil
}

func (taskService *TaskService) StartDeadlineChecker() {
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		taskService.CheckDadlines()

		for range ticker.C {
			taskService.CheckDadlines()
		}
	}()
}

func (taskService *TaskService) CheckDadlines() {
	tasks, err := taskService.taskRepository.GetTodayTasks()
	if err != nil {
		return
	}

	for _, task := range tasks {
		taskService.publishDeadlineNotification(task)
	}
}

func (taskService *TaskService) publishDeadlineNotification(task *model.Task) {
	ch, err := taskService.rabbitMQConn.Channel()
	if err != nil {
		log.Println("Failed to open RabbitMQ channel:", err)
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"task_deadline",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Failed to declare exchange:", err)
		return
	}

	notificationPayload := struct {
		TaskID   uint      `json:"task_id"`
		OwnerID  uint      `json:"owner_id"`
		Title    string    `json:"title"`
		Deadline time.Time `json:"deadline"`
	}{
		TaskID:   task.ID,
		OwnerID:  task.OwnerID,
		Title:    task.Title,
		Deadline: task.Deadline,
	}

	payload, err := json.Marshal(notificationPayload)
	if err != nil {
		log.Println("Failed to marshal notification payload:", err)
		return
	}

	err = ch.Publish(
		"task_deadline",
		"task_deadline",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payload,
		},
	)
	if err != nil {
		log.Println("Failed to publish message:", err)
		return
	}
	log.Printf("Published deadline notification for task ID %d to RabbitMQ\n", task.ID)
}

func (taskService *TaskService) publishTaskCreatedNotification(createTaskRequest dto.CreateTaskRequest) {
	ch, err := taskService.rabbitMQConn.Channel()
	if err != nil {
		log.Println("Failed to open RabbitMQ channel:", err)
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"task_created",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Failed to declare exchange:", err)
		return
	}
	notificationPayload := struct {
		OwnerID  uint      `json:"owner_id"`
		Title    string    `json:"title"`
		Deadline time.Time `json:"deadline"`
	}{
		OwnerID:  createTaskRequest.OwnerID,
		Title:    createTaskRequest.Title,
		Deadline: createTaskRequest.Deadline,
	}

	payload, err := json.Marshal(notificationPayload)
	if err != nil {
		log.Println("Failed to marshal notification payload:", err)
		return
	}
	err = ch.Publish(
		"task_created",
		"task_created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payload,
		},
	)
	if err != nil {
		log.Println("Failed to publish message:", err)
		return
	}
	log.Printf("Published task created notification for task %s to RabbitMQ\n", createTaskRequest.Title)
}
