package entity

import "github.com/google/uuid"

type (
	DatabaseModel struct {
		Data map[uuid.UUID]Student
		Temp map[string]bool
	}
)
