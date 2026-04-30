package main

import (
	"log"
	adapter_cli "task-manager/internal/adapter/cli"
	"task-manager/internal/app"
	"task-manager/internal/infra"
)

func main() {
	store, _ := infra.NewTaskRepository(infra.FormatJSON)
	svc := app.NewTaskService(store)
	svc.Load()

	rootCmd := adapter_cli.RegisterCLI(svc)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
