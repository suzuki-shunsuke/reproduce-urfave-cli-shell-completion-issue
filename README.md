# reproduce-urfave-cli-shell-completion-issue

This repository includes a document and sample code to reproduce the issue of [urfave/cli](https://github.com/urfave/cli) shell completion.

The version of urfave/cli: [v2.27.2](go.mod).

## How to reproduce the issue

Go is required.

```sh
go version
```
```
# Output
go version go1.22.4 darwin/arm64
```

There are two commands built with urfave/cli/v2.

- [root](cmd/root/main.go)
- [child](cmd/child/main.go)

1. Build these commands

```sh
go build -o dist/root ./cmd/root
go build -o dist/child ./cmd/child
```

Bash completion is enabled.

```sh
./dist/root --generate-bash-completion
```
```
# Output
exec:execute a command
help:Shows a list of commands or help for one command
h:Shows a list of commands or help for one command
```
```sh
./dist/child --generate-bash-completion
```
```
# Output
add:add a new task
list:list tasks
help:Shows a list of commands or help for one command
h:Shows a list of commands or help for one command
```

`./dist/root exec -- <command> [<argument> ...]` executes a command `<command> [<argument> ...]`.
For example, `./dist/root exec -- git version` executes `git version`, and `./dist/root exec -- ./dist/child --help` executes `./dist/child --help`.

```sh
./dist/root exec -- git version
```
```
# Output
git version 2.39.3 (Apple Git-146)
```
```sh
./dist/root exec -- ./dist/child --help
```
```
# Output
NAME:
   child - A new cli application

USAGE:
   child [global options] command [command options] 

COMMANDS:
   add      add a new task
   list     list tasks
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

So we expect `./dist/root exec -- ./dist/child --generate-bash-completion` executes `./dist/child --generate-bash-completion`, but actually it isn't.

```sh
./dist/root exec -- ./dist/child --generate-bash-completion
```
```
# Output
help:Shows a list of commands or help for one command
h:Shows a list of commands or help for one command
```

The output is different from the output of `./dist/child --generate-bash-completion`.

```sh
./dist/child --generate-bash-completion
```
```
# Output
add:add a new task
list:list tasks
help:Shows a list of commands or help for one command
h:Shows a list of commands or help for one command
```

## Note

After double dash `--` only positional arguments are accepted.

https://unix.stackexchange.com/a/11382

So when we executes `./dist/root exec -- ./dist/child --generate-bash-completion`, `./dist-root` should treat the argument `--generate-bash-completion` as a positional argument, but actually it treats the argument as a flag.
This is the cause of this issue.

<!-- This file is generated by yodoc.
https://github.com/suzuki-shunsuke/yodoc
Please don't edit this code comment because yodoc depends on this code comment.
-->
