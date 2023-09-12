package config

// Judge
type Judge struct {
	System  System  `yaml:"system"`
	Exe     Exe     `yaml:"exe"`
	Sandbox Sandbox `yaml:"sandbox"`
	File    File    `yaml:"file"`
	Name    string  `yaml:"name"`
	Version string  `yaml:"version"`
	Host    string  `yaml:"host"`
	Port    int     `yaml:"port"`
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

// Default
type Default struct {
	Redis    Redis    `yaml:"redis"`
	Server   Server   `yaml:"server"`
	Judge    Judge    `yaml:"judge"`
	Problem  Problem  `yaml:"problem"`
	Etcd     Etcd     `yaml:"etcd"`
	Rabbitmq Rabbitmq `yaml:"rabbitmq"`
	Mysql    Mysql    `yaml:"mysql"`
}

// Server
type Server struct {
	Salt  string `yaml:"salt"`
	Node  int    `yaml:"node"`
	Email Email  `yaml:"email"`
	Auth  Auth   `yaml:"auth"`
}

// Auth
type Auth struct {
	Issuer string `yaml:"issuer"`
	Key    string `yaml:"key"`
	Expire int    `yaml:"expire"`
}

// Mysql
type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

// Email
type Email struct {
	Auth   string `yaml:"auth"`
	Expire int    `yaml:"expire"`
	Addr   string `yaml:"addr"`
	Host   string `yaml:"host"`
	From   string `yaml:"from"`
	Email  string `yaml:"email"`
}

// Sandbox
type Sandbox struct {
	DefaultMaxMemory     int    `yaml:"defaultMaxMemory"`
	DefaultMaxOutputSize int    `yaml:"defaultMaxOutputSize"`
	ExePath              string `yaml:"exePath"`
	LogPath              string `yaml:"logPath"`
	DefaultMaxTime       int    `yaml:"defaultMaxTime"`
}

// Rabbitmq
type Rabbitmq struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

// Exe
type Exe struct {
	Gcc    string `yaml:"gcc"`
	Gpp    string `yaml:"gpp"`
	Go     string `yaml:"go"`
	Java   string `yaml:"java"`
	Javac  string `yaml:"javac"`
	Python string `yaml:"python"`
}

// Redis
type Redis struct {
	Ttl      int    `yaml:"ttl"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	LockKey  string `yaml:"lockKey"`
}

// Problem
type Problem struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

// Etcd
type Etcd struct {
	Addr string `yaml:"addr"`
	Ttl  int    `yaml:"ttl"`
}

