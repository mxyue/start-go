package mongodb

import (
	"context"
	"reflect"
	"time"

	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Client 客户端
type Client struct {
	Database *mongo.Database
}

//SetSession 设置session
func (client *Client) SetSession(url, dbName string) {
	sess, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = sess.Connect(ctx)
	if err != nil {
		logrus.Error("[db.SetSession.Connect]", err)
	}
	pinErr := sess.Ping(ctx, nil)
	if pinErr != nil {
		logrus.Error("[db.SetSession.Ping]", pinErr)
		panic(pinErr)
	}
	logrus.Info("mongo connected, use database: ", dbName)
	client.Database = sess.Database(dbName)
}

func getDefaultCtx(args ...time.Duration) (context.Context, context.CancelFunc) {
	timeout := time.Duration(3)
	for i, arg := range args {
		if i == 0 && arg > 0 {
			timeout = arg
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	return ctx, cancel
}

func getTimeout(opt bson.M) time.Duration {
	if opt["timeout"] != nil {
		return time.Duration(opt["timeout"].(int))
	}
	return time.Duration(0)
}

func getFindOpt(opt bson.M) *options.FindOptions {
	findOptions := options.Find()
	if opt["limit"] != nil {
		findOptions.SetLimit(int64(opt["limit"].(int)))
	}
	if opt["select"] != nil {
		findOptions.SetProjection(opt["select"])
	}
	if opt["skip"] != nil {
		findOptions.SetSkip(int64(opt["skip"].(int)))
	}
	if opt["sort"] != nil {
		findOptions.SetSort(opt["sort"])
	}
	return findOptions
}

//通过反射将cursor中的数据存到 result中
func decodeCursor(ctx context.Context, cur *mongo.Cursor, resultv reflect.Value) error {
	defer cur.Close(ctx)
	slicev := resultv.Elem()
	elemType := slicev.Type().Elem()

	defer func() {
		if err := recover(); err != nil {
			logrus.Error("[db.Find.Next]", err)
		}
	}()

	for cur.Next(ctx) {
		//根据type生成新的对象
		elemPtr := reflect.New(elemType)
		//elemPtr.Interface()提取地址
		if err := cur.Decode(elemPtr.Interface()); err != nil {
			logrus.Error("[db.Find.Decode]", err)
			return err
		}
		//将元素的值append到 slicev 中
		slicev.Set(reflect.Append(slicev, elemPtr.Elem()))
	}

	return nil
}

//Find 通过反射直接将传入的struct赋值 limit,select,skip,sort
func (client *Client) Find(collName string, filter bson.D, opt bson.M, result interface{}) error {
	resultv := reflect.ValueOf(result)
	if resultv.Kind() != reflect.Ptr {
		panic("result argument must be a slice address")
	}
	ctx, cancel := getDefaultCtx(getTimeout(opt))
	defer cancel()
	coll := client.Database.Collection(collName)
	cur, err := coll.Find(ctx, filter, getFindOpt(opt))
	if err != nil {
		logrus.Error("[db.Find.Find]", err)
		return err
	}
	return decodeCursor(ctx, cur, resultv)
}

//Aggregate 聚合查询
func (client *Client) Aggregate(collName string, pipeline interface{}, result interface{}) error {
	resultv := reflect.ValueOf(result)
	if resultv.Kind() != reflect.Ptr {
		panic("result argument must be a slice address")
	}
	ctx, cancel := getDefaultCtx()
	defer cancel()
	cur, err := client.Database.Collection(collName).Aggregate(ctx, pipeline)
	if err != nil {
		logrus.Error("[db.Aggregate.Aggregate]", err)
		return err
	}
	return decodeCursor(ctx, cur, resultv)
}

//AggregateDoc 聚合查询
func (client *Client) AggregateDoc(collName string, pipeline interface{}) ([]*bson.D, error) {
	results := make([]*bson.D, 0)

	ctx, cancel := getDefaultCtx()
	defer cancel()
	cur, err := client.Database.Collection(collName).Aggregate(ctx, pipeline)
	if err != nil {
		logrus.Error("[db.Aggregate.Aggregate]", err)
		return results, err
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		elem := &bson.D{}
		if err := cur.Decode(elem); err != nil {
			logrus.Error(err)
		}
		results = append(results, elem)
		// do something with elem....
	}

	if err := cur.Err(); err != nil {
		logrus.Error(err)
	}
	return results, err
}

//FindReturnDoc 直接将查询结果通过bson.M返回
func (client *Client) FindReturnDoc(collName string, filter bson.D, opt bson.M) ([]bson.M, error) {
	ctx, cancel := getDefaultCtx()
	defer cancel()
	coll := client.Database.Collection(collName)
	cur, err := coll.Find(ctx, filter, getFindOpt(opt))
	if err != nil {
		logrus.Error("[db.FindReturnDoc.Find]", err)
	}

	defer cur.Close(ctx)

	results := make([]bson.M, 0)
	for cur.Next(ctx) {
		doc := bson.M{}
		if err := cur.Decode(&doc); err != nil {
			logrus.Error("[Find.Decode]", err)
		}
		results = append(results, doc)
	}

	return results, err
}

//Count 查询数量
func (client *Client) Count(collName string, filter bson.D) int {
	ctx, cancel := getDefaultCtx()
	defer cancel()
	coll := client.Database.Collection(collName)
	total, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		logrus.Error("[db.Count]", err)
	}
	return int(total)
}

func getFindOneOpt(opt bson.M) *options.FindOneOptions {
	findOptions := options.FindOne()
	if opt["select"] != nil {
		findOptions.SetProjection(opt["select"])
	}
	if opt["skip"] != nil {
		findOptions.SetSkip(int64(opt["skip"].(int)))
	}
	if opt["sort"] != nil {
		findOptions.SetSort(opt["sort"])
	}
	return findOptions
}

//FindOne 查询单个
func (client *Client) FindOne(collName string, filter bson.D, opt bson.M, result interface{}) error {
	resultv := reflect.ValueOf(result)
	if resultv.Kind() != reflect.Ptr {
		panic("result argument must be a slice address")
	}
	ctx, cancel := getDefaultCtx()
	defer cancel()
	coll := client.Database.Collection(collName)

	err := coll.FindOne(ctx, filter, getFindOneOpt(opt)).Decode(result)
	return err
}

//InsertMany 插入多个自定义对象
func (client *Client) InsertMany(collName string, objs []interface{}) (*mongo.InsertManyResult, error) {
	ctx, cancel := getDefaultCtx()
	defer cancel()
	return client.Database.Collection(collName).InsertMany(ctx, objs)
}

//InsertOne 添加一个
func (client *Client) InsertOne(collName string, data interface{}) (interface{}, error) {
	ctx, cancel := getDefaultCtx()
	defer cancel()
	coll := client.Database.Collection(collName)
	res, err := coll.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	id := res.InsertedID
	return id, err
}

//UpdateOne 更新一个
func (client *Client) UpdateOne(collName string, filter bson.D, data interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := getDefaultCtx()
	defer cancel()
	coll := client.Database.Collection(collName)
	res, err := coll.UpdateOne(ctx, filter, data)
	return res, err
}

//UpdateMany 更新多个
func (client *Client) UpdateMany(collName string, filter bson.D, data interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := getDefaultCtx()
	defer cancel()
	coll := client.Database.Collection(collName)
	res, err := coll.UpdateMany(ctx, filter, data)
	return res, err
}

//DeleteOne 删除一个
func (client *Client) DeleteOne(collName string, filter bson.D) error {
	ctx, cancel := getDefaultCtx()
	defer cancel()
	coll := client.Database.Collection(collName)
	_, err := coll.DeleteOne(ctx, filter)
	return err
}

//DeleteMany 删除多个
func (client *Client) DeleteMany(collName string, filter bson.D) error {
	ctx, cancel := getDefaultCtx()
	defer cancel()
	coll := client.Database.Collection(collName)
	_, err := coll.DeleteMany(ctx, filter)
	return err
}

//getFindOneAndUpdateOpt  用于FindOneAndUpdate中处理option参数
func getFindOneAndUpdateOpt(opt bson.M) *options.FindOneAndUpdateOptions {
	findOptions := options.FindOneAndUpdate()
	if opt["upsert"] != nil {
		findOptions.SetUpsert(opt["upsert"].(bool))
	}
	if opt["new"] == true {
		findOptions.SetReturnDocument(1)
	}
	return findOptions
}

//FindOneAndUpdate 查询或更新 options[upsert:bool,new:bool]
func (client *Client) FindOneAndUpdate(collName string, filter bson.D, update interface{}, opts bson.M, result interface{}) error {
	resultv := reflect.ValueOf(result)
	if resultv.Kind() != reflect.Ptr {
		panic("result argument must be a slice address")
	}
	ctx, cancel := getDefaultCtx()
	defer cancel()
	return client.Database.Collection(collName).FindOneAndUpdate(ctx, filter, update, getFindOneAndUpdateOpt(opts)).Decode(result)
}

//Distinct distinct
func (client *Client) Distinct(collName string, fieldName string, filter bson.D) ([]interface{}, error) {
	ctx, cancel := getDefaultCtx()
	defer cancel()
	return client.Database.Collection(collName).Distinct(ctx, fieldName, filter)
}

//BulkWrite 批量写入
func (client *Client) BulkWrite(collName string, models []mongo.WriteModel) (*mongo.BulkWriteResult, error) {
	ctx, cancel := getDefaultCtx()
	defer cancel()
	return client.Database.Collection(collName).BulkWrite(ctx, models)
}

//LastNumberInc inc并返回最新的numId
func (client *Client) LastNumberInc(modelName string) (uint, error) {
	doc := struct {
		Count uint
	}{}
	err := client.FindOneAndUpdate("identitycounters", bson.D{{"model", modelName}}, bson.M{"$inc": bson.M{"count": 1}}, bson.M{"upsert": true, "new": true}, &doc)
	return doc.Count, err
}

//ListIndex 打印indexes。目前未使用，待完善
//func (client *Client) ListIndexes(collName string) error {
//
//	idxs := client.Database.Collection(collName).Indexes()
//	ctx, cancel := getDefaultCtx()
//	defer cancel()
//
//	cur, err := idxs.List(ctx)
//	if err != nil {
//		return errors.Wrap(err, "unable to list indexes")
//	}
//
//	for cur.Next(ctx) {
//		d := bson.M{}
//
//		if err := cur.Decode(&d); err != nil {
//			return errors.Wrap(err, "unable to decode bson index document")
//		}
//
//		fmt.Println("---index>", d)
//	}
//
//	return nil
//}

func buildIndexOpt(opt bson.D) *options.IndexOptions {
	optM := opt.Map()
	optionBuilder := options.Index()
	if optM["unique"] == true {
		optionBuilder.SetUnique(true)
	}
	if optM["sparse"] == true {
		optionBuilder.SetSparse(true)
	}
	if optM["background"] == true {
		optionBuilder.SetBackground(true)
	}
	return optionBuilder
}

//EnsureIndexes 添加多个索引
func (client *Client) EnsureIndexes(collName string, opts [][]bson.D) error {
	idxs := client.Database.Collection(collName).Indexes()
	ctx, cancel := getDefaultCtx()
	defer cancel()
	indexModels := make([]mongo.IndexModel, 0)
	for _, opt := range opts {
		indexModels = append(indexModels, mongo.IndexModel{
			Keys:    opt[0],
			Options: buildIndexOpt(opt[1]),
		})
	}
	results, err := idxs.CreateMany(ctx, indexModels)
	if err != nil {
		logrus.Error("create indexes error:", collName, err)
	}
	logrus.Info("create indexes results>", collName, results)
	return err
}
