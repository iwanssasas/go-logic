package service

import (
	"context"
	"errors"
	"fmt"
	"go-logic/003.dependencyInjection/entity"

	"github.com/google/uuid"
)

type Service interface {
	TestPingService(ctx context.Context) (string, error)
	AddStudentService(ctx context.Context, params entity.Student) (*uuid.UUID, error)
	GetStudentService(ctx context.Context) (entity.Students, error)
	DeleteStudents(ctx context.Context, ID uuid.UUID) (*string, error)
	UpdateStudentService(ctx context.Context, params entity.Student, ID uuid.UUID) (*uuid.UUID, error)
}

type service struct {
}

var database entity.DatabaseModel

func NewService() Service {
	database = entity.DatabaseModel{
		Data: make(map[uuid.UUID]entity.Student),
		Temp: make(map[string]bool),
	}
	return service{}
}

func (s service) TestPingService(ctx context.Context) (string, error) {
	status := "Succes"
	return status, nil
}

func (s service) AddStudentService(ctx context.Context, params entity.Student) (*uuid.UUID, error) {

	_, hasName := database.Temp[params.Name]
	if hasName {
		return nil, errors.New("NAME IS EXSIST")
	}

	id := uuid.New()
	database.Data[id] = params
	database.Temp[params.Name] = true

	return &id, nil
}

func (s service) GetStudentService(ctx context.Context) (entity.Students, error) {

	var result []entity.Student

	for _, val := range database.Data {
		result = append(result, val)
	}
	fmt.Println(database.Data)
	return result, nil

}

func (s service) DeleteStudents(ctx context.Context, ID uuid.UUID) (*string, error) {
	_, hasID := database.Data[ID]
	if !hasID {
		return nil, errors.New("ID IS NOT FOUND")
	}

	delete(database.Data, ID)

	response := "success"
	return &response, nil
}

func (s service) UpdateStudentService(ctx context.Context, params entity.Student, ID uuid.UUID) (*uuid.UUID, error) {
	_, hasName := database.Temp[params.Name]
	if hasName {
		return nil, errors.New("NAME IS EXSIST")
	}

	database.Data[ID] = params
	database.Temp[params.Name] = true

	return &ID, nil
}
