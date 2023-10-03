package config

// Mysql
type Mysql struct {
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

// Problem
type Problem struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	File    File   `yaml:"file"`
}

// File
type File struct {
	TestcasePath string `yaml:"testcasePath"`
}

// User
type User struct {
	Email   Email  `yaml:"email"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

// Private
type Private struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

// Rabbitmq
type Rabbitmq struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
}

// System
type System struct {
	SudoPwd string `yaml:"sudoPwd"`
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

// Submit
type Submit struct {
	Port    int    `yaml:"port"`
	Jobs    Jobs   `yaml:"jobs"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
}

// Mongodb
type Mongodb struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
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
	Name           string    `yaml:"name"`
	Version        string    `yaml:"version"`
	Port           int       `yaml:"port"`
	MaxJudgerCount int       `yaml:"maxJudgerCount"`
	Host           string    `yaml:"host"`
	System         System    `yaml:"system"`
	Exe            Exe       `yaml:"exe"`
	Sandbox        Sandbox   `yaml:"sandbox"`
	File           JudgeFile `yaml:"file"`
}

// Default
type Default struct {
	Etcd     Etcd     `yaml:"etcd"`
	Mysql    Mysql    `yaml:"mysql"`
	Mongodb  Mongodb  `yaml:"mongodb"`
	Problem  Problem  `yaml:"problem"`
	Chatgpt  Chatgpt  `yaml:"chatgpt"`
	Auth     Auth     `yaml:"auth"`
	User     User     `yaml:"user"`
	Private  Private  `yaml:"private"`
	Jaeger   Jaeger   `yaml:"jaeger"`
	Rabbitmq Rabbitmq `yaml:"rabbitmq"`
	Redis    Redis    `yaml:"redis"`
	Judge    Judge    `yaml:"judge"`
	Submit   Submit   `yaml:"submit"`
	Contest  Contest  `yaml:"contest"`
}

// Openai
type Openai struct {
	BaseUrl string `yaml:"baseUrl"`
	ApiKey  string `yaml:"apiKey"`
	Model   string `yaml:"model"`
}

// Contest
type Contest struct {
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Name    string `yaml:"name"`
}

// Etcd
type Etcd struct {
	Addr string `yaml:"addr"`
	Ttl  int    `yaml:"ttl"`
}

// Chatgpt
type Chatgpt struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Openai  Openai `yaml:"openai"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
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

// Jaeger
type Jaeger struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

// Redis
type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	LockKey  string `yaml:"lockKey"`
	Ttl      int    `yaml:"ttl"`
	ShortTtl int    `yaml:"shortTtl"`
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

// JudgeFile
type JudgeFile struct {
	TempPath string `yaml:"tempPath"`
	CodePath string `yaml:"codePath"`
	DemoPath string `yaml:"demoPath"`
}

// Jobs
type Jobs struct {
	Time int `yaml:"time"`
}
