package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	h, err := NewHub(sugar)
	if err != nil {
		sugar.Fatalf("error while initializing hub: %v", err)
	}

	app := &cli.App{
		Name:  "pandora",
		Usage: "Stores secrets safely",
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Initializes a pandora's box",
				Action:  h.initBox,
			},
			{
				Name:    "lock",
				Aliases: []string{"l"},
				Usage:   "Lock's the pandora's box",
				Action:  h.LockBox,
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add a file to pandora's box",
				Action:  h.AddToBox,
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatalln(app.Run(os.Args))
	}
}
