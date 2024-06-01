package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
)

func main() {
    for {
        fmt.Fprint(os.Stdout, "$ ")

        // Wait for user input
        input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
        input = strings.TrimSpace(input)
        cmd := strings.Split(input, " ")
        switch input {
        case "echo":
            fmt.Fprint(os.Stdout, strings.Join(command[1:], " ") + "\n")
        case "exit 0":
            os.Exit(0)
        default:
            fmt.Fprint(os.Stdout, input[:len(input)] + ": command not found\n")
        }
    }
}
