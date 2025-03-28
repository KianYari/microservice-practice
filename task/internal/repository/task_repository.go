package repository

import (
	"github.com/kianyari/microservice-practice/task-service/internal/dto"
	"github.com/kianyari/microservice-practice/task-service/internal/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(createTaskRequest dto.CreateTaskRequest) error
	GetTasks(ownerID uint) ([]*model.Task, error)
	UpdateTask(task *model.Task) error
	GetTaskByID(taskID uint) (*model.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(
	db *gorm.DB,
) *taskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) CreateTask(createTaskRequest dto.CreateTaskRequest) error {
	task := model.Task{
		OwnerID:  createTaskRequest.OwnerID,
		Title:    createTaskRequest.Title,
		Deadline: createTaskRequest.Deadline,
		Status:   createTaskRequest.Status,
	}
	if err := r.db.Create(&task).Error; err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) GetTasks(ownerID uint) ([]*model.Task, error) {
	var tasks []*model.Task
	if err := r.db.Where("owner_id = ?", ownerID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) UpdateTask(task *model.Task) error {
	if err := r.db.Save(task).Error; err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) GetTaskByID(taskID uint) (*model.Task, error) {
	var task model.Task
	if err := r.db.Where("id = ?", taskID).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
