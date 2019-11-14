package redis

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestRedisSetAndGet(t *testing.T) {
	value := "value2"
	err := Client.Set("key", value, 0).Err()
	if err != nil {
		t.Error(err)
	}

	val, err := Client.Get("key").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("key", val)
	if val != value {
		t.Error("获取出错")
	}
	Client.Del("key")
}

func TestRedisInc(t *testing.T) {
	val, err := Client.Incr("num").Result()
	if err != nil {
		t.Error(err)
	}

	if success, err := Client.Expire("num", 2*time.Second).Result(); err != nil {
		fmt.Println("success", success)
	}

	fmt.Println(val)
}

func BenchmarkStringJoin1(b *testing.B) {
	b.ReportAllocs()
	input := []string{"Hello", "World"}
	for i := 0; i < b.N; i++ {
		result := strings.Join(input, " ")
		if result != "Hello World" {
			b.Error("Unexpected result: " + result)
		}
	}
}
