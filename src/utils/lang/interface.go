package lang

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UNArrToStrArr interface转string arr
func UNArrToStrArr(items []interface{}) (strArr []string) {
	for _, item := range items {
		if strItem, ok := item.(string); ok {
			strArr = append(strArr, strItem)
		}
	}
	return strArr
}

//UNJoin join
func UNJoin(items []interface{}, symbol string) string {
	str := ""
	for i, item := range items {
		if i == len(items)-1 {
			str += fmt.Sprintf("%v", item)
		} else {
			str += fmt.Sprintf("%v%s", item, symbol)
		}
	}
	return str
}

//UNArrToObjectIDArr interface数组转uint数组
func UNArrToObjectIDArr(arr []interface{}) []primitive.ObjectID {
	list := make([]primitive.ObjectID, 0)
	for _, in := range arr {
		if id, ok := in.(primitive.ObjectID); ok {
			list = append(list, id)
		} else {
			oid, err := primitive.ObjectIDFromHex(fmt.Sprint(in))
			if err != nil {
				panic(err)
			}
			list = append(list, oid)
		}
	}
	return list
}
