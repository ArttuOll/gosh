package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/ArttuOll/gosh/internal/builtin"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	stdInReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("gosh> ")
		input, err := stdInReader.ReadString('\n')
		if err != nil {
			log.Fatal("failed to read from stdin: %w", err)
		}

		parts := strings.Fields(input)
		command := parts[0]
		arguments := parts[1:]

		if builtin.IsBuiltInCommand(command) {
			builtin.EvaluateBuiltInCommand(command, arguments)
			continue
		}

		cmd := exec.Command(command, arguments...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
