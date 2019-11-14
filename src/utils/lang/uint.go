package lang

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

//FindUIntIndex 查询uint元素在数组中的索引
func FindUIntIndex(arr []uint, ele uint) int {
	for index, item := range arr {
		if item == ele {
			return index
		}
	}

	return -1
}

//UintIncludes 是否包含
func UintIncludes(arr []uint, ele uint) bool {
	return FindUIntIndex(arr, ele) != -1
}

//ToUintArr interface数组转uint数组
func ToUintArr(arr []interface{}) []uint {
	defer func() {
		if err := recover(); err != nil {
			logrus.Error("ToUintArr err:", err)
			logrus.Error("origin arr:", arr)
		}
	}()
	list := make([]uint, 0)
	for _, in := range arr {
		if num, ok := in.(int64); ok {
			list = append(list, uint(num))
		} else if num, ok := in.(int32); ok {
			list = append(list, uint(num))
		}
	}
	return list
}

//UintArrNotIncludes 查找一个数组在另外一个数组不存在的元素
func UintArrNotIncludes(smallArr, bigArr []uint) []uint {
	diff := make([]uint, 0)
	for _, bItem := range bigArr {
		if FindUIntIndex(smallArr, bItem) == -1 {
			diff = append(diff, bItem)
		}
	}
	return diff
}

//RandomUint 随机生成uint
func RandomUint(max uint, skip ...uint) uint {
	rand.Seed(time.Now().UnixNano())
	count := 0
	for {
		i := uint(rand.Intn(int(max)))
		if FindUIntIndex(skip, i) == int(-1) {
			return i
		}
		if count > 1000 {
			logrus.Warn("random uint count>", count)

			for j := uint(0); j <= max; j++ {
				if FindUIntIndex(skip, j) == int(-1) {
					return j
				}
			}
			panic("无有效数据")
		}
		count++
	}
}

//RandFromRange 从给定范围中取出随机值，包含最小值，不含最大值
func RandFromRange(min, max uint) uint {
	rand.Seed(time.Now().UnixNano())
	return uint(rand.Intn(int(max-min)) + int(min))
}

var symbol = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}
var symbolLength = uint(len(symbol))

//FormatUint 将数字转成指定进制的字符串
func FormatUint(num, scale uint) string {

	if scale < 2 {
		panic("scale must large than 1")
	}
	if scale > symbolLength {
		panic(fmt.Sprintf("scale mast small than: %d", symbolLength))
	}

	str := string(symbol[num%scale])
	nemNum := num / scale

	for nemNum > scale {
		str = string(symbol[nemNum%scale]) + str
		nemNum = nemNum / scale
	}
	if nemNum > 0 {
		str = string(symbol[nemNum]) + str
	}
	return str
}

//UintToStr uint to string
func UintToStr(num uint) string {
	return strconv.FormatInt(int64(num), 10)
}
