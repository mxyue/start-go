package lang

import "reflect"

//ToSlice 把一个本质是数组的interface{}转成数组
//参考 https://segmentfault.com/q/1010000000198391
//quiz.biz.Value可能是数组，也可能是单个元素，所以声明是interfae{}
//是数组的时候直接迭代 biz.Value.([]interface{}) 会引发panic
func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}
