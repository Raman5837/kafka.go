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
func (instance *RebalanceService) RebalanceGroup(groupId uint64, topicId uint64) (exception error) {

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
			return nil
		}

		// Get All Partitions For The Given Topic
		partitions, queryErr := repository.GetPartitionByTopicId(topicId)
		if queryErr != nil {
			return queryErr
		}

		allPartitions := *partitions
		if len(allPartitions) == 0 {
			return nil
		}

		ids := make([]uint64, 0)
		for _, consumer := range allConsumers {
			ids = append(ids, uint64(consumer.Id))
		}

		// Delete Existing ConsumerAssignment
		if queryErr := repository.DeleteAllConsumerAssignment(ids); queryErr != nil {
			return queryErr
		}

		// Lets Re-assign All Partitions In Round Robin Fashion
		for index, partition := range allPartitions {

			consumer := allConsumers[index%count]

			instance := model.ConsumerAssignment{
				ConsumerID:  uint64(consumer.Id),
				PartitionID: uint64(partition.Id),
			}

			if _, queryErr := repository.AssignPartitionToConsumer(instance); queryErr != nil {
				return queryErr
			}
		}

		return nil

	})

}
