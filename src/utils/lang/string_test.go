package lang

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplitToUintArray(t *testing.T) {
	str := "1,2,3,4"
	arr, err := SplitToUintArray(str)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("arr->", arr)
	if !reflect.DeepEqual(arr, []uint{1, 2, 3, 4}) {
		t.Error("生成slice不符合预期")
	}
}

func TestRandomStr(t *testing.T) {
	str := RandomStr(30)
	fmt.Println("str>>", str)
	if len(str) != 30 {
		t.Error("生成长度不对")
	}
}
