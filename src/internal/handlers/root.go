package handlers

import (
	"brodsky/internal/info"
	"brodsky/internal/log"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RootPersistentPreRunE(cmd *cobra.Command, _ []string) error {
	vVerbose, _ := cmd.Flags().GetBool("verbose")
	if vVerbose {
		log.InitializeLogger(logrus.DebugLevel)
	} else {
		log.InitializeLogger(logrus.InfoLevel)
	}

	return nil
}

func RootPersistentRunE(cmd *cobra.Command, _ []string) error {
	vFlag, _ := cmd.Flags().GetBool("version")
	if vFlag {
		fmt.Printf("%s version: %s\n", info.GetAppName(), info.GetVersion())
		return nil
	}

	fmt.Println("Please provide a command.")
	err := cmd.Help()
	if err != nil {
		return err
	}
	return nil
}
