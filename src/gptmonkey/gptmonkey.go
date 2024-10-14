package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/Wason1797/gptmonkey/actions"
)

func main() {
	app := &cli.App{
		Name:   "gptmonkey",
		Usage:  "Ask Codellama about anything. Ideally for shell commands ()",
		Action: actions.MainAction,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
