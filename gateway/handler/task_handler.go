package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/kianyari/microservice-practice/common/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskHandler struct {
	taskClient pb.TaskServiceClient
	jwtService pb.JWTServiceClient
}

func NewTaskHandler(
	taskClient pb.TaskServiceClient,
	jwtService pb.JWTServiceClient,
) *TaskHandler {
	return &TaskHandler{
		taskClient: taskClient,
		jwtService: jwtService,
	}
}

func (h *TaskHandler) RegisterRoutes(ginEngine *gin.Engine) {
	taskGroup := ginEngine.Group("/tasks")
	{
		taskGroup.POST("/create", h.CreateTask)
		taskGroup.GET("/get-list", h.GetTasks)
		taskGroup.POST("/complete", h.CompleteTask)
		taskGroup.DELETE("/delete", h.DeleteTask)

	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req struct {
		Title    string `json:"title"`
		Deadline string `json:"deadline"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var deadlineTimestamp *timestamppb.Timestamp
	if req.Deadline != "" {
		parsedTime, err := time.Parse(time.RFC3339, req.Deadline)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid deadline format. Use RFC3339 format (e.g., 2025-03-28T15:04:05Z)."})
			return
		}
		deadlineTimestamp = timestamppb.New(parsedTime)
	}

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "Authorization header is required"})
		return
	}
	token := authHeader[len("Bearer "):]
	validateTokenRequest := &pb.ValidateTokenRequest{
		Token: token,
	}
	response, err := h.jwtService.ValidateToken(c, validateTokenRequest)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	ownerID := response.Id

	protoReq := &pb.CreateTaskRequest{
		OwnerId:  ownerID,
		Title:    req.Title,
		Deadline: deadlineTimestamp,
	}

	res, err := h.taskClient.CreateTask(c, protoReq)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	var req pb.GetTasksRequest

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "Authorization header is required"})
		return
	}
	token := authHeader[len("Bearer "):]
	validateTokenRequest := &pb.ValidateTokenRequest{
		Token: token,
	}
	response, err := h.jwtService.ValidateToken(c, validateTokenRequest)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	ownerID := response.Id
	req.OwnerId = ownerID

	res, err := h.taskClient.GetTasks(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	type TaskResponse struct {
		ID       int32  `json:"id"`
		Title    string `json:"title"`
		Deadline string `json:"deadline"`
		Status   string `json:"status"`
	}

	var tasks []TaskResponse
	for _, task := range res.Tasks {
		tasks = append(tasks, TaskResponse{
			ID:       task.Id,
			Title:    task.Title,
			Deadline: time.Unix(task.Deadline.Seconds, 0).Format(time.RFC3339),
			Status:   task.Status,
		})
	}

	c.JSON(200, tasks)
}

func (h *TaskHandler) CompleteTask(c *gin.Context) {
	var req struct {
		TaskId int32 `json:"task_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "Authorization header is required"})
		return
	}
	token := authHeader[len("Bearer "):]
	validateTokenRequest := &pb.ValidateTokenRequest{
		Token: token,
	}
	response, err := h.jwtService.ValidateToken(c, validateTokenRequest)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	ownerID := response.Id

	protoReq := &pb.CompleteTaskRequest{
		Id:      req.TaskId,
		OwnerId: ownerID,
	}

	res, err := h.taskClient.CompleteTask(c, protoReq)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	var req struct {
		TaskID int32 `json:"task_id" validate:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "Authorization header is required"})
		return
	}
	token := authHeader[len("Bearer "):]
	validateTokenRequest := &pb.ValidateTokenRequest{
		Token: token,
	}
	response, err := h.jwtService.ValidateToken(c, validateTokenRequest)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	ownerID := response.Id
	protoReq := &pb.DeleteTaskRequest{
		Id:      req.TaskID,
		OwnerId: ownerID,
	}
	res, err := h.taskClient.DeleteTask(c, protoReq)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}
