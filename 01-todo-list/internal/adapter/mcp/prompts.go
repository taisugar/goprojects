package adapter_mcp

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func RegisterPrompts(server *mcp.Server) {

	server.AddPrompt(&mcp.Prompt{
		Name:        "generate-random-tasks",
		Description: "Generate random tasks",
		Arguments: []*mcp.PromptArgument{
			{
				Name:        "goal",
				Title:       "Goal",
				Description: "The goal for which to generate tasks",
			},
		},
	}, func(ctx context.Context, req *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {

		return &mcp.GetPromptResult{
			Messages: []*mcp.PromptMessage{
				{
					Role: "user",
					Content: &mcp.TextContent{
						Text: "Generate a list of random tasks: " + req.Params.Arguments["goal"] + ". The tasks should be actionable and specific.",
					},
				},
			},
		}, nil
	})

	server.AddPrompt(&mcp.Prompt{
		Name:        "update-all-descriptions",
		Description: "Update all task descriptions",
	}, func(ctx context.Context, req *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {

		return &mcp.GetPromptResult{
			Messages: []*mcp.PromptMessage{
				{
					Role: "user",
					Content: &mcp.TextContent{
						Text: "Update all task 'Description' fields that are missing ensure each description is clear and actionable. Keep the title and other fields the same. Only update the description field. If a description is already present, keep it as is.",
					},
				},
			},
		}, nil
	})
}
