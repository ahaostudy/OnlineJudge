package config

var Config = new(Default)

type Default struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`

	Static Static `yaml:"static"`
}

type Static struct {
	URI  string `yaml:"uri"`
	Path string `yaml:"path"`
}
