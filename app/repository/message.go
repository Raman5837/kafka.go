package repository

import (
	"github.com/Raman5837/kafka.go/app/model"
	"github.com/Raman5837/kafka.go/app/types"
	"github.com/Raman5837/kafka.go/base/database"
)

// Get Messages For Given Partition And Offset >= Given Offset Number With Given Limit
func GetMessages(partitionId uint64, offset uint64, limit int) (Messages *[]types.GetMessage, exception error) {

	model := model.Message{}
	DB := database.DBManager.SqliteDB
	responseInstance := &[]types.GetMessage{}
	queryResponse := DB.Table(model.TableName()).
		Where("partition_id = ? AND offset >= ?", partitionId, offset).
		Order("offset ASC").Limit(limit).Find(responseInstance)

	return responseInstance, queryResponse.Error

}

// Get All Message Present In System
func GetAllMessageCount() (Count int64, exception error) {

	var star int64
	model := model.Message{}
	DB := database.DBManager.SqliteDB
	queryResponse := DB.Table(model.TableName()).Count(&star)

	return star, queryResponse.Error

}

// Add New Message
func AddNewMessage(instance *model.Message) (Message *model.Message, exception error) {

	DB := database.DBManager.SqliteDB
	queryResponse := DB.Table(instance.TableName()).Create(&instance)

	return instance, queryResponse.Error

}
