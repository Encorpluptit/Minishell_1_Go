package builtins

import (
	"Minishell_1_Go/display"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func exitCmd(exitStatus *uint8, name string, argv []string) bool {
	cmdLen := len(argv)

	switch cmdLen {
	case 0:
		os.Exit(0)
	case 1:
		strings.Trim(argv[0], "()")
		exitCode, err := strconv.Atoi(argv[0])
		if err == nil {
			os.Exit(exitCode)
		}
		fmt.Println(err) // debug while coding
		return display.PrintError(exitStatus, "exit: Badly formed number.", display.CUSTOMPRINT)
	default:
		display.PrintError(exitStatus, name, display.MANY)
	}
	return true
}
