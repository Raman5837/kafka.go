package repository

import (
	"github.com/Raman5837/kafka.go/app/model"
	"github.com/Raman5837/kafka.go/app/types"
	"github.com/Raman5837/kafka.go/base/database"
)

// Get Partition With Given Topic Id And Partition Id
func GetPartition(topicId uint64, partitionId uint64) (Partition *types.GetPartition, exception error) {

	model := model.Partition{}
	DB := database.DBManager.SqliteDB
	responseInstance := &types.GetPartition{}
	queryResponse := DB.Table(model.TableName()).Where("topic_id = ? AND partition_id = ?", topicId, partitionId).First(responseInstance)

	return responseInstance, queryResponse.Error

}

// Get Partition With Given Topic Id
func GetPartitionByTopicId(topicId uint64) (Partition *[]types.GetPartition, exception error) {

	model := model.Partition{}
	DB := database.DBManager.SqliteDB
	allPartitions := &[]types.GetPartition{}
	queryResponse := DB.Table(model.TableName()).Where("topic_id = ?", topicId).Find(allPartitions)

	return allPartitions, queryResponse.Error

}

// Create New Partition
func CreatePartition(instance *model.Partition) (Partition *model.Partition, exception error) {

	DB := database.DBManager.SqliteDB
	queryResponse := DB.Table(instance.TableName()).Create(&instance)

	return instance, queryResponse.Error

}
