package defaultdb

import (
	"testing"

	"go.mongodb.org/mongo-driver/mongo"

	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)


var defaultDB = GetClient()

func TestMain(m *testing.M) {
	fmt.Println("begin")
	m.Run()
	fmt.Println("end")
}

func TestEnsureIndex(t *testing.T) {
	type d = bson.D
	err := defaultDB.EnsureIndexes("tests", [][]d{
		{d{{"name", 1}}, d{{"unique", true}}},
		{d{{"title", -1}}, d{{"unique", true}, {"sparse", true}}},
	})
	if err != nil {
		t.Error(err)
		fmt.Println("get error>", err)
	}
}

func TestBulkWrite(t *testing.T) {
	type testType struct {
		Name string
	}
	test := testType{"test"}
	updateDoc := mongo.NewUpdateOneModel()
	updateDoc.SetFilter(bson.D{})
	updateDoc.SetUpdate(bson.M{"$set": test})
	updateDoc.SetUpsert(true)
	result, err := defaultDB.BulkWrite("test", []mongo.WriteModel{updateDoc})
	if err != nil {
		t.Error(err)
	}
	fmt.Println("result>", result)
}
