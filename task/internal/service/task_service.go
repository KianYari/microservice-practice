package service

import (
	"context"
	"errors"

	pb "github.com/kianyari/microservice-practice/common/api"
	"github.com/kianyari/microservice-practice/task-service/internal/dto"
	repository "github.com/kianyari/microservice-practice/task-service/internal/repository"
)

type TaskService interface {
	CreateTask(createTaskRequest dto.CreateTaskRequest) error
	GetTasks(ownerID uint) (dto.TaskList, error)
	CompleteTask(completeTaskRequest dto.CompleteTaskRequest) error
}

type taskService struct {
	userClient     pb.UserServiceClient
	taskRepository repository.TaskRepository
}

func NewTaskService(
	userClient pb.UserServiceClient,
	taskRepository repository.TaskRepository,
) *taskService {
	return &taskService{
		userClient:     userClient,
		taskRepository: taskRepository,
	}
}

func (taskService *taskService) CreateTask(createTaskRequest dto.CreateTaskRequest) error {
	getUserByIdRequest := &pb.GetUserByIdRequest{
		Id: int32(createTaskRequest.OwnerID),
	}
	ctx := context.Background()
	_, exist := taskService.userClient.GetUserByID(ctx, getUserByIdRequest)
	if exist != nil {
		return errors.New("user not found")
	}

	err := taskService.taskRepository.CreateTask(createTaskRequest)

	if err != nil {
		return errors.New("failed to create task")
	}
	return nil
}

func (taskService *taskService) GetTasks(ownerID uint) (dto.TaskList, error) {
	ctx := context.Background()
	getUserByIdRequest := &pb.GetUserByIdRequest{
		Id: int32(ownerID),
	}
	_, exist := taskService.userClient.GetUserByID(ctx, getUserByIdRequest)
	if exist != nil {
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

func (taskService *taskService) CompleteTask(completeTaskRequest dto.CompleteTaskRequest) error {
	user, exist := taskService.userClient.GetUserByID(context.Background(), &pb.GetUserByIdRequest{
		Id: int32(completeTaskRequest.OwnerID),
	})
	if exist != nil {
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
	err := taskService.taskRepository.UpdateTask(task)
	if err != nil {
		return errors.New("failed to update task")
	}
	return nil
}
