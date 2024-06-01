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
        switch cmd[0] {
        case "type":
            switch cmd[1] {
            case "exit", "echo", "type":
                fmt.Fprint(os.Stdout, cmd[1] + " is a shell builtin\n")
            default:
                fmt.Fprintf(os.Stdout, cmd[1] + " not found\n")
            }
        case "echo":
            fmt.Fprint(os.Stdout, strings.Join(cmd[1:], " ") + "\n")
        case "exit":
            os.Exit(0)
        default:
            fmt.Fprint(os.Stdout, input[:len(input)] + ": command not found\n")
        }
    }
}
