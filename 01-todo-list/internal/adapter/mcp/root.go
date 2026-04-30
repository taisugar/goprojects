package adapter_mcp

import (
	"task-manager/internal/app"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func RegisterMCP(server *mcp.Server, svc *app.TaskService) {
	RegisterTools(server, svc)
	RegisterResources(server, svc)
	RegisterPrompts(server)
}
