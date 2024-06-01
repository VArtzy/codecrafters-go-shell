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
    for {
        fmt.Fprint(os.Stdout, "$ ")

        // Wait for user input
        input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
        input = strings.TrimSpace(input)
        cmd := strings.Split(input, " ")
        switch cmd[0] {
        case "cd":
            path := handleCd(cmd[1])
            cmd[1] = path
        case "type":
            handleType(cmd[1])
        case "echo":
            fmt.Fprint(os.Stdout, strings.Join(cmd[1:], " ") + "\n")
        case "exit":
            os.Exit(0)
        default:
            handleExec(cmd[0], cmd[1:])
        }
    }
}

func handleCd(path string) string {
            if strings.TrimSpace(path) == "~" {
                return os.Getenv("HOME")
            }
            if err := os.Chdir(path); err != nil {
                fmt.Fprintf(os.Stdout, path + ": No such file or directory\n")
            }
}

func handleExec(cmd string, args []string) error {
            exe := exec.Command(cmd, args...)
            exe.Stderr = os.Stderr
            exe.Stdout = os.Stdout
            err :=  exe.Run()
            if err != nil {
                fmt.Fprint(os.Stdout, input[:len(input)] + ": command not found\n")
            }
}

func handleType(arg string) {
            switch arg {
            case "exit", "echo", "type":
                fmt.Fprint(os.Stdout, arg + " is a shell builtin\n")
            default:
                paths := strings.Split(os.Getenv("PATH"), ":")
                for _, path := range paths {
                    fp := filepath.Join(path, arg)
                    if _, err := os.Stat(fp); err == nil {
                        fmt.Fprint(os.Stdout, fp + "\n")
                        return
                    }
                }
                fmt.Fprintf(os.Stdout, arg + " not found\n")
            }
}
