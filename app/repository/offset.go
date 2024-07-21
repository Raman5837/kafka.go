package repository

import (
	"github.com/Raman5837/kafka.go/app/model"
	"github.com/Raman5837/kafka.go/app/types"
	"github.com/Raman5837/kafka.go/base/database"
)

// Get Offset For Given Consumer And Partition
func GetOffset(consumerId uint64, partitionId uint64) (Offset *types.GetOffset, exception error) {

	model := model.Offset{}
	DB := database.DBManager.SqliteDB
	responseInstance := &types.GetOffset{}
	queryResponse := DB.Table(model.TableName()).
		Where("consumer_id = ? AND partition_id = ?", consumerId, partitionId).First(responseInstance)

	return responseInstance, queryResponse.Error

}

// Update Offset For Given Consumer And Partition
func UpdateOffset(consumerId uint64, partitionId uint64, newOffset uint64) (exception error) {

	model := model.Offset{}
	DB := database.DBManager.SqliteDB
	queryResponse := DB.
		Table(model.TableName()).
		Where("consumer_id = ? AND partition_id = ?", consumerId, partitionId).
		Update("number", newOffset)

	return queryResponse.Error

}
