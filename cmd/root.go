package cmd

import (
	"os"

	"github.com/ArttuOll/gosh/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gosh",
	Short: "A simple shell written in Go",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: internal.Run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
