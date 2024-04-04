package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

const (
	Exit = "exit"
)

func main() {
	fmt.Println("Hello to TODO app")

	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	for {
		runCommand(*command)

		fmt.Println("please enter another command:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(command string) {

	switch command {
	case Exit:
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}
}
