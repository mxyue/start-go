package admin

import "go.mongodb.org/mongo-driver/bson/primitive"

//TABLE admin
const TABLE = "admins"

//Admin admin 后台管理员
type Admin struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name"`               //名称
	Phone    string             `bson:"phone" json:"phone"` //手机号
	Password string             `bson:"password" json:"-"`  //密码
}
