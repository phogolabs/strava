package tool

// Config represents the tool config
type Config struct {
	Vendor    []*VendorConfig    `yaml:"vendor"`
	Transform []*TransformConfig `yaml:"transform"`
}

// VendorConfig represents the vendor config
type VendorConfig struct {
	Name   string   `yaml:"name"`
	Source []string `yaml:"source"`
}

// TransformConfig represents the update config
type TransformConfig struct {
	Name  string        `yaml:"name"`
	Rules []*RuleConfig `yaml:"rules"`
}

// RuleConfig represents an experesion
type RuleConfig struct {
	Regexp string `yaml:"regexp"`
	Value  string `yaml:"value"`
}
