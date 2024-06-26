# Creating own shell in Go Programming Languange
Creating shell REPL in go that can run built-in OS shell command to learn about how shell works.

- REPL using go for loop
- Execute executable file
- Print output
- Use OS shell command with go os/exec
- Learn type, cd, pwd, echo deeper as builtin command

## Future?

- Error handling

[![progress-banner](https://backend.codecrafters.io/progress/shell/987db4eb-110e-4ee1-8724-d9033bdb9fef)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

This is a starting point for Go solutions to the
["Build Your Own Shell" Challenge](https://app.codecrafters.io/courses/shell/overview).

_Add a description of your course here_

**Note**: If you're viewing this repo on GitHub, head over to
[codecrafters.io](https://codecrafters.io) to try the challenge.

# Passing the first stage

The entry point for your `shell` implementation is in `cmd/myshell/main.go`.
Study and uncomment the relevant code, and push your changes to pass the first
stage:

```sh
git add .
git commit -m "pass 1st stage" # any msg
git push origin master
```

Time to move on to the next stage!

# Stage 2 & beyond

Note: This section is for stages 2 and beyond.

1. Ensure you have `go (1.19)` installed locally
1. Run `./your_shell.sh` to run your program, which is implemented in
   `cmd/myshell/main.go`.
1. Commit your changes and run `git push origin master` to submit your solution
   to CodeCrafters. Test output will be streamed to your terminal.
