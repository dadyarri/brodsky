package handlers

import (
	"brodsky/internal/log"
	"github.com/spf13/cobra"
)

func ServeRunE(cmd *cobra.Command, _ []string) error {
	return log.ExecutionTime(func() error { return handleServeRun(cmd) })
}

func handleServeRun(cmd *cobra.Command) error {
	return nil
}
