package main

import (
	"io/ioutil"
	"os"

	"github.com/phogolabs/cli"
	"github.com/phogolabs/protoc-gen-vendor/tool"
	"google.golang.org/protobuf/proto"
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
			FilePath: "vendor.yaml",
			Value:    &tool.Config{},
		},
	}
)

func main() {
	app := &cli.App{
		Name:      "protoc-gen-vendor",
		HelpName:  "protoc-gen-vendor",
		Usage:     "protoc-gen-vendor",
		UsageText: "protoc-gen-vendor [global options]",
		Copyright: cli.Copyright("Phogo Labs"),
		Version:   version,
		Flags:     flags,
		Writer:    os.Stdout,
		ErrWriter: os.Stderr,
		Action:    run,
		Commands: []*cli.Command{
			&cli.Command{
				Name:   "run",
				Action: run,
			},
		},
	}

	app.Run(os.Args)
}

func run(ctx *cli.Context) error {
	var (
		config, _ = ctx.GlobalGet("config").(*tool.Config)
		input     = &tool.CodeGeneratorRequest{}
	)

	if ctx.Command.Name == "protoc-gen-vendor" {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		// You must use the requests unmarshal method to handle this type
		if err := proto.Unmarshal(data, input); err != nil {
			return err
		}
	}

	plugin := &tool.Plugin{
		Config: config,
	}

	output, err := plugin.Handle(input)
	if err != nil {
		return err
	}

	payout, err := proto.Marshal(output)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(payout)
	return err
}
