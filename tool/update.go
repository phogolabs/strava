package tool

import (
	"io/ioutil"
	"regexp"

	"github.com/phogolabs/cli"
)

// Transform represents the tool transform command
type Transform struct {
	config *Config
}

// NewTransform creates a new transform
func NewTransform(cfg *Config) *Transform {
	return &Transform{
		config: cfg,
	}
}

// Run runs the vendor command
func (p *Transform) Run(_ *cli.Context) error {
	for _, vendor := range p.config.Transform {
		data, err := ioutil.ReadFile(vendor.Name)
		if err != nil {
			return err
		}

		for _, rule := range vendor.Rules {
			expr, err := regexp.Compile(rule.Regexp)
			if err != nil {
				return err
			}

			data = expr.ReplaceAll(data, []byte(rule.Value))
		}

		if err := ioutil.WriteFile(vendor.Name, data, 0700); err != nil {
			return err
		}
	}

	return nil
}
