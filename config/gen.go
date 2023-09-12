package config

// Server
type Server struct {
	Salt  string `yaml:"salt"`
	Node  int    `yaml:"node"`
	Email Email  `yaml:"email"`
	Auth  Auth   `yaml:"auth"`
}

// Default
type Default struct {
	Judge    Judge    `yaml:"judge"`
	Etcd     Etcd     `yaml:"etcd"`
	Rabbitmq Rabbitmq `yaml:"rabbitmq"`
	Mysql    Mysql    `yaml:"mysql"`
	Redis    Redis    `yaml:"redis"`
	Server   Server   `yaml:"server"`
}

// Exe
type Exe struct {
	Python string `yaml:"python"`
	Gcc    string `yaml:"gcc"`
	Gpp    string `yaml:"gpp"`
	Go     string `yaml:"go"`
	Java   string `yaml:"java"`
	Javac  string `yaml:"javac"`
}

// Sandbox
type Sandbox struct {
	ExePath              string `yaml:"exePath"`
	LogPath              string `yaml:"logPath"`
	DefaultMaxTime       int    `yaml:"defaultMaxTime"`
	DefaultMaxMemory     int    `yaml:"defaultMaxMemory"`
	DefaultMaxOutputSize int    `yaml:"defaultMaxOutputSize"`
}

// Mysql
type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
}

// Redis
type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	LockKey  string `yaml:"lockKey"`
	Ttl      int    `yaml:"ttl"`
}

// Email
type Email struct {
	Expire int    `yaml:"expire"`
	Addr   string `yaml:"addr"`
	Host   string `yaml:"host"`
	From   string `yaml:"from"`
	Email  string `yaml:"email"`
	Auth   string `yaml:"auth"`
}

// Auth
type Auth struct {
	Issuer string `yaml:"issuer"`
	Key    string `yaml:"key"`
	Expire int    `yaml:"expire"`
}

// Judge
type Judge struct {
	Host    string  `yaml:"host"`
	Port    int     `yaml:"port"`
	System  System  `yaml:"system"`
	Exe     Exe     `yaml:"exe"`
	Sandbox Sandbox `yaml:"sandbox"`
	File    File    `yaml:"file"`
	Name    string  `yaml:"name"`
	Version string  `yaml:"version"`
}

// System
type System struct {
	SudoPwd string `yaml:"sudoPwd"`
}

// File
type File struct {
	TempPath string `yaml:"tempPath"`
	DemoPath string `yaml:"demoPath"`
}

// Etcd
type Etcd struct {
	Addr string `yaml:"addr"`
	Ttl  int    `yaml:"ttl"`
}

// Rabbitmq
type Rabbitmq struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
}

