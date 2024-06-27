---
dest: README.md
---

# reproduce-urfave-cli-shell-completion-issue

This repository includes a document and sample code to reproduce the issue of [urfave/cli](https://github.com/urfave/cli) shell completion.

The version of urfave/cli: [v2.27.2](go.mod).

## How to reproduce the issue

Go is required.

```sh
#-yodoc run
go version
```

```
# Output
{{.CombinedOutput -}}
```

There are two commands built with urfave/cli/v2.

- [root](cmd/root/main.go)
- [child](cmd/child/main.go)

1. Build these commands

```sh
#-yodoc run
go build -o dist/root ./cmd/root
go build -o dist/child ./cmd/child
```

Bash completion is enabled.

```sh
#-yodoc run
./dist/root --generate-bash-completion
```

```
# Output
{{.CombinedOutput -}}
```

```sh
#-yodoc run
./dist/child --generate-bash-completion
```

```
# Output
{{.CombinedOutput -}}
```

`./dist/root exec -- <command> [<argument> ...]` executes a command `<command> [<argument> ...]`.
For example, `./dist/root exec -- git version` executes `git version`, and `./dist/root exec -- ./dist/child --help` executes `./dist/child --help`.

```sh
#-yodoc run
./dist/root exec -- git version
```

```
# Output
{{.CombinedOutput -}}
```

```sh
#-yodoc run
./dist/root exec -- ./dist/child --help
```

```
# Output
{{.CombinedOutput -}}
```

So we expect `./dist/root exec -- ./dist/child --generate-bash-completion` executes `./dist/child --generate-bash-completion`, but actually it isn't.

```sh
#-yodoc run
./dist/root exec -- ./dist/child --generate-bash-completion
```

```
# Output
{{.CombinedOutput -}}
```

The output is different from the output of `./dist/child --generate-bash-completion`.

```sh
#-yodoc run
./dist/child --generate-bash-completion
```

```
# Output
{{.CombinedOutput -}}
```

## Note

After double dash `--` only positional arguments are accepted.

https://unix.stackexchange.com/a/11382

So when we executes `./dist/root exec -- ./dist/child --generate-bash-completion`, `./dist-root` should treat the argument `--generate-bash-completion` as a positional argument, but actually it treats the argument as a flag.
This is the cause of this issue.
