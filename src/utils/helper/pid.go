package helper

import (
	"project/src/config"
	"project/src/utils/file"
	"fmt"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"
)

//WritePid 写入pid文件
func WritePid() {
	pid := syscall.Getpid()
	logrus.Info("Actual pid is: ", pid)
	pidDirPath := strings.Trim(config.Config.PidPath, fmt.Sprintf("%d", config.Config.Port))
	file.Mkdir(pidDirPath)
	file.Overwrite(config.Config.PidPath, fmt.Sprintf("%d", pid))
}
