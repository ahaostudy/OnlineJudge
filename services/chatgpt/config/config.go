package config

var Config = new(Default)

type Default struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Openai  Openai `yaml:"openai"`
}

type Openai struct {
	BaseUrl string `yaml:"baseUrl"`
	ApiKey  string `yaml:"apiKey"`
	Model   string `yaml:"model"`
}
