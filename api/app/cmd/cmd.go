package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = RootCommand()

func RootCommand() *cobra.Command {
	return &cobra.Command{}
}

func Execute(port string) {
	rootCmd.AddCommand(NewRESTCommand(port))
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
