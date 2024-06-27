package main

import (
	"context"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {
	if err := core(); err != nil {
		log.Fatal(err)
	}
}

func core() error {
	app := cli.App{
		Name:                 "root",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:   "exec",
				Usage:  "execute a command",
				Action: execAction,
			},
		},
	}

	return app.RunContext(context.Background(), os.Args)
}

func execAction(c *cli.Context) error {
	cmd := exec.CommandContext(c.Context, c.Args().First(), c.Args().Tail()...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
