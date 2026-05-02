package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	flag.Parse()
	ctx := context.Background()

	// Create MCP client
	client := mcp.NewClient(&mcp.Implementation{Name: "todo-cli-client", Version: "0.0.1"}, nil)

	// Connect to the MCP server via stdio transport
	// The server is expected to be running or passed via command
	serverPath := os.Args[0]
	fmt.Printf("Connecting to MCP server at %s...\n", serverPath)
	if len(os.Args) > 1 {
		serverPath = os.Args[1]
	}

	// serverPath := "main.go"
	transport := &mcp.CommandTransport{Command: exec.Command(serverPath)}
	session, err := client.Connect(context.Background(), transport, nil)
	if err != nil {
		log.Fatal("Failed to connect to MCP server:", err)
	}
	defer session.Close()

	// Test: List available tools
	tools, err := session.ListTools(context.Background(), nil)
	if err != nil {
		log.Fatal("Failed to list tools:", err)
	}
	fmt.Println("Available tools:")
	for _, t := range tools.Tools {
		fmt.Printf("  - %s: %s\n", t.Name, t.Description)
	}

	// // Example: Call create_task tool
	// params := &mcp.CallToolParams{
	// 	Name: "create_task",
	// 	Arguments: map[string]any{
	// 		"title":       "Test task from CLI client",
	// 		"description": "This is a test task created via MCP client",
	// 	},
	// }
	// res, err := session.CallTool(context.Background(), params)
	// if err != nil {
	// 	log.Fatal("CallTool failed:", err)
	// }
	// fmt.Printf("Create task result: %+v\n", res)

	// // Example: Call list_tasks tool (if available)
	// listParams := &mcp.CallToolParams{
	// 	Name:      "list_tasks",
	// 	Arguments: map[string]any{},
	// }
	// res, err = session.CallTool(context.Background(), listParams)
	// if err != nil {
	// 	log.Printf("list_tasks not available: %v", err)
	// } else {
	// 	fmt.Printf("List tasks result: %+v\n", res)
	// }
	// 🔥 background loop
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch {
		case strings.HasPrefix(input, "add "):
			title := strings.TrimPrefix(input, "add ")

			callTool(ctx, session, "create_task", map[string]any{
				"title":       title,
				"description": "",
			})

		case input == "list":
			callTool(ctx, session, "list_tasks", map[string]any{})

		case input == "exit":
			fmt.Println("bye")
			return

		default:
			fmt.Println("commands: add <title> | list | exit")
		}
	}
}

func callTool(ctx context.Context, s *mcp.ClientSession, name string, args map[string]any) {
	res, err := s.CallTool(ctx, &mcp.CallToolParams{
		Name:      name,
		Arguments: args,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("result:", res)
}
