package config

var Config = new(Default)

// Submit
type Default struct {
	Port     int      `yaml:"port"`
	Name     string   `yaml:"name"`
	Version  string   `yaml:"version"`
	Host     string   `yaml:"host"`
	Jobs     Jobs     `yaml:"jobs"`
	Redis    Redis    `yaml:"redis"`
	Mysql    Mysql    `yaml:"mysql"`
	Rabbitmq Rabbitmq `yaml:"rabbitmq"`
}

// Jobs
type Jobs struct {
	Time int `yaml:"time"`
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

// Rabbitmq
type Rabbitmq struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
}
