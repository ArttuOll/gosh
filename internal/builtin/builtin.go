package builtin

import (
	"os"
)

func IsBuiltInCommand(command string) bool {
	return command == "exit" || command == "cd"
}

func EvaluateBuiltInCommand(command string, args []string) {
	switch command {
	case "exit":
		Exit()
	default:
		return
	}
}

func Exit() {
	os.Exit(0)
}
