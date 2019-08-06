package display

import (
	"fmt"
	"os"
	"strings"
)

var saveUser string // ou  mettre la var dans l'env

// TODO dotbg(background)

func Prompt(exitStatus *uint8) {
	user := searchUser()
	pwd := getPwd()
	fmt.Printf("%d %s → %s ", *exitStatus, user, pwd)
}

func getPwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		return err.Error()
	}

	wd := strings.SplitAfter(pwd, "/")
	if len(wd) > 2 {
		return wd[len(wd)-1]
	} else {
		return pwd
	}
}

func searchUser() string {
	if saveUser != "" {
		return saveUser
	}

	userPath := [3]string{"USER", "LOGNAME", "GROUP"}
	for _, path := range userPath {
		key, ok := os.LookupEnv(path)
		if ok {
			saveUser = key
			return saveUser
		}
	}
	saveUser = "Unknown"
	return saveUser
}
