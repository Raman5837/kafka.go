package utils

import (
	interfaces "github.com/Raman5837/kafka.go/app/interface"
	"github.com/Raman5837/kafka.go/app/model"
	"github.com/Raman5837/kafka.go/app/repository"
	"github.com/Raman5837/kafka.go/app/types"
	"github.com/Raman5837/kafka.go/base/utils"
	"gorm.io/gorm"
)

type partitionAssigner struct{}

// Returns New PartitionAssigner Instance
func NewPartitionAssigner() interfaces.PartitionAssignerInterface {
	return &partitionAssigner{}
}

// Get The Next Partition To Assign (Using Round-Robin Fashion)
func (assignor *partitionAssigner) Next(topicId uint) (*types.GetPartition, error) {

	partitions, queryErr := repository.GetPartitionByTopicId(topicId)

	if queryErr != nil {
		return nil, queryErr
	}

	allPartitions := *partitions
	if len(allPartitions) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	assigned, queryErr := repository.GetLastAssignedPartitionForTopic(topicId)

	if queryErr != nil {
		return nil, queryErr
	}

	lastAssignedIndex := getPartitionIndex(*partitions, assigned.PartitionId)
	nextPartitionIndex := (lastAssignedIndex + 1) % len(*partitions)
	nextPartition := allPartitions[nextPartitionIndex]

	// Update The LastAssignedPartition
	newObject := &model.LastAssignedPartition{TopicID: topicId, PartitionId: nextPartition.PartitionId}
	if updatedAssigned, queryErr := repository.SaveLastAssignedPartition(newObject); queryErr == nil {
		utils.Logger.InfoF("LastAssignedPartition For Topic %v Updated With Partition: %v", topicId, updatedAssigned.PartitionId)
	}

	return &nextPartition, nil

}

// Helper Function To Get Partition Index If Exists Else Returns -1
func getPartitionIndex(partitions []types.GetPartition, partitionId uint) int {

	for index, partition := range partitions {

		if partition.PartitionId == partitionId {
			return index
		}
	}

	return -1
}
