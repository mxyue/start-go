package lang

import (
	"reflect"
)

//IsEmpty 判定对象是否为空
func IsEmpty(object interface{}) bool {
	if object == nil {
		return true
	}

	objValue := reflect.ValueOf(object)

	switch objValue.Kind() {

	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0

	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}
		deref := objValue.Elem().Interface()
		return IsEmpty(deref)
	default:
		zero := reflect.Zero(objValue.Type())
		return reflect.DeepEqual(object, zero.Interface())
	}
}

//Contains 判断数组或者Slice是否包含目标元素
func Contains(source interface{}, target interface{}) bool {
	sourceValue := reflect.ValueOf(source)
	switch reflect.TypeOf(source).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < sourceValue.Len(); i++ {
			if sourceValue.Index(i).Interface() == target {
				return true
			}
		}
	}
	return false
}
