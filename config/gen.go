package config

// Openai
type Openai struct {
	BaseUrl string `yaml:"baseUrl"`
	ApiKey  string `yaml:"apiKey"`
	Model   string `yaml:"model"`
}

// Problem
type Problem struct {
	Name    string      `yaml:"name"`
	Version string      `yaml:"version"`
	Host    string      `yaml:"host"`
	Port    int         `yaml:"port"`
	File    ProblemFile `yaml:"file"`
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
	Judge    Judge    `yaml:"judge"`
	Submit   Submit   `yaml:"submit"`
	Chatgpt  Chatgpt  `yaml:"chatgpt"`
	Contest  Contest  `yaml:"contest"`
	Redis    Redis    `yaml:"redis"`
	Mongodb  Mongodb  `yaml:"mongodb"`
	Server   Server   `yaml:"server"`
	Problem  Problem  `yaml:"problem"`
	Private  Private  `yaml:"private"`
	Etcd     Etcd     `yaml:"etcd"`
	Rabbitmq Rabbitmq `yaml:"rabbitmq"`
	Mysql    Mysql    `yaml:"mysql"`
}

// File
type File struct {
	CodePath string `yaml:"codePath"`
	DemoPath string `yaml:"demoPath"`
	TempPath string `yaml:"tempPath"`
}

// Redis
type Redis struct {
	ShortTtl int    `yaml:"shortTtl"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	LockKey  string `yaml:"lockKey"`
	Ttl      int    `yaml:"ttl"`
}

// Email
type Email struct {
	Host   string `yaml:"host"`
	From   string `yaml:"from"`
	Email  string `yaml:"email"`
	Auth   string `yaml:"auth"`
	Expire int    `yaml:"expire"`
	Addr   string `yaml:"addr"`
}

// Etcd
type Etcd struct {
	Addr string `yaml:"addr"`
	Ttl  int    `yaml:"ttl"`
}

// Rabbitmq
type Rabbitmq struct {
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
	Host     string `yaml:"host"`
}

// Mongodb
type Mongodb struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Judge
type Judge struct {
	File    File    `yaml:"file"`
	Name    string  `yaml:"name"`
	Version string  `yaml:"version"`
	Host    string  `yaml:"host"`
	Port    int     `yaml:"port"`
	System  System  `yaml:"system"`
	Exe     Exe     `yaml:"exe"`
	Sandbox Sandbox `yaml:"sandbox"`
}

// System
type System struct {
	SudoPwd string `yaml:"sudoPwd"`
}

// Exe
type Exe struct {
	Go     string `yaml:"go"`
	Java   string `yaml:"java"`
	Javac  string `yaml:"javac"`
	Python string `yaml:"python"`
	Gcc    string `yaml:"gcc"`
	Gpp    string `yaml:"gpp"`
}

// Submit
type Submit struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

// Chatgpt
type Chatgpt struct {
	Openai  Openai `yaml:"openai"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

// Contest
type Contest struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

// Sandbox
type Sandbox struct {
	DefaultMaxMemory     int    `yaml:"defaultMaxMemory"`
	DefaultMaxOutputSize int    `yaml:"defaultMaxOutputSize"`
	ExePath              string `yaml:"exePath"`
	LogPath              string `yaml:"logPath"`
	DefaultMaxTime       int    `yaml:"defaultMaxTime"`
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

// ProblemFile
type ProblemFile struct {
	TestcasePath string `yaml:"testcasePath"`
}

// Private
type Private struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

