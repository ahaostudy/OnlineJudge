package config

var Config = new(Default)

// Default
type Default struct {
	Name           string  `yaml:"name"`
	Version        string  `yaml:"version"`
	Port           int     `yaml:"port"`
	MaxJudgerCount int     `yaml:"maxJudgerCount"`
	Host           string  `yaml:"host"`
	Exe            Exe     `yaml:"exe"`
	Sandbox        Sandbox `yaml:"sandbox"`
	File           File    `yaml:"file"`

	Rabbitmq Rabbitmq `yaml:"rabbitmq"`
	Redis    Redis    `yaml:"redis"`
}

// Sandbox
type Sandbox struct {
	DefaultMaxMemory     int    `yaml:"defaultMaxMemory"`
	DefaultMaxOutputSize int    `yaml:"defaultMaxOutputSize"`
	ExePath              string `yaml:"exePath"`
	LogPath              string `yaml:"logPath"`
	DefaultMaxCpuTime    int    `yaml:"defaultMaxCpuTime"`
	DefaultMaxRealTime   int    `yaml:"defaultMaxRealTime"`
}

// Exe
type Exe struct {
	Javac  string `yaml:"javac"`
	Python string `yaml:"python"`
	Gcc    string `yaml:"gcc"`
	Gpp    string `yaml:"gpp"`
	Go     string `yaml:"go"`
	Java   string `yaml:"java"`
}

// File
type File struct {
	TempPath     string `yaml:"tempPath"`
	CodePath     string `yaml:"codePath"`
	TestcasePath string `yaml:"testcasePath"`
}

// Rabbitmq
type Rabbitmq struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
}

// Redis
type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	LockKey  string `yaml:"lockKey"`
	Ttl      int    `yaml:"ttl"`
	ShortTtl int    `yaml:"shortTtl"`
}
