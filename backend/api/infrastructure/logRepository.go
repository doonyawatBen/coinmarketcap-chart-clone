package infrastructure

import (
	"github.com/lodashventure/nlp/model"
)

type LogRepository struct{}

func (Property *LogRepository) Insert(LogHTTP *model.LogHTTP) error {
	ctx, cancel := DataBaseGlobal.GetCTX()
	defer cancel()

	collection := DataBaseGlobal.Store.Collection(ConfigGlobal.DataBase.Collection.LogHTTP)

	_, err := collection.InsertOne(ctx, LogHTTP)
	if err != nil {
		return err
	}

	return nil
}
