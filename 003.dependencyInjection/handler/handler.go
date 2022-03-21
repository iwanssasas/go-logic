package handler

import (
	"context"
	"go-logic/003.dependencyInjection/entity"
	"go-logic/003.dependencyInjection/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler interface {
	TestPing(c *gin.Context)
	AddStudents(c *gin.Context)
	GetAllStudents(c *gin.Context)
	DeleteStudents(c *gin.Context)
	UpdateStudents(c *gin.Context)
}

type handler struct {
	service service.Service
}

func NewHandler() Handler {
	return handler{
		service: service.NewService(),
	}
}

func (h handler) TestPing(c *gin.Context) {
	ctx := context.Background()
	result, err := h.service.TestPingService(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": result})
}

func (h handler) AddStudents(c *gin.Context) {
	var params entity.Student

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	result, err := h.service.AddStudentService(ctx, params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": result})
}

func (h handler) GetAllStudents(c *gin.Context) {
	ctx := context.Background()
	result, err := h.service.GetStudentService(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": result})
}

func (h handler) DeleteStudents(c *gin.Context) {
	IdString := c.Param("id")
	ID := uuid.MustParse(IdString)
	ctx := context.Background()
	result, err := h.service.DeleteStudents(ctx, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": result})
}

func (h handler) UpdateStudents(c *gin.Context) {
	var params entity.Student
	ctx := context.Background()
	IdString := c.Param("id")

	ID := uuid.MustParse(IdString)

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.service.UpdateStudentService(ctx, params, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": result})
}
