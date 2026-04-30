package infra

import (
	"fmt"
	"task-manager/internal/domain"
	"task-manager/internal/infra/storage"
)

type Format string

const (
	FormatCSV  Format = "csv"
	FormatJSON Format = "json"
)

func NewTaskRepository(format Format) (domain.TaskRepository, error) {
	switch format {
	case FormatCSV:
		return storage.NewCSV("tasks.csv"), nil
	case FormatJSON:
		return storage.NewJSON("tasks.json"), nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
}
