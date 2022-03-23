package service

import (
	"context"
	"errors"
	"fmt"
	"go-logic/003.dependencyInjection/entity"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	TestPingService(ctx context.Context) (string, error)
	AddStudentService(ctx context.Context, params entity.Student) (*uuid.UUID, error)
	GetStudentService(ctx context.Context) (entity.Students, error)
	DeleteStudents(ctx context.Context, ID uuid.UUID) (*string, error)
	UpdateStudentService(ctx context.Context, params entity.Student, ID uuid.UUID) (*uuid.UUID, error)

	UploadExcelService(ctx context.Context, dataExcel [][]string) error
	GetAllUploadExcel(ctx context.Context) (entity.UploadExcels, error)
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

func (s service) UploadExcelService(ctx context.Context, dataExcel [][]string) error {

	wg := &sync.WaitGroup{}
	uploadExcelChan := make(chan entity.UploadExcelDatabase)

	for ir, row := range dataExcel {
		wg.Add(1)

		go s.upsertExcel(wg, ir, row, uploadExcelChan)
	}

	go func(wg *sync.WaitGroup, uploadExcelChan chan entity.UploadExcelDatabase) {
		wg.Wait()
		close(uploadExcelChan)
	}(wg, uploadExcelChan)

	mapUploadExcel := make(entity.MapUploadExcel)

	for val := range uploadExcelChan {

		keyMap := fmt.Sprintf("%v", val.ID)
		mapUploadExcel[keyMap] = val
	}

	database.DataExcel = mapUploadExcel

	return nil

}

func (s service) upsertExcel(wg *sync.WaitGroup, ir int, row []string, uploadExcelChan chan entity.UploadExcelDatabase) {
	defer wg.Done()
	var err error
	var uploadExcelDatabase entity.UploadExcelDatabase

	isImportant := func(colPosition int, col string) {
		if col == "" {
			errors.New("ROW KOSONG")
		}
	}

	for ic, col := range row {
		switch ic {
		case 0:
			isImportant(ic, col)
			idx, err := uuid.Parse(col)
			if err != nil {
				panic(err)
			}

			uploadExcelDatabase.ID = idx
		case 1:
			isImportant(ic, col)
			uploadExcelDatabase.Name = strings.ToUpper(col)
		case 2:
			isImportant(ic, col)
			uploadExcelDatabase.IsPublish, err = strconv.ParseBool(col)
			if err != nil {
				panic(err)
			}
		case 3:
			isImportant(ic, col)
			dateInt, err := strconv.Atoi(col)
			if err != nil {
				panic(err)
			}
			d := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
			d2 := d.AddDate(0, 0, dateInt)
			uploadExcelDatabase.PublishDate = d2
		case 4:
			isImportant(ic, col)
			uploadExcelDatabase.Notes = strings.TrimSpace(col)
		case 5:
			uploadExcelDatabase.Quantity, err = strconv.Atoi(col)
			if err != nil {
				panic(err)
			}

		}
	}

	uploadExcelChan <- uploadExcelDatabase

}

func (s service) GetAllUploadExcel(ctx context.Context) (entity.UploadExcels, error) {
	var result []entity.UploadExcelDatabase

	for _, val := range database.DataExcel {
		result = append(result, val)
	}
	return result, nil
}
