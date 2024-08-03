package utils

import (
	"github.com/Raman5837/kafka.go/app/model"
	"github.com/Raman5837/kafka.go/app/repository"
	"github.com/Raman5837/kafka.go/base/database"
	"github.com/Raman5837/kafka.go/base/utils"
	"gorm.io/gorm"
)

// Consumer Rebalancing
type RebalanceService struct{}

// Returns A New Instance Of RebalanceService
func NewRebalanceService() *RebalanceService {
	return &RebalanceService{}
}

// Re-balance The Consumer Group
func (instance *RebalanceService) RebalanceGroup(groupId uint, topicId uint) (exception error) {

	DB := database.DBManager.SqliteDB

	return utils.WithTransactionAtomic(DB, func(transaction *gorm.DB) error {

		// Get All Consumers Of Given Group
		consumers, queryErr := repository.GetConsumersOfAGroup(groupId)

		if queryErr != nil {
			return queryErr
		}

		allConsumers := *consumers
		count := len(allConsumers)

		if count == 0 {
			utils.Logger.InfoF("[CG Rebalancing]: No Consumer Found For Group Id: %d", groupId)
			return nil
		}

		// Get All Partitions For The Given Topic
		partitions, queryErr := repository.GetPartitionByTopicId(topicId)
		if queryErr != nil {
			utils.Logger.ErrorF(queryErr, "[CG Rebalancing]: Error While Finding Partitions For Topic: %d", topicId)
			return queryErr
		}

		allPartitions := *partitions
		if len(allPartitions) == 0 {
			utils.Logger.InfoF("[CG Rebalancing]: No Partitions Found For Topic: %d", topicId)
			return nil
		}

		ids := make([]uint, 0)
		for _, consumer := range allConsumers {
			ids = append(ids, consumer.Id)
		}

		// Delete Existing ConsumerAssignment
		if queryErr := repository.DeleteAllConsumerAssignment(ids); queryErr != nil {
			return queryErr
		}

		// Lets Re-assign All Partitions In Round Robin Fashion
		for index, partition := range allPartitions {

			consumer := allConsumers[index%count]

			instance := model.ConsumerAssignment{
				ConsumerId:  consumer.Id,
				PartitionId: partition.Id,
			}

			if _, queryErr := repository.AssignPartitionToConsumer(instance); queryErr != nil {
				return queryErr
			}
		}

		return nil

	})

}
