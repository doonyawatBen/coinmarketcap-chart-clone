package repository

import (
	"fmt"
	"time"

	"github.com/lodashventure/nlp/infrastructure"
	"github.com/lodashventure/nlp/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QuotaRepository struct{}

func (Property *QuotaRepository) GetByAppName(appName string) (*model.Quota, error) {
	var result *model.Quota

	ctx, cancel := infrastructure.DataBaseGlobal.GetCTX()
	defer cancel()

	collection := infrastructure.DataBaseGlobal.Store.Collection(infrastructure.ConfigGlobal.DataBase.Collection.Quota)

	filter := bson.M{
		"appName": appName,
	}

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, fmt.Errorf("Not found Quota %s : %s", appName, err.Error())
	}

	return result, nil
}

func (Property *QuotaRepository) UpdateByAppName(appName string, quota *model.UpdateQuotaRequest) error {
	ctx, cancel := infrastructure.DataBaseGlobal.GetCTX()
	defer cancel()

	collection := infrastructure.DataBaseGlobal.Store.Collection(infrastructure.ConfigGlobal.DataBase.Collection.Quota)

	str := time.Now().Format(time.RFC3339)
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		infrastructure.Log.Error(nil, "quota", appName, "Time Parse invalid", err)
		return err
	}

	filter := bson.M{
		"appName": appName,
	}

	update := bson.M{
		"$set": bson.M{
			"quota":      quota.Quota,
			"updated_at": t,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (Property *QuotaRepository) UpdateResetByAppName(appName string) error {
	ctx, cancel := infrastructure.DataBaseGlobal.GetCTX()
	defer cancel()

	collection := infrastructure.DataBaseGlobal.Store.Collection(infrastructure.ConfigGlobal.DataBase.Collection.Quota)

	str := time.Now().Format(time.RFC3339)
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		infrastructure.Log.Error(nil, "quota", appName, "Time Parse invalid", err)
		return err
	}

	filter := bson.M{
		"appName": appName,
	}

	update := bson.M{
		"$set": bson.M{
			"quota_used": 0,
			"updated_at": t,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (Property *QuotaRepository) UpdateResetAll() error {
	ctx, cancel := infrastructure.DataBaseGlobal.GetCTX()
	defer cancel()

	collection := infrastructure.DataBaseGlobal.Store.Collection(infrastructure.ConfigGlobal.DataBase.Collection.Quota)

	str := time.Now().Format(time.RFC3339)
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		infrastructure.Log.Error(nil, "quota", "", "Time Parse invalid", err)
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"quota_used": 0,
			"updated_at": t,
		},
	}

	_, err = collection.UpdateMany(ctx, bson.M{}, update)
	if err != nil {
		return err
	}

	return nil
}

func (Property *QuotaRepository) GetResultByAppName(appName string) (*mongo.UpdateResult, error) {
	var result *mongo.UpdateResult

	ctx, cancel := infrastructure.DataBaseGlobal.GetCTX()
	defer cancel()

	collection := infrastructure.DataBaseGlobal.Store.Collection(infrastructure.ConfigGlobal.DataBase.Collection.Quota)

	filter := bson.M{"appName": appName}

	str := time.Now().Format(time.RFC3339)
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		infrastructure.Log.Error(nil, "quota", appName, "Time Parse invalid", err)
		return result, err
	}

	// Define the pipeline stages
	pipeline := []bson.M{
		{
			"$addFields": bson.M{
				"currentMonth": bson.M{
					"$month": bson.M{"$toDate": bson.M{"$toLong": bson.M{"$dateFromParts": bson.M{"year": bson.M{"$year": bson.M{"$toDate": bson.M{"$toLong": "$$NOW"}}}, "month": bson.M{"$month": bson.M{"$toDate": bson.M{"$toLong": "$$NOW"}}}}}}},
				},
			},
		},
		{
			"$addFields": bson.M{
				"quota_used": bson.M{
					"$cond": bson.M{
						"if": bson.M{
							"$ne": []interface{}{"$currentMonth", bson.M{"$month": bson.M{
								"$toDate": "$updated_at",
							}}},
						},
						"then": 0,
						"else": "$quota_used",
					},
				},
				"updated_at": bson.M{
					"$cond": bson.M{
						"if": bson.M{
							"$ne": []interface{}{"$currentMonth", bson.M{"$month": bson.M{
								"$toDate": "$updated_at",
							}}},
						},
						"then": t,
						"else": "$updated_at",
					},
				},
			},
		},
		{
			"$addFields": bson.M{
				"quota_used": bson.M{
					"$cond": bson.M{
						"if":   bson.M{"$gt": bson.A{"$quota", "$quota_used"}},
						"then": bson.M{"$add": bson.A{"$quota_used", 1}},
						"else": "$quota_used",
					},
				},
				"updated_at": bson.M{
					"$cond": bson.M{
						"if":   bson.M{"$gt": bson.A{"$quota", "$quota_used"}},
						"then": t,
						"else": "$updated_at",
					},
				},
			},
		},
		{
			"$unset": "currentMonth",
		},
	}

	// Define options for the update operation
	opts := options.Update().SetUpsert(false)

	result, err = collection.UpdateOne(ctx, filter, pipeline, opts)
	if err != nil {
		return result, fmt.Errorf("Can't Update Quota %s : %s", appName, err.Error())
	}

	return result, nil
}

func (Property *QuotaRepository) InsertByAppName(appName string, quotaLimit int) error {
	ctx, cancel := infrastructure.DataBaseGlobal.GetCTX()
	defer cancel()

	collection := infrastructure.DataBaseGlobal.Store.Collection(infrastructure.ConfigGlobal.DataBase.Collection.Quota)

	str := time.Now().Format(time.RFC3339)
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		infrastructure.Log.Error(nil, "quota", appName, "Time Parse invalid", err)
		return fmt.Errorf("Sorry, we were unable to process your request at the moment. Please try again later or contact customer support for assistance.")
	}

	quota := &model.QuotaDB{
		AppName:   appName,
		Quota:     quotaLimit,
		QuotaUsed: 1,
		CreatedAt: t,
		UpdatedAt: t,
	}

	_, err = collection.InsertOne(ctx, quota)
	if err != nil {
		infrastructure.Log.Error(nil, "quota", appName, "Can't insert quota", err)
		return fmt.Errorf("Sorry, we were unable to process your request at the moment. Please try again later or contact customer support for assistance.")
	}

	return nil
}
