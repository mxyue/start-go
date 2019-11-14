package codewrap

import (
	"project/src/utils/lang"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const symbol = "<#>"
const defaultCode = 0

//HTTPCodes http状态码
var HTTPCodes = struct {
	Unauthorized,
	NoPermission int
}{
	401,
	403,
}

//Codes 错误码定义
var Codes = struct {
	ParamsError,
	NotFound,
	SystemError,
	CustomMessage,
	DuplicateBuy int
}{
	400,
	404,
	500,
	601,
	701,
}

//Error code，message 转error
func Error(code int, message string) error {
	if lang.StrIncludeWord(message, symbol) {
		return errors.New(message)
	}
	return fmt.Errorf("%d%s%s", code, symbol, message)
}

//String code，message 转string
func String(code int, message string) string {
	if lang.StrIncludeWord(message, symbol) {
		return message
	}
	return fmt.Sprintf("%d%s%s", code, symbol, message)
}

//Parse 转成code，message
func Parse(message string) (int, string) {
	items := strings.Split(message, symbol)
	if len(items) == 2 {
		code, err := strconv.Atoi(items[0])
		if err != nil {
			return defaultCode, message
		}
		return code, items[1]
	}
	return defaultCode, message
}

//CreateSystemError 快捷创建系统错误信息
func CreateSystemError(message string) error {
	return Error(Codes.SystemError, message)
}

//CreateCustomMessage  快捷创建自定义错误信息
func CreateCustomMessage(message string) error {
	return Error(Codes.CustomMessage, message)
}
