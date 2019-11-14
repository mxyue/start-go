package helper

import (
	"project/src/utils/codewrap"
	"project/src/utils/lang"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetPathObjectID 获取gin params中的 objectId
func GetPathObjectID(c *gin.Context, key string) (primitive.ObjectID, error) {
	strID := c.Params.ByName(key)
	if strID == "" {
		return primitive.NilObjectID, codewrap.Error(codewrap.Codes.ParamsError, fmt.Sprintf("%s 参数不存在", key))
	}
	return primitive.ObjectIDFromHex(strID)
}

//GetQueryObjectID 获取gin 中的query objectId
func GetQueryObjectID(c *gin.Context, key string) (primitive.ObjectID, error) {
	strID, present := c.GetQuery(key)
	if !present {
		return primitive.NilObjectID, codewrap.Error(codewrap.Codes.ParamsError, fmt.Sprintf("%s 参数不存在", key))
	}
	return primitive.ObjectIDFromHex(strID)
}

//GetQueryObjIds 从query中获取objectIds
func GetQueryObjIds(c *gin.Context, key string) ([]primitive.ObjectID, error) {
	objIds, present := c.GetQuery(key)
	if !present {
		return []primitive.ObjectID{}, codewrap.Error(codewrap.Codes.ParamsError, fmt.Sprintf("%s 参数不存在", key))
	}
	return lang.SplitToObjIDArray(objIds)
}

//GetQueryNumIds 从query中获取numIds
func GetQueryNumIds(c *gin.Context, key string) ([]uint, error) {
	numIds, present := c.GetQuery(key)
	if !present {
		return []uint{}, codewrap.Error(codewrap.Codes.ParamsError, fmt.Sprintf("%s 参数不存在", key))
	}
	return lang.SplitToUintArray(numIds)
}

//GetQueryInt 获取query参数中的int数据
func GetQueryInt(c *gin.Context, key string) (int, error) {
	num, present := c.GetQuery(key)
	if !present {
		return 0, codewrap.Error(codewrap.Codes.ParamsError, fmt.Sprintf("%s 参数不存在", key))
	}
	return strconv.Atoi(num)
}
