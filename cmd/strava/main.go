package main

import (
	"os"

	"github.com/phogolabs/cli"
	"github.com/phogolabs/strava/tool"
)

var (
	// version variable is injected at compile time via
	// -ldflags "-X main.version=<sha>"
	version = "dev"

	// application command arguments
	flags = []cli.Flag{
		&cli.YAMLFlag{
			Name:     "config",
			Usage:    "Configuration file",
			FilePath: "strava.yaml",
			Value:    &tool.Config{},
		},
	}

	// application commands
	commands = []*cli.Command{
		&cli.Command{
			Name:      "vendor",
			Usage:     "Downloads all vendor dependencies",
			UsageText: "strava vendor [global options]",
			Action:    vendor,
		},
		&cli.Command{
			Name:      "apply",
			Usage:     "Transforms a given list of files",
			UsageText: "strava transform [global options]",
			Action:    apply,
		},
	}
)

func main() {
	app := &cli.App{
		Name:      "strava",
		HelpName:  "strava",
		Usage:     "Manages your vendor content",
		UsageText: "strava [global options]",
		Copyright: cli.Copyright("Phogo Labs"),
		Version:   version,
		Flags:     flags,
		Writer:    os.Stdout,
		ErrWriter: os.Stderr,
		Commands:  commands,
	}

	app.Run(os.Args)
}

func vendor(ctx *cli.Context) error {
	config, _ := ctx.GlobalGet("config").(*tool.Config)
	// prepare the runner
	runner := tool.NewVendor(config)
	// execute
	return runner.Run(ctx)
}

func apply(ctx *cli.Context) error {
	config, _ := ctx.GlobalGet("config").(*tool.Config)
	// prepare the runner
	runner := tool.NewTransform(config)
	// execute
	return runner.Run(ctx)
}
