package main

import (
	"bufio"
	"fmt"
	"os"
    "os/exec"
    "strings"
    "path/filepath"
)

func main() {
    Exit:
    for {
        fmt.Fprint(os.Stdout, "$ ")

        // Wait for user input
        input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
        input = strings.TrimSpace(input)
        cmd := strings.Split(input, " ")
        switch cmd[0] {
        case "cd":
            if err := os.Chdir(cmd[1]); err != nil {
                fmt.Fprintf(os.Stdout, cmd[1] + ": No such file or directory\n")
            }
        case "type":
            switch cmd[1] {
            case "exit", "echo", "type":
                fmt.Fprint(os.Stdout, cmd[1] + " is a shell builtin\n")
            default:
                paths := strings.Split(os.Getenv("PATH"), ":")
                for _, path := range paths {
                    fp := filepath.Join(path, cmd[1])
                    if _, err := os.Stat(fp); err == nil {
                        fmt.Fprint(os.Stdout, fp + "\n")
                        continue Exit
                    }
                }
                fmt.Fprintf(os.Stdout, cmd[1] + " not found\n")
            }
        case "echo":
            fmt.Fprint(os.Stdout, strings.Join(cmd[1:], " ") + "\n")
        case "exit":
            os.Exit(0)
        default:
            exe := exec.Command(cmd[0], cmd[1:]...)
            exe.Stderr = os.Stderr
            exe.Stdout = os.Stdout
            err := exe.Run()
            if err != nil {
                fmt.Fprint(os.Stdout, input[:len(input)] + ": command not found\n")
            }
        }
    }
}
