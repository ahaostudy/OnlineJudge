package config

// Sandbox
type Sandbox struct {
	DefaultMaxOutputSize int    `yaml:"defaultMaxOutputSize"`
	ExePath              string `yaml:"exePath"`
	LogPath              string `yaml:"logPath"`
	DefaultMaxCpuTime    int    `yaml:"defaultMaxCpuTime"`
	DefaultMaxRealTime   int    `yaml:"defaultMaxRealTime"`
	DefaultMaxMemory     int    `yaml:"defaultMaxMemory"`
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
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

// Chatgpt
type Chatgpt struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Openai  Openai `yaml:"openai"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

// Openai
type Openai struct {
	BaseUrl string `yaml:"baseUrl"`
	ApiKey  string `yaml:"apiKey"`
	Model   string `yaml:"model"`
}

// Etcd
type Etcd struct {
	Addr string `yaml:"addr"`
	Ttl  int    `yaml:"ttl"`
}

// Auth
type Auth struct {
	Issuer string `yaml:"issuer"`
	Key    string `yaml:"key"`
	Expire int    `yaml:"expire"`
	Salt   string `yaml:"salt"`
	Node   int    `yaml:"node"`
}

// Judge
type Judge struct {
	Port           int       `yaml:"port"`
	System         System    `yaml:"system"`
	Exe            Exe       `yaml:"exe"`
	File           JudgeFile `yaml:"file"`
	MaxJudgerCount int       `yaml:"maxJudgerCount"`
	Name           string    `yaml:"name"`
	Version        string    `yaml:"version"`
	Host           string    `yaml:"host"`
	Sandbox        Sandbox   `yaml:"sandbox"`
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

// Problem
type Problem struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	File    File   `yaml:"file"`
}

// Rabbitmq
type Rabbitmq struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
}

// Mongodb
type Mongodb struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// System
type System struct {
	SudoPwd string `yaml:"sudoPwd"`
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

// Default
type Default struct {
	Problem  Problem  `yaml:"problem"`
	Submit   Submit   `yaml:"submit"`
	Chatgpt  Chatgpt  `yaml:"chatgpt"`
	Etcd     Etcd     `yaml:"etcd"`
	Rabbitmq Rabbitmq `yaml:"rabbitmq"`
	Mongodb  Mongodb  `yaml:"mongodb"`
	Auth     Auth     `yaml:"auth"`
	Judge    Judge    `yaml:"judge"`
	User     User     `yaml:"user"`
	Private  Private  `yaml:"private"`
	Mysql    Mysql    `yaml:"mysql"`
	Redis    Redis    `yaml:"redis"`
	Contest  Contest  `yaml:"contest"`
}

// File
type File struct {
	TestcasePath string `yaml:"testcasePath"`
}

// Jobs
type Jobs struct {
	Time int `yaml:"time"`
}

// JudgeFile
type JudgeFile struct {
	TempPath string `yaml:"tempPath"`
	CodePath string `yaml:"codePath"`
	DemoPath string `yaml:"demoPath"`
}

// User
type User struct {
	Port    int    `yaml:"port"`
	Email   Email  `yaml:"email"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
}

// Submit
type Submit struct {
	Port    int    `yaml:"port"`
	Jobs    Jobs   `yaml:"jobs"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
}

// Redis
type Redis struct {
	ShortTtl int    `yaml:"shortTtl"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	LockKey  string `yaml:"lockKey"`
	Ttl      int    `yaml:"ttl"`
}

// Contest
type Contest struct {
	Port    int    `yaml:"port"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
}

