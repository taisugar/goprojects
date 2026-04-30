package main

import (
	"context"
	_ "embed"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"

	adapter_mcp "task-manager/internal/adapter/mcp"
	"task-manager/internal/app"
	"task-manager/internal/infra"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var (
	httpAddr  = flag.String("http", "", "if set, use streamable HTTP at this address, instead of stdin/stdout")
	pprofAddr = flag.String("pprof", "", "if set, host the pprof debugging server at this address")
)

func main() {
	flag.Parse()

	if *pprofAddr != "" {
		// For debugging memory leaks, add an endpoint to trigger a few garbage
		// collection cycles and ensure the /debug/pprof/heap endpoint only reports
		// reachable memory.
		http.DefaultServeMux.Handle("/gc", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			for range 3 {
				runtime.GC()
			}
			fmt.Fprintln(w, "GC'ed")
		}))
		go func() {
			// DefaultServeMux was mutated by the /debug/pprof import.
			http.ListenAndServe(*pprofAddr, http.DefaultServeMux)
		}()
	}

	// For simplicity, we use a JSON file for storage. In a real application, this could be a database or other persistent storage.
	store, _ := infra.NewTaskRepository(infra.FormatJSON)

	svc := app.NewTaskService(store)
	svc.Load()

	opts := &mcp.ServerOptions{
		Instructions:      "Use this server!",
		CompletionHandler: complete, // support completions by setting this handler
	}
	icons := mcpIcons()
	server := mcp.NewServer(
		&mcp.Implementation{Name: "todo-mcp", Version: "0.0.1", Icons: icons},
		opts,
	)
	adapter_mcp.RegisterMCP(server, svc)

	if *httpAddr != "" {
		handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server {
			return server
		}, nil)
		log.Printf("MCP handler listening at %s", *httpAddr)
		if *pprofAddr != "" {
			log.Printf("pprof listening at http://%s/debug/pprof", *pprofAddr)
		}
		log.Fatal(http.ListenAndServe(*httpAddr, handler))
	} else {
		t := &mcp.LoggingTransport{Transport: &mcp.StdioTransport{}, Writer: os.Stderr}
		if err := server.Run(context.Background(), t); err != nil {
			log.Printf("Server failed: %v", err)
		}
	}

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	} else {
		log.Println("MCP server stopped")
	}
}

func complete(ctx context.Context, req *mcp.CompleteRequest) (*mcp.CompleteResult, error) {
	return &mcp.CompleteResult{
		Completion: mcp.CompletionResultDetails{
			Total:  1,
			Values: []string{req.Params.Argument.Value + "x"},
		},
	}, nil
}

//go:embed mcp.png
var mcpIconData []byte

func mcpIcons() []mcp.Icon {
	return []mcp.Icon{{
		Source:   "data:image/png;base64," + base64.StdEncoding.EncodeToString(mcpIconData),
		MIMEType: "image/png",
		Sizes:    []string{"48x48"},
		Theme:    mcp.IconThemeLight,
	}}
}
