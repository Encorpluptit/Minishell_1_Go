package builtins

import (
	"Minishell_1_Go/display"
	"fmt"
	"os"
)

func env(exitStatus *uint8, name string, argv []string) bool {
	if len(argv) > 0 {
		return display.PrintError(exitStatus, name, display.MANY)
	}
	for _, key := range os.Environ() {
		fmt.Println(key)
	}
	return true
}
