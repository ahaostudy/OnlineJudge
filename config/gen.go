package config

// Rabbitmq
type Rabbitmq struct {
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
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

// Problem
type Problem struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	File    File   `yaml:"file"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

// Redis
type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	LockKey  string `yaml:"lockKey"`
	Ttl      int    `yaml:"ttl"`
	ShortTtl int    `yaml:"shortTtl"`
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
	Sandbox Sandbox   `yaml:"sandbox"`
	File    JudgeFile `yaml:"file"`
	Name    string    `yaml:"name"`
	Version string    `yaml:"version"`
	Host    string    `yaml:"host"`
	Port    int       `yaml:"port"`
	System  System    `yaml:"system"`
	Exe     Exe       `yaml:"exe"`
}

// Sandbox
type Sandbox struct {
	ExePath              string `yaml:"exePath"`
	LogPath              string `yaml:"logPath"`
	DefaultMaxTime       int    `yaml:"defaultMaxTime"`
	DefaultMaxMemory     int    `yaml:"defaultMaxMemory"`
	DefaultMaxOutputSize int    `yaml:"defaultMaxOutputSize"`
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

// Private
type Private struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

// Auth
type Auth struct {
	Issuer string `yaml:"issuer"`
	Key    string `yaml:"key"`
	Expire int    `yaml:"expire"`
	Salt   string `yaml:"salt"`
	Node   int    `yaml:"node"`
}

// Default
type Default struct {
	Problem  Problem  `yaml:"problem"`
	Submit   Submit   `yaml:"submit"`
	Chatgpt  Chatgpt  `yaml:"chatgpt"`
	Contest  Contest  `yaml:"contest"`
	Etcd     Etcd     `yaml:"etcd"`
	Redis    Redis    `yaml:"redis"`
	Mongodb  Mongodb  `yaml:"mongodb"`
	Judge    Judge    `yaml:"judge"`
	User     User     `yaml:"user"`
	Private  Private  `yaml:"private"`
	Rabbitmq Rabbitmq `yaml:"rabbitmq"`
	Mysql    Mysql    `yaml:"mysql"`
	Auth     Auth     `yaml:"auth"`
}

// Submit
type Submit struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Jobs    Jobs   `yaml:"jobs"`
}

// Openai
type Openai struct {
	BaseUrl string `yaml:"baseUrl"`
	ApiKey  string `yaml:"apiKey"`
	Model   string `yaml:"model"`
}

// Etcd
type Etcd struct {
	Ttl  int    `yaml:"ttl"`
	Addr string `yaml:"addr"`
}

// System
type System struct {
	SudoPwd string `yaml:"sudoPwd"`
}

// User
type User struct {
	Email   Email  `yaml:"email"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
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

// File
type File struct {
	TestcasePath string `yaml:"testcasePath"`
}

// Jobs
type Jobs struct {
	Time int `yaml:"time"`
}

// Chatgpt
type Chatgpt struct {
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Openai  Openai `yaml:"openai"`
	Name    string `yaml:"name"`
}

// Contest
type Contest struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

// JudgeFile
type JudgeFile struct {
	TempPath string `yaml:"tempPath"`
	CodePath string `yaml:"codePath"`
	DemoPath string `yaml:"demoPath"`
}

