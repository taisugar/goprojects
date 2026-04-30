package adapter_mcp

import (
	"context"
	"encoding/json"

	"task-manager/internal/app"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func RegisterResources(server *mcp.Server, svc *app.TaskService) {

	server.AddResource(&mcp.Resource{
		Name:        "tasks_info",
		Description: "List all tasks",
		MIMEType:    "text/plain",
		URI:         "tasks://list",
	}, func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {

		tasks := svc.List()

		data, _ := json.MarshalIndent(tasks, "", "  ")

		return &mcp.ReadResourceResult{
			Contents: []*mcp.ResourceContents{
				{URI: req.Params.URI, MIMEType: "text/plain", Text: string(data)},
			},
		}, nil
	})
}
