package builtin

import (
	"fmt"
	"os"
)

func IsBuiltInCommand(command string) bool {
	return command == "exit" || command == "cd"
}

func EvaluateBuiltInCommand(command string, args []string) error {
	switch command {
	case "exit":
		exit()
		return nil
	case "cd":
		return changeDirectory(args)
	default:
		return nil
	}
}

func exit() {
	os.Exit(0)
}

func changeDirectory(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("the cd built-in command only accepts one argument")
	}

	path := args[0]

	os.Chdir(path)

	return nil
}
