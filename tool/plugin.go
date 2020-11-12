package tool

import (
	"context"
	"os"

	"github.com/hashicorp/go-getter"
)

// Plugin represents the tool plugin
type Plugin struct {
	Config *Config
}

// Handle handle the plugin
func (p *Plugin) Handle(req *CodeGeneratorRequest) (*CodeGeneratorResponse, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	for _, vendor := range p.Config.Vendor {
		if err := os.MkdirAll(vendor.Name, 0700); err != nil {
			return nil, err
		}

		for _, source := range vendor.Dependencies {
			client := &getter.Client{
				Ctx:     context.TODO(),
				Src:     source,
				Dst:     vendor.Name,
				Pwd:     pwd,
				Mode:    getter.ClientModeAny,
				Options: []getter.ClientOption{},
			}

			if err := client.Get(); err != nil {
				return nil, err
			}
		}
	}

	return &CodeGeneratorResponse{}, nil
}
