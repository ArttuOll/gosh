package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) error {
	stdInReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("gosh> ")
		input, err := stdInReader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read from stdin: %w", err)
		}

		parts := strings.Fields(input)
		command := parts[0]

		if command == "exit" {
			os.Exit(0)
		}
	}
}
