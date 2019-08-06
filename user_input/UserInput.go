package user_input

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetInput() (cmd string, argv []string) {
	var input string
	var err error

	for {
		reader := bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		switch {
		case err == nil && !mismatchChar(input) && input != "\n":
			return "", nil
		case err == nil && input != "\n":
			break
		case err == io.EOF:
			fmt.Println("exit\n")
			os.Exit(0)
		default:
			if err != nil {
				fmt.Println(err)
			}
			return "", nil
		}
		break
	}
	argv = strings.Fields(input)
	name := argv[0]
	argv = argv[1:]
	return name, argv
}

func mismatchChar(input string) bool {
	switch {
	case strings.Count(input, "'")%2 != 0:
		fmt.Println("Unmatched : '")
		return false
	case strings.Count(input, "\"")%2 != 0:
		fmt.Println("Unmatched : \"")
		return false
	default:
		return (parenthesis(input))
	}
}

func parenthesis(input string) bool {
	leftParen := strings.Count(input, "(")
	rightParen := strings.Count(input, ")")
	switch {
	case leftParen > rightParen:
		fmt.Println("Too many's '('")
		return false
	case rightParen > leftParen:
		fmt.Println("Too many's ')'")
		return false
	}
	return true
}
