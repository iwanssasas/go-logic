package handler

import (
	"bytes"
	"context"
	"encoding/base64"
	"go-logic/003.dependencyInjection/entity"
	"go-logic/003.dependencyInjection/service"
	"go-logic/003.dependencyInjection/utils"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler interface {
	TestPing(c *gin.Context)
	AddStudents(c *gin.Context)
	GetAllStudents(c *gin.Context)
	DeleteStudents(c *gin.Context)
	UpdateStudents(c *gin.Context)

	UploadExcel(c *gin.Context)
	GetUploadExcel(c *gin.Context)
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
	ctx := context.Background()
	ID, err := uuid.Parse(IdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	ID, err := uuid.Parse(IdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

func (h handler) UploadExcel(c *gin.Context) {
	var params entity.UploadExcelModel
	ctx := context.Background()

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	file, err := base64.StdEncoding.DecodeString(params.Excel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	r := bytes.NewReader(file)
	e, err := excelize.OpenReader(r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	dataExcels := e.GetRows("Sheet1")
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	dataExcels = dataExcels[1:]
	if len(dataExcels) < 1 {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	go func(ctx context.Context, dataExcels [][]string) {
		err := h.service.UploadExcelService(ctx, dataExcels)
		if err != nil {
			panic(err)
		}

	}(ctx, dataExcels)
	c.JSON(http.StatusOK, gin.H{
		"data":  "SUCCES",
		"error": nil,
	})
}

func (h handler) GetUploadExcel(c *gin.Context) {
	ctx := context.Background()
	result, err := h.service.GetAllUploadExcel(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  result,
		"error": nil,
	})
}
