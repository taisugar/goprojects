package cmd

import (
	"task-manager/model"
	"task-manager/service"
	"task-manager/storage"

	"github.com/spf13/cobra"
)

type format string

const (
	formatCSV  format = "csv"
	formatJSON format = "json"
)

var rootCmd = &cobra.Command{
	Use: "todo",
}

func Execute() {
	rootCmd.Execute()
}

var svc *service.TaskService

func initService(format format) {
	var store storage.Strategy[[]model.Task]

	if format == formatCSV {
		store = storage.NewCSV("tasks.csv")
	} else {
		store = storage.NewJSON[[]model.Task]("tasks.json")
	}

	svc = service.NewTaskService(store)
	svc.Load()
}
