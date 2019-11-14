package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/imdario/mergo"
)

func init() {
	initConfig(args, &Config, &mode, &env)
	showVersion()
	gin.SetMode(mode)
	printConfigInfo()
}

//envs 环境变量值
var envs = struct {
	Test, Dev, Stage, Prod string
}{
	"test", "dev", "stage", "prod",
}
var ginModel = []string{gin.TestMode, gin.DebugMode, gin.ReleaseMode}

func parseArgs() Arguments {
	mode := flag.String(
		"mode",
		"",
		fmt.Sprintf("--mode=[%s]", strings.Join(ginModel, "|")),
	)
	env := flag.String(
		"env",
		"",
		fmt.Sprintf("--env=[%s]", strings.Join(getConfigKeys(envConfigs), "|")),
	)
	host := flag.String(
		"host",
		"",
		"--host=localhost",
	)
	port := flag.Uint(
		"port",
		0,
		"--port=8001",
	)
	logLevel := flag.String(
		"logLevel",
		"info",
		"--logLevel=debug",
	)
	version := flag.Bool(
		"version",
		false,
		"Show version information",
	)

	flag.Parse()
	//checkArg(mode, modeConfigs)
	checkArg(env, envConfigs)
	return Arguments{
		mode:     *mode,
		env:      *env,
		Host:     *host,
		Port:     *port,
		logLevel: *logLevel,
		Version:  *version,
	}
}

func checkArg(arg *string, configMap map[string]Configuration) {
	_, ok := configMap[*arg]
	if len(*arg) != 0 && !ok {
		panic(fmt.Sprintf("Error: The arg %s is error", *arg))
	}
}

func coverString(defaultValue string, values ...string) string {
	value := defaultValue
	for _, item := range values {
		if len(item) != 0 {
			value = item
		}
	}
	return value
}

func getTestEnv() string {
	if flag.Lookup("test.v") != nil {
		return "test"
	}
	return ""
}

var envConfigs = map[string]Configuration{
	"test":  getTestConfig(),
	"dev":   getDevConfig(),
	"stage": getStageConfig(),
	"prod":  getProdConfig(),
}

func initConfig(args Arguments, config *Configuration, mode *string, env *string) {
	*mode = coverString(*mode, getTestEnv(), args.mode)
	*env = coverString(*env, getTestEnv(), args.env)
	argsConfiguration := Configuration{
		Host: args.Host,
		Port: args.Port,
	}

	_ = mergo.Merge(config, defaultConfig, mergo.WithOverride)
	_ = mergo.Merge(config, envConfigs[*env], mergo.WithOverride)
	_ = mergo.Merge(config, argsConfiguration, mergo.WithOverride)
	configAfterBuild(config)

}

// 获取当前可执行文件所在路径
func getRootPath() string {
	if BuildTime == "" {
		return "."
	}
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func configAfterBuild(config *Configuration) {
	config.PidPath = getPidPath(config)
}

func getPidPath(config *Configuration) string {
	rootPath := getRootPath()
	return fmt.Sprintf("%s/temp/pid/%d", rootPath, config.Port)
}

// 获取一个配置集的 key
func getConfigKeys(mapper map[string]Configuration) []string {
	keys := make([]string, 0)
	for key := range mapper {
		keys = append(keys, key)
	}
	return keys
}

func getHostFromEnv(name string) string {
	env := os.Getenv(name)
	part := strings.Split(env, "://")
	if len(part) != 2 {
		return ""
	}
	return part[1]
}

func getDefaultMongoHost() string {
	mongoHost := getHostFromEnv("MONGO_PORT")
	if mongoHost != "" {
		fmt.Println("mongo custom host:", mongoHost)
		return mongoHost
	}
	return "localhost:27017"
}

func getDefaultRedisConf() (string, uint) {
	redis := getHostFromEnv("REDIS_PORT")
	if redis == "" {
		return "127.0.0.1", uint(6379)
	}

	redisPart := strings.Split(redis, ":")
	redisHost := redisPart[0]
	redisPort, _ := strconv.Atoi(redisPart[1])
	fmt.Printf("redis custom host %s, port %d", redisHost, redisPort)
	return redisHost, uint(redisPort)
}
