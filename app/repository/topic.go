package repository

import (
	"github.com/Raman5837/kafka.go/app/model"
	"github.com/Raman5837/kafka.go/app/types"
	"github.com/Raman5837/kafka.go/base/database"
)

// Get Topic With Given Name
func GetTopic(name string) (Topic *types.GetTopic, exception error) {

	model := model.Topic{}
	DB := database.DBManager.SqliteDB
	responseInstance := &types.GetTopic{}
	queryResponse := DB.Table(model.TableName()).Where("name = ?", name).First(responseInstance)

	return responseInstance, queryResponse.Error

}

// Create New Topic
func CreateTopic(instance *model.Topic) (Topic *model.Topic, exception error) {

	DB := database.DBManager.SqliteDB
	queryResponse := DB.Table(instance.TableName()).Create(&instance)

	return instance, queryResponse.Error

}
