package repository

import (
	"github.com/Raman5837/kafka.go/app/model"
	"github.com/Raman5837/kafka.go/app/types"
	"github.com/Raman5837/kafka.go/base/database"
)

// Get Consumer With Given PK And Group Id
func GetConsumer(id uint, groupId int) (Consumer *types.GetConsumer, exception error) {

	model := model.Consumer{}
	DB := database.DBManager.SqliteDB
	responseInstance := &types.GetConsumer{}
	queryResponse := DB.Table(model.TableName()).Where("id = ? AND group_id = ?", id, groupId).First(responseInstance)

	return responseInstance, queryResponse.Error

}

// Add New Consumer
func AddNewConsumer(instance *model.Consumer) (Consumer *model.Consumer, exception error) {

	DB := database.DBManager.SqliteDB
	queryResponse := DB.Table(instance.TableName()).Create(&instance)

	return instance, queryResponse.Error

}

// Get ConsumerGroup With Given Name And isActive Boolean Value
func GetConsumerGroup(name string, isActive bool) (ConsumerGroup *types.GetConsumerGroup, exception error) {

	model := model.ConsumerGroup{}
	DB := database.DBManager.SqliteDB
	responseInstance := &types.GetConsumerGroup{}
	queryResponse := DB.Table(model.TableName()).Where("name = ? AND is_active = ?", name, isActive).First(responseInstance)

	return responseInstance, queryResponse.Error

}

// Add New ConsumerGroup
func AddNewConsumerGroup(instance *model.ConsumerGroup) (ConsumerGroup *model.ConsumerGroup, exception error) {

	DB := database.DBManager.SqliteDB
	queryResponse := DB.Table(instance.TableName()).Create(&instance)

	return instance, queryResponse.Error

}
