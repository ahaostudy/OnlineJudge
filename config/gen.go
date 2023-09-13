package config

// Redis
type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	LockKey  string `yaml:"lockKey"`
	Ttl      int    `yaml:"ttl"`
}

// Testcase
type Testcase struct {
	Name    string       `yaml:"name"`
	Version string       `yaml:"version"`
	Host    string       `yaml:"host"`
	Port    int          `yaml:"port"`
	File    TestcaseFile `yaml:"file"`
}

// TestcaseFile
type TestcaseFile struct {
	Path string `yaml:"path"`
}

// Etcd
type Etcd struct {
	Addr string `yaml:"addr"`
	Ttl  int    `yaml:"ttl"`
}

// Default
type Default struct {
	Rabbitmq Rabbitmq `yaml:"rabbitmq"`
	Redis    Redis    `yaml:"redis"`
	Server   Server   `yaml:"server"`
	Judge    Judge    `yaml:"judge"`
	Problem  Problem  `yaml:"problem"`
	Testcase Testcase `yaml:"testcase"`
	Etcd     Etcd     `yaml:"etcd"`
	Private  Private  `yaml:"private"`
	Mysql    Mysql    `yaml:"mysql"`
}

// Server
type Server struct {
	Email Email  `yaml:"email"`
	Auth  Auth   `yaml:"auth"`
	Salt  string `yaml:"salt"`
	Node  int    `yaml:"node"`
}

// Email
type Email struct {
	Addr   string `yaml:"addr"`
	Host   string `yaml:"host"`
	From   string `yaml:"from"`
	Email  string `yaml:"email"`
	Auth   string `yaml:"auth"`
	Expire int    `yaml:"expire"`
}

// System
type System struct {
	SudoPwd string `yaml:"sudoPwd"`
}

// Exe
type Exe struct {
	Java   string `yaml:"java"`
	Javac  string `yaml:"javac"`
	Python string `yaml:"python"`
	Gcc    string `yaml:"gcc"`
	Gpp    string `yaml:"gpp"`
	Go     string `yaml:"go"`
}

// Sandbox
type Sandbox struct {
	ExePath              string `yaml:"exePath"`
	LogPath              string `yaml:"logPath"`
	DefaultMaxTime       int    `yaml:"defaultMaxTime"`
	DefaultMaxMemory     int    `yaml:"defaultMaxMemory"`
	DefaultMaxOutputSize int    `yaml:"defaultMaxOutputSize"`
}

// Problem
type Problem struct {
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Name    string `yaml:"name"`
}

// Private
type Private struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
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

// Rabbitmq
type Rabbitmq struct {
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
}

// Auth
type Auth struct {
	Key    string `yaml:"key"`
	Expire int    `yaml:"expire"`
	Issuer string `yaml:"issuer"`
}

// Judge
type Judge struct {
	Sandbox Sandbox `yaml:"sandbox"`
	File    File    `yaml:"file"`
	Name    string  `yaml:"name"`
	Version string  `yaml:"version"`
	Host    string  `yaml:"host"`
	Port    int     `yaml:"port"`
	System  System  `yaml:"system"`
	Exe     Exe     `yaml:"exe"`
}

// File
type File struct {
	DemoPath string `yaml:"demoPath"`
	TempPath string `yaml:"tempPath"`
}

