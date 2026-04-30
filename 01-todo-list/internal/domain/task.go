package domain

import (
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	IsCompleted bool
	CompletedAt *time.Time
}

type Tasks []Task

type TaskRepository interface {
	Load(*Tasks) error
	Save(Tasks) error
}
