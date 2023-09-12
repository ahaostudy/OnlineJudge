package config

// import "time"

// type config struct {
// 	RabbitMQ rabbitmq `yaml:"rabbitmq"`
// 	Auth     auth     `yaml:"auth"`
// 	MySQL    mysql    `yaml:"mysql"`
// 	Key      key      `yaml:"key"`
// 	Email    email    `yaml:"email"`
// 	Redis    redis    `yaml:"redis"`
// }

// type rabbitmq struct {
// 	Host     string `yaml:"host"`
// 	Port     int    `yaml:"port"`
// 	Username string `yaml:"username"`
// 	Password string `yaml:"password"`
// 	Vhost    string `yaml:"vhost"`
// }

// type auth struct {
// 	Issuer string `yaml:"issuer"`
// 	Key    string `yaml:"key"`
// 	Expire int64  `yaml:"expire"`
// }

// type key struct {
// 	Salt string `yaml:"salt"`
// 	Node int64  `yaml:"node"`
// }

// type mysql struct {
// 	Host     string `yaml:"host"`
// 	Port     string `yaml:"port"`
// 	DBName   string `yaml:"dbname"`
// 	Username string `yaml:"username"`
// 	Password string `yaml:"password"`
// 	Charset  string `yaml:"charset"`
// }

// type email struct {
// 	Addr   string `yaml:"addr"`
// 	Host   string `yaml:"host"`
// 	From   string `yaml:"from"`
// 	Email  string `yaml:"email"`
// 	Auth   string `yaml:"auth"`
// 	Expire int64  `yaml:"expire"`
// }

// type redis struct {
// 	Addr     string `yaml:"addr"`
// 	Password string `yaml:"password"`
// 	KeyLock  string
// 	TTL      time.Duration
// }
