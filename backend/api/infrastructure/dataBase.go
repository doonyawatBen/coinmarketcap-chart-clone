package infrastructure

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataBase struct {
	Store *mongo.Database
}

func NewDataBase() *DataBase {
	Store := mongoInit()

	DB := &DataBase{
		Store,
	}

	migrate(DB)

	return DB
}

func (db *DataBase) GetCTX() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(CtxGlobal, time.Duration(ConfigGlobal.DataBase.TimeOutSecond)*time.Second)
	return ctx, cancel
}

func mongoInit() *mongo.Database {
	clientOptions := options.Client().ApplyURI(ConfigGlobal.DataBase.Url).SetCompressors([]string{"snappy", "zlib", "zstd"})
	client, err := mongo.Connect(CtxGlobal, clientOptions)
	if err != nil {
		Log.Fatalln(nil, "init", "", "Can't Connection Database", err)
	}

	ServiceCloses = append(ServiceCloses, func() {
		client.Disconnect(CtxGlobal)
	})

	err = client.Ping(CtxGlobal, nil)
	if err != nil {
		Log.Fatalln(nil, "init", "", "Failed to Ping Database", err)
	}

	Log.Info("init", "", "Connected to Database")

	store := client.Database(ConfigGlobal.DataBase.Name)

	return store
}

func migrate(DB *DataBase) {
	indexInitQuota(DB)
}

func indexInitQuota(DB *DataBase) {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "appName", Value: 1},
			{Key: "quota", Value: 1},
			{Key: "quota_used", Value: 1},
			{Key: "updated_at", Value: 1},
		},
	}

	ctx, cancel := DB.GetCTX()
	defer cancel()

	collection := DB.Store.Collection(ConfigGlobal.DataBase.Collection.Quota)

	collection.Indexes().DropAll(ctx)

	collection.Indexes().CreateOne(ctx, indexModel)
}
