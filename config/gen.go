package config

// Redis
type Redis struct {
	ShortTtl int    `yaml:"shortTtl"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	LockKey  string `yaml:"lockKey"`
	Ttl      int    `yaml:"ttl"`
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
	Name    string    `yaml:"name"`
	Version string    `yaml:"version"`
	Host    string    `yaml:"host"`
	Port    int       `yaml:"port"`
	System  System    `yaml:"system"`
	Exe     Exe       `yaml:"exe"`
	Sandbox Sandbox   `yaml:"sandbox"`
	File    JudgeFile `yaml:"file"`
}

// System
type System struct {
	SudoPwd string `yaml:"sudoPwd"`
}

// Sandbox
type Sandbox struct {
	ExePath              string `yaml:"exePath"`
	LogPath              string `yaml:"logPath"`
	DefaultMaxTime       int    `yaml:"defaultMaxTime"`
	DefaultMaxMemory     int    `yaml:"defaultMaxMemory"`
	DefaultMaxOutputSize int    `yaml:"defaultMaxOutputSize"`
}

// Jobs
type Jobs struct {
	Time int `yaml:"time"`
}

// Contest
type Contest struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

// Openai
type Openai struct {
	ApiKey  string `yaml:"apiKey"`
	Model   string `yaml:"model"`
	BaseUrl string `yaml:"baseUrl"`
}

// Chatgpt
type Chatgpt struct {
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Openai  Openai `yaml:"openai"`
	Name    string `yaml:"name"`
}

// File
type File struct {
	TestcasePath string `yaml:"testcasePath"`
}

// Rabbitmq
type Rabbitmq struct {
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
}

// JudgeFile
type JudgeFile struct {
	DemoPath string `yaml:"demoPath"`
	TempPath string `yaml:"tempPath"`
	CodePath string `yaml:"codePath"`
}

// Submit
type Submit struct {
	Jobs    Jobs   `yaml:"jobs"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

// Default
type Default struct {
	Etcd     Etcd     `yaml:"etcd"`
	Mysql    Mysql    `yaml:"mysql"`
	Mongodb  Mongodb  `yaml:"mongodb"`
	Problem  Problem  `yaml:"problem"`
	Contest  Contest  `yaml:"contest"`
	Private  Private  `yaml:"private"`
	Rabbitmq Rabbitmq `yaml:"rabbitmq"`
	Redis    Redis    `yaml:"redis"`
	Server   Server   `yaml:"server"`
	Judge    Judge    `yaml:"judge"`
	Submit   Submit   `yaml:"submit"`
	Chatgpt  Chatgpt  `yaml:"chatgpt"`
}

// Email
type Email struct {
	From   string `yaml:"from"`
	Email  string `yaml:"email"`
	Auth   string `yaml:"auth"`
	Expire int    `yaml:"expire"`
	Addr   string `yaml:"addr"`
	Host   string `yaml:"host"`
}

// Mysql
type Mysql struct {
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
}

// Mongodb
type Mongodb struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Problem
type Problem struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	File    File   `yaml:"file"`
}

// Private
type Private struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
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

// Etcd
type Etcd struct {
	Addr string `yaml:"addr"`
	Ttl  int    `yaml:"ttl"`
}

