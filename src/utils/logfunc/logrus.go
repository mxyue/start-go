package logfunc

import (
	"project/src/config"
	"project/src/utils/file"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var logPath = fmt.Sprintf("%s/temp/logs", config.RootPath)

var logLevels = map[string]logrus.Level{
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
	"error": logrus.ErrorLevel,
	"warn":  logrus.WarnLevel,
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
	"trace": logrus.TraceLevel,
}

func init() {
	if config.EnvIsTest() {
		return
	}
	file.Mkdir(logPath)
	level := config.GetLogLevel()
	logrus.AddHook(new(LineHook))
	logrus.AddHook(allOutHook(level, 20))
	logrus.AddHook(errOutHook(20))
	setErrorLog()
}

func setErrorLog() {
	logErr := io.MultiWriter(logrus.StandardLogger().WriterLevel(logrus.ErrorLevel), io.Writer(os.Stderr))
	gin.DefaultErrorWriter = logErr
}

func allOutHook(logLevel string, maxRemainCnt uint) logrus.Hook {
	stdLogName := fmt.Sprintf("%s/std", logPath)
	writer, err := rotatelogs.New(
		stdLogName+".%Y%m%d%H",
		// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
		rotatelogs.WithLinkName(stdLogName),

		// WithRotationTime设置日志分割的时间,这里设置为一天分割一次
		rotatelogs.WithRotationTime(24*time.Hour),

		// WithMaxAge和WithRotationCount二者只能设置一个,
		// WithMaxAge设置文件清理前的最长保存时间,
		// WithRotationCount设置文件清理前最多保存的个数.
		//rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationCount(maxRemainCnt),
	)

	if err != nil {
		logrus.Errorf("[allOutHook] config local file system for logger error: %v", err)
	}

	level, ok := logLevels[logLevel]

	if ok {
		logrus.SetLevel(level)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{})

	if !config.EnvIsDev() {
		logrus.StandardLogger().Out = ioutil.Discard
	}
	return lfsHook
}

func errOutHook(maxRemainCnt uint) logrus.Hook {
	errLogName := fmt.Sprintf("%s/err", logPath)
	errWriter, err := rotatelogs.New(
		errLogName+".%Y%m%d%H",
		rotatelogs.WithRotationTime(72*time.Hour),
		rotatelogs.WithRotationCount(maxRemainCnt),
	)

	if err != nil {
		logrus.Errorf("[errOutHook] config local file system for logger error: %v", err)
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.ErrorLevel: errWriter,
		logrus.FatalLevel: errWriter,
		logrus.PanicLevel: errWriter,
	}, &logrus.TextFormatter{})

	return lfsHook
}
