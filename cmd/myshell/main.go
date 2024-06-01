package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    for true {
        fmt.Fprint(os.Stdout, "$ ")

        // Wait for user input
        input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
        fmt.Fprint(os.Stdout, input[:len(input) - 1] + ": command not found\n")
    }
}
