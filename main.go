package main

import (
	"project/src/config"
	_ "project/src/db/indexes"
	"project/src/routes"
	"project/src/utils/helper"
	_ "project/src/utils/logfunc"
	"fmt"
	"io/ioutil"
	"strconv"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/sirupsen/logrus"
)


func init() {
	helper.WritePid()
}

func beforeStop() {
	str, _ := ioutil.ReadFile(config.Config.PidPath)
	pid, _ := strconv.ParseInt(string(str), 10, 64)
	fmt.Println("current pid>", syscall.Getpid(), " file pid>", int(pid))
	if syscall.Getpid() == int(pid) {
		fmt.Println("go to unregister")
	}
}

/**
Swagger 统一配置
@title web后台API接口
@version 0.0.2
@BasePath /api
@securityDefinitions.apikey appKey
@in header
@name Authorization
**/
func main() {
	// Endless
	router := routes.Router()

	host := fmt.Sprintf("%s:%d", config.Config.Host, config.Config.Port)
	server := endless.NewServer(host, router)
	server.ReadTimeout = 5 * time.Second

	server.BeforeBegin = func(add string) {
		//consulManager.Register()
	}

	for _, signal := range []syscall.Signal{syscall.SIGINT, syscall.SIGTERM} {
		err := server.RegisterSignalHook(endless.PRE_SIGNAL, signal, func() {
			beforeStop()
		})
		if err != nil {
			logrus.Error("register signal err:", err)
		}
	}

	if err := server.ListenAndServe(); err != nil {
		logrus.Warn("main.Server err: ", err)
	}
}
