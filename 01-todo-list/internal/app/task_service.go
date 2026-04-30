package app

import (
	"errors"
	"fmt"
	"time"

	"task-manager/internal/domain"
)

type TaskService struct {
	repo  domain.TaskRepository
	tasks domain.Tasks
}

func NewTaskService(s domain.TaskRepository) *TaskService {
	return &TaskService{repo: s}
}

func (s *TaskService) Load() { s.repo.Load(&s.tasks) }
func (s *TaskService) save() { s.repo.Save(s.tasks) }

func (s *TaskService) validateIndex(i int) error {
	if i < 0 || i >= len(s.tasks) {
		err := errors.New("invalid task index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *TaskService) Add(title, desc string) {
	s.tasks = append(s.tasks, domain.Task{
		ID:          len(s.tasks) + 1,
		Title:       title,
		Description: desc,
		CreatedAt:   time.Now(),
	})
	s.repo.Save(s.tasks)
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
		s.repo.Save(s.tasks)
		return nil
	}
	t.IsCompleted = false
	s.repo.Save(s.tasks)
	return nil
}

func (s *TaskService) Remove(i int) error {
	if err := s.validateIndex(i); err != nil {
		return err
	}
	s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
	s.repo.Save(s.tasks)
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
	s.repo.Save(s.tasks)
	return nil
}

func (s *TaskService) List() domain.Tasks {
	return s.tasks
}
