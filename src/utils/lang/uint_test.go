package lang

import (
	"reflect"
	"testing"
)

func TestUintArrNotIncludes(t *testing.T) {
	bizArr := []uint{1, 2, 3, 4}
	smallArr := []uint{1, 2}
	diffArr := UintArrNotIncludes(smallArr, bizArr)
	if !reflect.DeepEqual(diffArr, []uint{3, 4}) {
		t.Error("arr数据不相等")
	}
}

func TestRandomUint(t *testing.T) {
	result := RandomUint(3, 0, 1)
	if result != uint(2) {
		t.Error("不是预期的数")
	}
}
