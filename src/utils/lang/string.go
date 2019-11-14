package lang

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//RandomStr 生成指定长度的随机字符串
func RandomStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//FindItemStr 根据谓词函数检测数组元素
func FindItemStr(array []string, check func(item string) bool) (index int, val string) {
	for index, val := range array {
		if check(val) {
			return index, val
		}
	}

	return -1, ""
}

//StrIncludes 字符串包含判断
func StrIncludes(arr []string, item string) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}

//StrIncludeWord 查询字符串是否包含某个词
func StrIncludeWord(fullStr, sub string) bool {
	return strings.IndexAny(fullStr, sub) >= 0
}

//SplitToUintArray 将英文分隔的字符串转成 uint数组
func SplitToUintArray(strArr string) ([]uint, error) {
	arr := make([]uint, 0)
	for _, str := range strings.Split(strArr, ",") {
		num, err := strconv.Atoi(str)
		if err != nil {
			return arr, err
		}
		if num < 0 {
			return arr, errors.New("数值不能小于0")
		}
		arr = append(arr, uint(num))
	}
	return arr, nil
}

//StrToUint 字符串转uint
func StrToUint(str string) (uint, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return uint(num), nil
}

//SplitToObjIDArray 将英文分隔的字符串转成 object id数组
func SplitToObjIDArray(strArr string) ([]primitive.ObjectID, error) {
	arr := make([]primitive.ObjectID, 0)
	for _, str := range strings.Split(strArr, ",") {
		id, err := primitive.ObjectIDFromHex(str)
		if err != nil {
			return arr, err
		}
		arr = append(arr, id)
	}
	return arr, nil
}

//StrFilter 移除不满足条件的功能
func StrFilter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

//StringArrayToObjectID 字符串数组到ObjectID数组
func StringArrayToObjectID(ids []string) []primitive.ObjectID {
	objectIds := make([]primitive.ObjectID, 0)

	for _, str := range ids {
		ID, _ := primitive.ObjectIDFromHex(str)
		objectIds = append(objectIds, ID)
	}

	return objectIds
}
