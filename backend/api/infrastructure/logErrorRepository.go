package infrastructure

import (
	"github.com/lodashventure/nlp/model"
)

type LogErrorRepository struct{}

func (Property *LogErrorRepository) Insert(LogError *model.LogError) error {
	ctx, cancel := DataBaseGlobal.GetCTX()
	defer cancel()

	collection := DataBaseGlobal.Store.Collection(ConfigGlobal.DataBase.Collection.LogError)

	_, err := collection.InsertOne(ctx, LogError)
	if err != nil {
		return err
	}

	return nil
}
