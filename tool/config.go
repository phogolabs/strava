package tool

// Config represents the tool config
type Config struct {
	Vendor []*VendorConfig `yaml:"vendor"`
}

// VendorConfig represents the vendor config
type VendorConfig struct {
	Name         string   `yaml:"name"`
	Dependencies []string `yaml:"dependencies"`
}
