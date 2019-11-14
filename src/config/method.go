package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

//EnvIsDev 环境是否为dev环境
func EnvIsDev() bool {
	return env == envs.Dev
}

//EnvIsTest 是否为测试环境
func EnvIsTest() bool {
	return env == envs.Test
}

//EnvIsProd 是否为正式环境
func EnvIsProd() bool {
	return env == envs.Prod
}

//GetLogLevel 获取日志等级
func GetLogLevel() string {
	if EnvIsDev() {
		return "debug"
	}
	return args.logLevel
}

func showVersion() {
	// Output version info
	if args.Version {
		printVersionInfo()
		os.Exit(0)
	}
}

//printVersionInfo 打印版本信息
func printVersionInfo() {
	fmt.Printf("version:\t%s\n", Version)
	fmt.Printf("git hash:\t%s\n", GitHash)
	fmt.Printf("build time:\t%s\n", BuildTime)
}

//printConfigInfo 打印配置信息
func printConfigInfo() {
	// Output base information
	logrus.Info("Run mode is ", mode)
	logrus.Info("Run environment is ", env)
	logrus.Info("Log level is ", GetLogLevel())

	cf := Config
	logrus.WithFields(logrus.Fields{
		"Host": cf.Host,
		"Port": cf.Port,
	}).Info("Basic config:")
	logrus.WithFields(logrus.Fields{
		"Hosts": cf.Mongo.Hosts,
		"Name":  cf.Mongo.DBName,
	}).Info("Mongodb config:")
}
