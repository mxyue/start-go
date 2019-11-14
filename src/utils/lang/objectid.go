package lang

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FindObjectIDIndex 查询object元素在数组中的索引
func FindObjectIDIndex(arr []primitive.ObjectID, ele primitive.ObjectID) int {
	for index, item := range arr {
		if item == ele {
			return index
		}
	}

	return -1
}

//ObjectIDIncludes 检查objectId是否存在于数组
func ObjectIDIncludes(arr []primitive.ObjectID, ele primitive.ObjectID) bool {
	return FindObjectIDIndex(arr, ele) != -1
}
