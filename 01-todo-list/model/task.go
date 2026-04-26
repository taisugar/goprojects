package model

import (
	"time"
)

type Task struct {
	Title       string
	Description string
	CreatedAt   time.Time
	IsCompleted bool
	CompletedAt *time.Time
}

type Tasks []Task

// func (tasks *Tasks) add(title string, description string) {
// 	task := Task{
// 		Title:       title,
// 		Description: description,
// 		CreatedAt:   time.Now(),
// 	}
// 	*tasks = append(*tasks, task)
// }

// func (tasks *Tasks) validateIndex(index int) error {
// 	if index < 0 || index >= len(*tasks) {
// 		err := errors.New("invalid task index")
// 		fmt.Println(err)
// 		return err
// 	}

// 	return nil
// }

// func (tasks *Tasks) delete(index int) error {
// 	if err := tasks.validateIndex(index); err != nil {
// 		return err
// 	}
// 	*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
// 	return nil
// }

// func (tasks *Tasks) complete(index int) error {
// 	if err := tasks.validateIndex(index); err != nil {
// 		return err
// 	}
// 	t := (*tasks)
// 	IsCompleted := t[index].IsCompleted
// 	if !IsCompleted {
// 		now := time.Now()
// 		t[index].CompletedAt = &now
// 	}
// 	t[index].IsCompleted = !IsCompleted

// 	return nil
// }

// func (tasks *Tasks) edit(index int, title string, description string) error {
// 	if err := tasks.validateIndex(index); err != nil {
// 		return err
// 	}
// 	t := (*tasks)
// 	t[index].Title = title
// 	t[index].Description = description
// 	return nil
// }

// func (tasks *Tasks) print() {
// 	table := table.New(os.Stdout)
// 	table.SetRowLines(false)
// 	table.SetHeaders("#", "Title", "Description", "Created At", "Completed", "Completed At")
// 	for i, task := range *tasks {
// 		completed := "❌"
// 		completedAt := ""
// 		if task.IsCompleted {
// 			completed = "✅"
// 		}
// 		if task.CompletedAt != nil {
// 			completedAt = timediff.TimeDiff(*task.CompletedAt)
// 		}
// 		table.AddRow(strconv.Itoa(i), task.Title, task.Description, task.CreatedAt.Format(time.RFC1123), completed, completedAt)
// 	}
// 	table.Render()
// }
