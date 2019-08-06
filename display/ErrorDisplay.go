package display

import (
	"fmt"
)

const (
	UNKNOWNCMD   uint8 = 0
	CUSTOMPRINT  uint8 = 1
	FEW          uint8 = 2
	MANY         uint8 = 3
	PATHNOTFOUND uint8 = 4
)

func PrintError(exitStatus *uint8, name string, message uint8) bool {
	switch message {
	case CUSTOMPRINT:
		fmt.Println(name)
	case MANY:
		fmt.Printf("%s: Too many arguments.\n", name)
	case FEW:
		fmt.Printf("%s: Too few arguments.\n", name)
	case PATHNOTFOUND:
		fmt.Printf("%s: No such file or directory.\n", name)
	default:
		fmt.Printf("%s: Unknown error.\n", name)
	}
	*exitStatus = 1
	return true
}
