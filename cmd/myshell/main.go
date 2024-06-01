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
        switch input {
        case "exit 0":
            os.Exit(0)
        default:
            fmt.Fprint(os.Stdout, input[:len(input) - 1] + ": command not found\n")
        }
    }
}
