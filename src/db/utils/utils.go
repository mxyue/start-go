package utils

import (
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

//BuildUpdateDoc 创建批量更新的doc
func BuildUpdateDoc(filter, updateDoc interface{}, upsert bool) *mongo.UpdateOneModel {
	uDoc := mongo.NewUpdateOneModel()
	uDoc.SetFilter(filter)
	uDoc.SetUpdate(updateDoc)
	uDoc.SetUpsert(upsert)
	return uDoc
}

//IsMultiInsert 判断是否为重复添加
func IsMultiInsert(err error) bool {
	return strings.HasPrefix(err.Error(), "multiple write errors")
}
