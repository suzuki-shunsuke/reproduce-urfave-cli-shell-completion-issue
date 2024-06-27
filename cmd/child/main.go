package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	if err := core(); err != nil {
		log.Fatal(err)
	}
}

func core() error {
	app := cli.App{
		Name:                 "child",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:   "add",
				Usage:  "add a new task",
				Action: action,
			},
			{
				Name:   "list",
				Usage:  "list tasks",
				Action: action,
			},
		},
	}

	return app.RunContext(context.Background(), os.Args)
}

func action(c *cli.Context) error {
	fmt.Println("Hello from child")
	return nil
}
