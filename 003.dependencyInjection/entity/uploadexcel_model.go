package entity

import (
	"time"

	"github.com/google/uuid"
)

type (
	UploadExcelModel struct {
		Excel string `json:"excel"`
	}

	UploadExcelDatabase struct {
		ID          uuid.UUID `json:"Id"`
		Name        string    `json:"name"`
		IsPublish   bool      `json:"is_publish"`
		PublishDate time.Time `json:"publish_date"`
		Notes       string    `json:"notes"`
		Quantity    int       `json:"quantity"`
	}

	UploadExcels []UploadExcelDatabase

	MapUploadExcel map[string]UploadExcelDatabase

	ErrorsExcel struct {
		Row          int    `json:"row"`
		ErrorMessage string `json:"error_message"`
	}

	ErrorsExcels []ErrorsExcel
	Response     struct {
		ErrorMessage ErrorsExcels `json:"error_message"`
		DataExcel    UploadExcels `json:"data_excel"`
	}
)

type StringSliceParams [][]string
