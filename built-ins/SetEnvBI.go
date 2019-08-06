package builtins

import (
	"Minishell_1_Go/display"
	"fmt"
	"os"
	"unicode"
)

//type chkS func(exitStatus *uint8, key string) bool
//var check chkS = func(exitStatus *uint8, key string) bool { return (checkKeySynthax(exitStatus, key))}
// il y a surement une autre solution mais ça m'entraine pour les functionsvariables

func setEnv(exitStatus *uint8, name string, argv []string) bool {
	argvLen := len(argv)
	switch {
	case argvLen == 0:
		return env(exitStatus, name, argv)
	case argvLen == 1 && checkKeySynthax(exitStatus, argv[0]):
		return putInEnv(exitStatus, argv[0], "")
	case argvLen == 2 && checkKeySynthax(exitStatus, argv[0]):
		return putInEnv(exitStatus, argv[0], argv[1])
	case argvLen > 2:
		return display.PrintError(exitStatus, name, display.MANY)
	}
	return true
}

func checkKeySynthax(exitStatus *uint8, key string) bool {
	f2 := func(ch rune) bool {
		return unicode.IsLetter(ch) || unicode.IsDigit(ch)
	}
	for _, ch := range key {
		if !f2(ch) {
			display.PrintError(exitStatus, "setenv : key most contain only alphanum characters.", display.CUSTOMPRINT)
			return false
		}
	}
	return true
}

func putInEnv(exitStatus *uint8, key string, value string) bool {
	if err := os.Setenv(key, value); err != nil {
		fmt.Println(err)
		*exitStatus = 2
	}
	return true
}
