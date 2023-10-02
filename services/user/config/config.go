package config

var Config = new(Default)

type Default struct {
	Email   Email  `yaml:"email"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Auth    Auth   `yaml:"auth"`
	Mysql   Mysql  `yaml:"mysql"`
	Redis   Redis  `yaml:"redis"`
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
