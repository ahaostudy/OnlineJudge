package config

// Redis
type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	LockKey  string `yaml:"lockKey"`
	Ttl      int    `yaml:"ttl"`
	ShortTtl int    `yaml:"shortTtl"`
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

// Chatgpt
type Chatgpt struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Openai  Openai `yaml:"openai"`
}

// Private
type Private struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

// Default
type Default struct {
	Testcase Testcase `yaml:"testcase"`
	Contest  Contest  `yaml:"contest"`
	Problem  Problem  `yaml:"problem"`
	Etcd     Etcd     `yaml:"etcd"`
	Rabbitmq Rabbitmq `yaml:"rabbitmq"`
	Mysql    Mysql    `yaml:"mysql"`
	Redis    Redis    `yaml:"redis"`
	Mongodb  Mongodb  `yaml:"mongodb"`
	Server   Server   `yaml:"server"`
	Judge    Judge    `yaml:"judge"`
	Submit   Submit   `yaml:"submit"`
	Chatgpt  Chatgpt  `yaml:"chatgpt"`
	Private  Private  `yaml:"private"`
}

// Mysql
type Mysql struct {
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
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

// Testcase
type Testcase struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	File    File   `yaml:"file"`
}

// Sandbox
type Sandbox struct {
	DefaultMaxTime       int    `yaml:"defaultMaxTime"`
	DefaultMaxMemory     int    `yaml:"defaultMaxMemory"`
	DefaultMaxOutputSize int    `yaml:"defaultMaxOutputSize"`
	ExePath              string `yaml:"exePath"`
	LogPath              string `yaml:"logPath"`
}

// Rabbitmq
type Rabbitmq struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
}

// JudgeFile
type JudgeFile struct {
	TempPath string `yaml:"tempPath"`
	CodePath string `yaml:"codePath"`
	DemoPath string `yaml:"demoPath"`
}

// System
type System struct {
	SudoPwd string `yaml:"sudoPwd"`
}

// Openai
type Openai struct {
	BaseUrl string `yaml:"baseUrl"`
	ApiKey  string `yaml:"apiKey"`
	Model   string `yaml:"model"`
}

// Problem
type Problem struct {
	Port    int    `yaml:"port"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
}

// Etcd
type Etcd struct {
	Addr string `yaml:"addr"`
	Ttl  int    `yaml:"ttl"`
}

// Mongodb
type Mongodb struct {
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
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

// Judge
type Judge struct {
	Sandbox Sandbox   `yaml:"sandbox"`
	File    JudgeFile `yaml:"file"`
	Name    string    `yaml:"name"`
	Version string    `yaml:"version"`
	Host    string    `yaml:"host"`
	Port    int       `yaml:"port"`
	System  System    `yaml:"system"`
	Exe     Exe       `yaml:"exe"`
}

// Submit
type Submit struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

// File
type File struct {
	Path string `yaml:"path"`
}

// Contest
type Contest struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

