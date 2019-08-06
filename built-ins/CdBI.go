package builtins

import (
	"Minishell_1_Go/display"
	"fmt"
	"os"
	"strings"
)

func cd(exitStatus *uint8, name string, argv []string) bool {
	argvLen := len(argv)
	switch {
	case argvLen == 0:
		key, ok := os.LookupEnv("HOME")
		if !ok {
			return display.PrintError(exitStatus, "cd: No home directory.", display.CUSTOMPRINT)
		}
		changeDirectory(exitStatus, key)
	case argvLen == 1 && argv[0][0] == '~':
		key, ok := os.LookupEnv("HOME")
		if !ok {
			return display.PrintError(exitStatus, "No $home variable set.", display.CUSTOMPRINT)
		}
		key += strings.TrimPrefix(argv[0], "~")
		changeDirectory(exitStatus, key)
	case argvLen == 1 && argv[0] == "-":
		lastDir(exitStatus)
	case argvLen == 1:
		changeDirectory(exitStatus, argv[0])
	default:
		display.PrintError(exitStatus, name, display.MANY)
	}
	return true
}

func changeDirectory(exitStatus *uint8, path string) {
	oldPwd := searchOldPwd()
	if err := os.Chdir(path); err != nil {
		fmt.Println(err.Error())
		return
	}
	changePwdPathEnvVar(exitStatus, oldPwd)
}

func searchOldPwd() string {
	oldPwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if oldPwd == "" {
		if key, ok := os.LookupEnv("PWD"); ok {
			oldPwd = key
		}
	}
	return oldPwd
}

func changePwdPathEnvVar(exitStatus *uint8, oldPwd string) {
	putInEnv(exitStatus, "OLDPWD", oldPwd)
	if oldPwd, err := os.Getwd(); err == nil {
		if err := os.Setenv("PWD", oldPwd); err != nil {
			fmt.Println(err)
			*exitStatus = 2
		}
	}
}

func lastDir(exitStatus *uint8) {
	oldPwd, ok := os.LookupEnv("OLDPWD")
	if !ok {
		display.PrintError(exitStatus, "", display.PATHNOTFOUND)
		return
	}
	changeDirectory(exitStatus, oldPwd)
	return
}
