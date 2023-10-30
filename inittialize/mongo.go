package initialize

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"new-mall/global"
	"new-mall/inittialize/internal"
	"strings"
)

var Mongo = new(mongo)

type (
	mongo struct{}
	Index struct {
		V    any      `bson:"v"`
		Ns   any      `bson:"ns"`
		Key  []bson.E `bson:"key"`
		Name string   `bson:"name"`
	}
)

func (m *mongo) Indexes(ctx context.Context) error {
	// Table name: Index list Column: "Table name": [][]string{{"index1", "index2"}}
	indexMap := map[string][][]string{}
	for collection, indexes := range indexMap {
		err := m.CreateIndexes(ctx, collection, indexes)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *mongo) Initialization() error {
	var opts []options.ClientOptions
	if global.Config.Mongo.IsZap {
		opts = internal.Mongo.GetClientOptions()
	}
	ctx := context.Background()
	client, err := qmgo.Open(ctx, &qmgo.Config{
		Uri:              global.Config.Mongo.Uri(),
		Coll:             global.Config.Mongo.Coll,
		Database:         global.Config.Mongo.Database,
		MinPoolSize:      &global.Config.Mongo.MinPoolSize,
		MaxPoolSize:      &global.Config.Mongo.MaxPoolSize,
		SocketTimeoutMS:  &global.Config.Mongo.SocketTimeoutMs,
		ConnectTimeoutMS: &global.Config.Mongo.ConnectTimeoutMs,
		Auth: &qmgo.Credential{
			Username: global.Config.Mongo.Username,
			Password: global.Config.Mongo.Password,
		},
	}, opts...)
	if err != nil {
		return errors.Wrap(err, "Failed to connect to mongodb database!")
	}
	global.Mongo = client
	err = m.Indexes(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *mongo) CreateIndexes(ctx context.Context, name string, indexes [][]string) error {
	collection, err := global.Mongo.Database.Collection(name).CloneCollection()
	if err != nil {
		return errors.Wrapf(err, "Failed to obtain table object of [%s]!", name)
	}
	list, err := collection.Indexes().List(ctx)
	if err != nil {
		return errors.Wrapf(err, "Failed to obtain index object of [%s]!", name)
	}
	var entities []Index
	err = list.All(ctx, &entities)
	if err != nil {
		return errors.Wrapf(err, "Failed to obtain index list of [%s]!", name)
	}
	length := len(indexes)
	indexMap1 := make(map[string][]string, length)
	for i := 0; i < length; i++ {
		length1 := len(indexes[i])
		keys := make([]string, 0, length1)
		for j := 0; j < length1; j++ {
			if indexes[i][i][0] == '-' {
				keys = append(keys, indexes[i][j], "-1")
				continue
			}
			keys = append(keys, indexes[i][j], "1")
		}
		key := strings.Join(keys, "_")
		_, o1 := indexMap1[key]
		if o1 {
			return errors.Errorf("Index [%s] is duplicated!", key)
		}
		indexMap1[key] = indexes[i]
	}
	length = len(entities)
	indexMap2 := make(map[string]map[string]string, length)
	for i := 0; i < length; i++ {
		v1, o1 := indexMap2[entities[i].Name]
		if !o1 {
			keyLength := len(entities[i].Key)
			v1 = make(map[string]string, keyLength)
			for j := 0; j < keyLength; j++ {
				v2, o2 := v1[entities[i].Key[j].Key]
				if !o2 {
					v1 = make(map[string]string)
				}
				v2 = entities[i].Key[j].Key
				v1[entities[i].Key[j].Key] = v2
				indexMap2[entities[i].Name] = v1
			}
		}
	}
	for k1, v1 := range indexMap1 {
		_, o2 := indexMap2[k1]
		if o2 {
			continue
		} // Index exists
		err = global.Mongo.Database.Collection(name).CreateOneIndex(ctx, options.IndexModel{Key: v1})
		if err != nil {
			return errors.Wrapf(err, "Failed to create index [%s]!!", k1)
		}
	}
	return nil
}
