package adapter_mcp

import (
	"context"
	"fmt"
	"strings"
	"task-manager/internal/app"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type CreateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type UpdateTaskInput struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type CompleteTaskInput struct {
	ID int `json:"id"`
}
type RemoveTaskInput struct {
	ID int `json:"id"`
}
type ListTasksInput struct{}
type ListTasksOutput struct {
	Tasks []string `json:"tasks"`
}

type Output struct {
	Message string `json:"message"`
}

func RegisterTools(server *mcp.Server, svc *app.TaskService) {

	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_task",
		Description: "Create a new task",
	}, func(ctx context.Context, req *mcp.CallToolRequest, input CreateTaskInput) (
		*mcp.CallToolResult,
		Output,
		error,
	) {
		svc.Add(input.Title, input.Description)

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Task created: " + input.Title},
			},
		}, Output{Message: "Task created successfully"}, nil
	})

	mcp.AddTool(server, &mcp.Tool{
		Name:        "complete_task",
		Description: "Mark task as completed",
	}, func(ctx context.Context, req *mcp.CallToolRequest, input CompleteTaskInput) (
		*mcp.CallToolResult,
		Output,
		error,
	) {

		err := svc.Complete(input.ID)
		if err != nil {
			return nil, Output{Message: "Failed to complete task"}, err
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Task %d completed", input.ID)},
			},
		}, Output{Message: "Task completed successfully"}, nil
	})

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_task",
		Description: "Update task",
	}, func(ctx context.Context, req *mcp.CallToolRequest, input UpdateTaskInput) (
		*mcp.CallToolResult,
		Output,
		error,
	) {
		svc.Edit(input.ID, input.Title, input.Description)

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Task %d updated, Title: %s, Description: %s", input.ID, input.Title, input.Description)},
			},
		}, Output{Message: "Task updated successfully"}, nil
	})

	mcp.AddTool(server, &mcp.Tool{
		Name:        "remove_task",
		Description: "Remove a task",
	}, func(ctx context.Context, req *mcp.CallToolRequest, input RemoveTaskInput) (
		*mcp.CallToolResult,
		Output,
		error,
	) {

		err := svc.Remove(input.ID)
		if err != nil {
			return nil, Output{Message: "Failed to remove task"}, err
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Task %d removed", input.ID)},
			},
		}, Output{Message: "Task removed successfully"}, nil
	})

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_tasks",
		Description: "List all tasks",
	}, func(ctx context.Context, req *mcp.CallToolRequest, input ListTasksInput) (
		*mcp.CallToolResult,
		ListTasksOutput,
		error,
	) {
		tasks := svc.List()

		var taskStrings []string
		for _, t := range tasks {
			status := "pending"
			if t.IsCompleted {
				status = "completed"
			}
			taskStrings = append(taskStrings, fmt.Sprintf("[%d] %s - %s (%s)", t.ID, t.Title, t.Description, status))
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Tasks:\n" + strings.Join(taskStrings, "\n")},
			},
		}, ListTasksOutput{Tasks: taskStrings}, nil
	})
}
