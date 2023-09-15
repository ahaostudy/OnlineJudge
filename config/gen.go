package config

// TestcaseFile
type TestcaseFile struct {
	Path string `yaml:"path"`
}

// File
type File struct {
	CodePath string `yaml:"codePath"`
	DemoPath string `yaml:"demoPath"`
	TempPath string `yaml:"tempPath"`
}

// Private
type Private struct {
	Port    int    `yaml:"port"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
}

// Server
type Server struct {
	Node  int    `yaml:"node"`
	Email Email  `yaml:"email"`
	Auth  Auth   `yaml:"auth"`
	Salt  string `yaml:"salt"`
}

// Rabbitmq
type Rabbitmq struct {
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
}

// Redis
type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	LockKey  string `yaml:"lockKey"`
	Ttl      int    `yaml:"ttl"`
	ShortTtl int    `yaml:"shortTtl"`
}

// System
type System struct {
	SudoPwd string `yaml:"sudoPwd"`
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

// Sandbox
type Sandbox struct {
	ExePath              string `yaml:"exePath"`
	LogPath              string `yaml:"logPath"`
	DefaultMaxTime       int    `yaml:"defaultMaxTime"`
	DefaultMaxMemory     int    `yaml:"defaultMaxMemory"`
	DefaultMaxOutputSize int    `yaml:"defaultMaxOutputSize"`
}

// Testcase
type Testcase struct {
	Host    string       `yaml:"host"`
	Port    int          `yaml:"port"`
	File    TestcaseFile `yaml:"file"`
	Name    string       `yaml:"name"`
	Version string       `yaml:"version"`
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

// Default
type Default struct {
	Redis    Redis    `yaml:"redis"`
	Judge    Judge    `yaml:"judge"`
	Testcase Testcase `yaml:"testcase"`
	Submit   Submit   `yaml:"submit"`
	Private  Private  `yaml:"private"`
	Etcd     Etcd     `yaml:"etcd"`
	Mysql    Mysql    `yaml:"mysql"`
	Server   Server   `yaml:"server"`
	Problem  Problem  `yaml:"problem"`
	Rabbitmq Rabbitmq `yaml:"rabbitmq"`
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

// Email
type Email struct {
	Email  string `yaml:"email"`
	Auth   string `yaml:"auth"`
	Expire int    `yaml:"expire"`
	Addr   string `yaml:"addr"`
	Host   string `yaml:"host"`
	From   string `yaml:"from"`
}

// Auth
type Auth struct {
	Issuer string `yaml:"issuer"`
	Key    string `yaml:"key"`
	Expire int    `yaml:"expire"`
}

// Problem
type Problem struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

// Submit
type Submit struct {
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

