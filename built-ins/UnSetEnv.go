package builtins

import (
	"Minishell_1_Go/display"
	"fmt"
	"os"
)

func unSetEnv(exitStatus *uint8, name string, argv []string) bool {
	if len(argv) == 0 {
		return display.PrintError(exitStatus, name, display.FEW)
	}
	for _, i := range argv {
		if err := os.Unsetenv(i); err != nil {
			fmt.Println(err)
		}
	}
	return true
}
