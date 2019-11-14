package lang

import (
	"math/rand"
	"time"
)

//RandomInt 随机数[0, max)
func RandomInt(max int) int {
	if max == 0 {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

//RandomChan 随机[0, max)，然后随机从里面取数据，取出的数将被移除随机数中
func RandomChan(max int, yield chan int) {
	value := make([]int, 0, max)
	for i := 0; i < max; i++ {
		value = append(value, i)
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < max; i++ {
		random := rand.Intn(len(value))
		yield <- value[random]
		value = append(value[:random], value[random+1:]...)
	}
	close(yield)
}

//Int8Includes int8包含
func Int8Includes(arr []int8, item int8) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}
