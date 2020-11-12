package tool

import (
	"os"

	"github.com/hashicorp/go-getter"
	"github.com/phogolabs/cli"
)

// Vendor represents the tool vendor command
type Vendor struct {
	config *Config
}

// NewVendor creates a new vendor
func NewVendor(cfg *Config) *Vendor {
	return &Vendor{
		config: cfg,
	}
}

// Run runs the vendor command
func (p *Vendor) Run(_ *cli.Context) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	tracker := NewProgressTracker()

	for _, vendor := range p.config.Vendor {
		if err := os.MkdirAll(vendor.Name, 0700); err != nil {
			return err
		}

		for _, source := range vendor.Source {
			client := &getter.Client{
				Src:  source,
				Dst:  vendor.Name,
				Pwd:  pwd,
				Mode: getter.ClientModeAny,
				Options: []getter.ClientOption{
					getter.WithProgress(tracker),
				},
			}

			tracker.Add(1)

			go func() {
				if err := client.Get(); err != nil {
					//TODO:
				}

				tracker.Add(-1)
			}()
		}
	}

	tracker.Wait()
	return nil
}

// func (p *Plugin) action() error {
// 	for _, action := range p.Config.Actions {
// 		data, err := ioutil.ReadFile(action.Name)
// 		if err != nil {
// 			return err
// 		}

// 		content := string(data)

// 		for _, expr := range action.Expressions {
// 			syntax, err := regexp.Compile(expr.Syntax)
// 			if err != nil {
// 				return err
// 			}

// 			content = syntax.ReplaceAllString(content, expr.Value)
// 		}

// 		data = []byte(content)

// 		if err := ioutil.WriteFile(action.Name, data, 0700); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
