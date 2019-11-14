package lang

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sort"
)

func findItemIndex(arr primitive.A, ele primitive.ObjectID) int {
	for index, item := range arr {
		if item == ele {
			return index
		}
	}

	return -1
}

//SortSliceByArrayIndex 根据数组出现的顺序，对切片排序
//其中 primitive.A 的参数无法使用 []interface{}，会有转换错误
func SortSliceByArrayIndex(ids primitive.A, items []interface{}) []interface{} {
	sort.SliceStable(items, func(i, j int) bool {
		before := items[i].(primitive.M)
		after := items[j].(primitive.M)
		ii := findItemIndex(ids, before["_id"].(primitive.ObjectID))
		jj := findItemIndex(ids, after["_id"].(primitive.ObjectID))
		return ii < jj
	})

	return items
}
