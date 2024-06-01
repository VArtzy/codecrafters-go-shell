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
        case "echo":
            fmt.Fprint(os.Stdout, strings.Join(cmd[1:], " ") + "\n")
        case "exit 0":
            os.Exit(0)
        default:
            fmt.Fprint(os.Stdout, input[:len(input)] + ": command not found\n")
        }
    }
}
