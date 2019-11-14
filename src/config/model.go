package config

// Mongo MongoDB服务器配置
type Mongo struct {
	Hosts  string
	DBName string
}

//Redis Redis服务器配置
type Redis struct {
	Host     string
	Port     uint
	Password string
}

//AliOSS 阿里云配置
type AliOSS struct {
	Prefix string
}

//Auth 登录校验
type Auth struct {
	JwtPrivateKey string
}

//Configuration <<first>> 本服务配置
type Configuration struct {
	Host    string
	Port    uint
	Mongo   Mongo
	Redis   Redis
	PidPath string
	Auth    Auth
}

//Arguments 命令行参数
type Arguments struct {
	mode     string
	env      string
	logLevel string
	Host     string
	Port     uint
	Version  bool
}
