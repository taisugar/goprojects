package service

import (
	"errors"
	"fmt"
	"time"

	"task-manager/model"
	"task-manager/storage"
)

type TaskService struct {
	store storage.Strategy[[]model.Task]
	tasks []model.Task
}

func NewTaskService(s storage.Strategy[[]model.Task]) *TaskService {
	return &TaskService{store: s}
}

func (s *TaskService) Load() { s.store.Load(&s.tasks) }
func (s *TaskService) Save() { s.store.Save(s.tasks) }

func (s *TaskService) validateIndex(i int) error {
	if i < 0 || i >= len(s.tasks) {
		err := errors.New("invalid task index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *TaskService) Add(title, desc string) {
	s.tasks = append(s.tasks, model.Task{
		Title:       title,
		Description: desc,
		CreatedAt:   time.Now(),
	})
}

func (s *TaskService) Complete(i int) error {
	if err := s.validateIndex(i); err != nil {
		return err
	}
	t := &s.tasks[i]

	if !t.IsCompleted {
		now := time.Now()
		t.CompletedAt = &now
		t.IsCompleted = true
		return nil
	}
	t.IsCompleted = false
	return nil
}

func (s *TaskService) Remove(i int) error {
	if err := s.validateIndex(i); err != nil {
		return err
	}
	s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
	return nil
}

func (s *TaskService) Edit(i int, title, desc string) error {
	if err := s.validateIndex(i); err != nil {
		return err
	}
	if title != "" {
		s.tasks[i].Title = title
	}
	if desc != "" {
		s.tasks[i].Description = desc
	}
	return nil
}

func (s *TaskService) List() []model.Task {
	return s.tasks
}
