package main

import (
	"Minishell_1_Go/built-ins"
	"Minishell_1_Go/display"
	"Minishell_1_Go/user_input"
	"fmt"
)

//func main() { os.Exit(myShell()) } // pour defer si jamais (Ex: close on defer)

func main() {
	var exitStatus uint8

SHELL_LOOP: // juste pour utiliser des labels
	for {
		display.Prompt(&exitStatus)
		name, argv := user_input.GetInput()
		switch {
		case name == "":
			continue SHELL_LOOP
		//alias.Check()
		case builtins.Check(&exitStatus, name, argv):
			continue SHELL_LOOP
			//case execve.Check(&exitStatus):
			//TO DO evaluate cmd
		default:
			fmt.Println("name", name, "\t", "argv", argv)
			display.PrintError(&exitStatus, "test", display.UNKNOWNCMD)
		}
	}

}
