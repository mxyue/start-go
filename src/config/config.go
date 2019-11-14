package config

import (
	"github.com/imdario/mergo"
)

var (
	//mode 开发模式
	mode = "debug"
	//env 运行环境
	env = "dev"
	//args 启动参数
	args = parseArgs()
	//BuildTime 编译时间
	BuildTime = ""
	//Version 版本
	Version = ""
	//GitHash Git提交hash码
	GitHash = ""
	//Config 配置信息
	Config Configuration
	//RootPath 当前可执行文件所在路径
	RootPath = getRootPath()
)

var defaultMongoHost = getDefaultMongoHost()
var defaultRedisHost, defaultRedisPort = getDefaultRedisConf()

/// 初始配置，其他配置根据不同的env环境进行替换 /////  ===================================
var defaultConfig = Configuration{
	Host: "0.0.0.0",
	Port: 9500,

	Mongo: Mongo{
		Hosts:  defaultMongoHost,
		DBName: "my-database-dev",
	},

	Redis: Redis{
		Host:     defaultRedisHost,
		Port:     defaultRedisPort,
		Password: "",
	},
}

/// 测试环境 /// ====================================================================
func getTestConfig() Configuration {
	return Configuration{
		Mongo: Mongo{
			DBName: "my-database-test",
		},
		Auth: Auth{JwtPrivateKey: "dev-private-key"},
	}
}

/// 本地开发环境 /// ================================================================
func getDevConfig() Configuration {
	return defaultConfig
}

/// 外网测试环境配置 /////============================================================
var stageConfig = Configuration{
	Mongo: Mongo{
		Hosts:  "localhost:27017",
		DBName: "my-database-dev",
	},
}

func getStageConfig() Configuration {
	doc := Configuration{}
	_ = mergo.Merge(&doc, stageConfig)
	return doc
}

/// 正式环境配置 /////================================================================
var productionConfig = Configuration{
	Mongo: Mongo{
		Hosts:  "127.0.0.1:27017,127.0.0.1:27018",
		DBName: "my-database",
	},

	Redis: Redis{
		Host: "127.0.0.1",
	},
}

func getProdConfig() Configuration {
	productionConfig.Host = "127.0.0.1"
	return productionConfig
}
