package logfunc

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

//LineHook 显示日志的文件和行数,正式使用时，不建议使用
type LineHook struct{}

//Levels logrus接口规范
func (hook LineHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

//Fire logrus接口规范
func (hook LineHook) Fire(entry *logrus.Entry) error {
	pathLine := findCaller(7)
	entry.Data["logLocation"] = pathLine
	return nil
}

func findCaller(skip int) string {
	file := ""
	line := 0
	for i := 0; i < 10; i++ {
		file, line = findSkip(skip + i)
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func findSkip(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0
	}
	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line
}
