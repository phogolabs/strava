package tool

// Config represents the tool config
type Config struct {
	Vendor  []*VendorConfig `yaml:"vendor"`
	Actions []*ActionConfig `yaml:"actions"`
}

// VendorConfig represents the vendor config
type VendorConfig struct {
	Name   string   `yaml:"name"`
	Source []string `yaml:"source"`
}

// ActionConfig represents the action config
type ActionConfig struct {
	Name        string        `yaml:"name"`
	Expressions []*ExprConfig `yaml:"expressions"`
}

// ExprConfig represents an experesion
type ExprConfig struct {
	Syntax string `yaml:"syntax"`
	Value  string `yaml:"value"`
}
