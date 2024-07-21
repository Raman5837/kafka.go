package repository

import (
	"github.com/Raman5837/kafka.go/app/model"
	"github.com/Raman5837/kafka.go/app/types"
	"github.com/Raman5837/kafka.go/base/database"
)

// Get ConsumerAssignment With Given Partition Id And Group Id
func GetConsumerAssignment(groupId int, partitionId int) (Assignment *types.GetConsumerAssignment, exception error) {

	DB := database.DBManager.SqliteDB
	model := model.ConsumerAssignment{}
	responseInstance := &types.GetConsumerAssignment{}
	queryResponse := DB.Table(model.TableName()).Where("group_id = ? AND partition_id = ?", groupId, partitionId).First(responseInstance)

	return responseInstance, queryResponse.Error

}

// Get All Assigned Consumers For Given Group Id
func GetAssignedConsumersOfAGroup(groupId uint64) (Assignment *[]types.GetConsumerAssignment, exception error) {

	DB := database.DBManager.SqliteDB
	model := model.ConsumerAssignment{}
	responseInstance := &[]types.GetConsumerAssignment{}
	queryResponse := DB.Table(model.TableName()).Where("group_id = ?", groupId).Find(responseInstance)

	return responseInstance, queryResponse.Error

}

// Create A ConsumerAssignment For Given Consumer And Partition
func AssignPartitionToConsumer(instance model.ConsumerAssignment) (Assignment model.ConsumerAssignment, exception error) {

	DB := database.DBManager.SqliteDB
	queryResponse := DB.Table(instance.TableName()).Create(&instance)

	return instance, queryResponse.Error

}

// Delete ConsumerAssignment For Given Consumer Id
func DeleteConsumerAssignment(consumerId uint64) (exception error) {

	model := model.ConsumerAssignment{}
	DB := database.DBManager.SqliteDB
	queryResponse := DB.Table(model.TableName()).Where("consumer_id = ?", consumerId).Delete(model)

	return queryResponse.Error

}

// Delete ConsumerAssignment For Given Consumer Ids
func DeleteAllConsumerAssignment(consumerIds []uint64) (exception error) {

	model := model.ConsumerAssignment{}
	DB := database.DBManager.SqliteDB
	queryResponse := DB.Table(model.TableName()).Where("consumer_id IN ?", consumerIds).Delete(model)

	return queryResponse.Error

}

// Get LastAssignedPartition With Given Topic Id And Partition Id
func GetLastAssignedPartition(topicId int, partitionId int) (LastAssignedPartition *types.GetLastAssignedPartition, exception error) {

	DB := database.DBManager.SqliteDB
	model := model.LastAssignedPartition{}
	responseInstance := &types.GetLastAssignedPartition{}
	queryResponse := DB.Table(model.TableName()).
		Where("topic_id = ? AND partition_id = ?", topicId, partitionId).
		First(responseInstance)

	return responseInstance, queryResponse.Error

}

// Get LastAssignedPartition With Given Topic
func GetLastAssignedPartitionForTopic(topicId uint64) (LastAssignedPartition *types.GetLastAssignedPartition, exception error) {

	DB := database.DBManager.SqliteDB
	model := model.LastAssignedPartition{TopicID: topicId}
	responseInstance := &types.GetLastAssignedPartition{}
	queryResponse := DB.Table(model.TableName()).Where("topic_id = ?", topicId).FirstOrInit(responseInstance)

	return responseInstance, queryResponse.Error

}

// Save LastAssignedPartition
func SaveLastAssignedPartition(instance *model.LastAssignedPartition) (LAP *model.LastAssignedPartition, exception error) {

	DB := database.DBManager.SqliteDB
	queryResponse := DB.Table(instance.TableName()).Save(&instance)

	return instance, queryResponse.Error

}
