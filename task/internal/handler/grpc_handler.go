package handler

import (
	"context"

	pb "github.com/kianyari/microservice-practice/common/api"
	"github.com/kianyari/microservice-practice/task-service/internal/dto"
	service "github.com/kianyari/microservice-practice/task-service/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type grpcHandler struct {
	pb.UnimplementedTaskServiceServer
	taskService service.TaskService
}

func NewGRPCHandler(
	grpcServer *grpc.Server,
	taskService service.TaskService,
) {
	handler := &grpcHandler{
		taskService: taskService,
	}
	pb.RegisterTaskServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	createTaskRequest := dto.CreateTaskRequest{
		OwnerID:  uint(req.OwnerId),
		Title:    req.Title,
		Deadline: req.Deadline.AsTime(),
		Status:   "pending",
	}

	err := h.taskService.CreateTask(createTaskRequest)
	if err != nil {
		return &pb.CreateTaskResponse{Message: err.Error()}, err
	}
	return &pb.CreateTaskResponse{Message: "task created successfully"}, nil
}

func (h *grpcHandler) GetTasks(ctx context.Context, req *pb.GetTasksRequest) (*pb.GetTasksResponse, error) {
	tasks, err := h.taskService.GetTasks(uint(req.OwnerId))
	if err != nil {
		return &pb.GetTasksResponse{Message: err.Error()}, err
	}

	var tasksResponse []*pb.Task
	for _, task := range tasks.Tasks {
		tasksResponse = append(tasksResponse, &pb.Task{
			Id:       int32(task.ID),
			Title:    task.Title,
			Deadline: timestamppb.New(task.Deadline),
			Status:   task.Status,
		})
	}

	return &pb.GetTasksResponse{Tasks: tasksResponse}, nil
}

func (h *grpcHandler) CompleteTask(ctx context.Context, req *pb.CompleteTaskRequest) (*pb.CompleteTaskResponse, error) {
	completeTaskRequest := dto.CompleteTaskRequest{
		TaskID:  uint(req.Id),
		OwnerID: uint(req.OwnerId),
	}

	err := h.taskService.CompleteTask(completeTaskRequest)
	if err != nil {
		return &pb.CompleteTaskResponse{Message: err.Error()}, err
	}
	return &pb.CompleteTaskResponse{Message: "task completed successfully"}, nil
}

func (h *grpcHandler) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	deleteTaskRequest := dto.DeleteTaskRequest{
		TaskID:  uint(req.Id),
		OwnerID: uint(req.OwnerId),
	}

	err := h.taskService.DeleteTask(deleteTaskRequest)
	if err != nil {
		return &pb.DeleteTaskResponse{Message: err.Error()}, err
	}
	return &pb.DeleteTaskResponse{Message: "task deleted successfully"}, nil
}
