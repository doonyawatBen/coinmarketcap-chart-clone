package service

import (
	"github.com/lodashventure/nlp/model"
	"github.com/lodashventure/nlp/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuotaService struct {
	quotaRepository repository.QuotaRepository
}

func (Property *QuotaService) GetByAppName(appName string) (*model.Quota, error) {
	return Property.quotaRepository.GetByAppName(appName)
}

func (Property *QuotaService) UpdateByAppName(appName string, quota *model.UpdateQuotaRequest) error {
	return Property.quotaRepository.UpdateByAppName(appName, quota)
}

func (Property *QuotaService) UpdateResetByAppName(appName string) error {
	return Property.quotaRepository.UpdateResetByAppName(appName)
}

func (Property *QuotaService) UpdateResetAll() error {
	return Property.quotaRepository.UpdateResetAll()
}

func (Property *QuotaService) GetResultByAppName(appName string) (*mongo.UpdateResult, error) {
	return Property.quotaRepository.GetResultByAppName(appName)
}

func (Property *QuotaService) InsertByAppName(appName string, quotaLimit int) error {
	return Property.quotaRepository.InsertByAppName(appName, quotaLimit)
}
