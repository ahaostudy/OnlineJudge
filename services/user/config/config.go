package config

var Config = new(Default)

type Default struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Static  Static `yaml:"static"`
	Auth    Auth   `yaml:"auth"`
	Email   Email  `yaml:"email"`
	Mysql   Mysql  `yaml:"mysql"`
	Redis   Redis  `yaml:"redis"`
}

type Static struct {
	Path string `yaml:"path"`
}

// Auth
type Auth struct {
	Issuer string `yaml:"issuer"`
	Key    string `yaml:"key"`
	Expire int    `yaml:"expire"`
	Salt   string `yaml:"salt"`
	Node   int    `yaml:"node"`
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

// Redis
type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	LockKey  string `yaml:"lockKey"`
	Ttl      int    `yaml:"ttl"`
	ShortTtl int    `yaml:"shortTtl"`
}

// Mysql
type Mysql struct {
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}
