package user

import "go.mongodb.org/mongo-driver/bson/primitive"

//TABLE user
const TABLE = "users"

//User user 用户账号
type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Name         string             `bson:"name"`                             //名称
	Portrait     string             `bson:"portrait"`                         //头像
	Password     string             `bson:"password" json:"-"`                //密码
}
