package defaultdb

import (
	"project/src/config"
	"project/src/db/mongodb"
	"fmt"
)

//DefaultClient 继承mongodb.client的type
type DefaultClient struct {
	*mongodb.Client
}

var dbName = config.Config.Mongo.DBName

var url = fmt.Sprintf("mongodb://%s", config.Config.Mongo.Hosts)

var sess *DefaultClient

//GetClient 获取DefaultClient实例
func GetClient() *DefaultClient {
	if sess != nil {
		return sess
	}
	client := &mongodb.Client{}
	sess = &DefaultClient{Client: client}
	sess.SetSession(url, dbName)
	return sess
}
