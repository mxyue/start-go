package indexes

import (
	"project/src/db/defaultdb"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var defaultClient = defaultdb.GetClient()

type d = bson.D

func createDefaultIndex(table string, index [][]d) {
	err := defaultClient.EnsureIndexes(table, index)

	if err != nil {
		logrus.Warn("default db create indexes err:", err)
	}
}

type indexCreateType struct {
	TableName string
	Index     [][]d
}

func init() {
	indexDocs := []indexCreateType{

	}

	for _, indexDoc := range indexDocs {
		createDefaultIndex(indexDoc.TableName, indexDoc.Index)
	}

}
