package service

import (
	"github.com/Raman5837/kafka.go/app/model"
	"github.com/Raman5837/kafka.go/app/repository"
	"github.com/Raman5837/kafka.go/app/types"
	"github.com/Raman5837/kafka.go/app/utils"
)

// Consumer Service
type ConsumerService struct {
}

// Returns New Consumer Service Instance
func NewConsumerService() *ConsumerService {
	return &ConsumerService{}
}

// Add A New Consumer
func (service *ConsumerService) AddNewConsumer(payload *types.AddNewConsumerRequestEntity) (Consumer *types.GetConsumer, exception error) {

	instance, queryErr := repository.AddNewConsumer(&model.Consumer{GroupID: payload.GroupID})
	newConsumer := types.GetConsumer{GroupID: instance.Group.ID, ConsumerID: instance.ConsumerID}

	// Let's Perform Re-balancing
	rebalance := utils.NewRebalanceService()
	if rebalanceErr := rebalance.RebalanceGroup(payload.GroupID, payload.TopicId); rebalanceErr != nil {
		return &newConsumer, rebalanceErr
	}

	return &newConsumer, queryErr
}

// Get Message To Consume For Requested Consumer
func (service *ConsumerService) GetMessages(payload *types.GetMessageToConsumeRequestEntity) (Messages *types.ConsumeMessageResponseEntity, exception error) {

	// Validate Consumer
	if _, consumerErr := repository.GetConsumer(payload.ConsumerId, payload.GroupID); consumerErr != nil {
		return nil, consumerErr
	}

	// Fetch New Message
	messages, messageErr := repository.GetMessages(payload.PartitionId, payload.Offset)
	if messageErr != nil {
		return nil, messageErr
	}

	allMessages := *messages
	count := len(allMessages)

	// Update Offset
	if count > 0 {
		newOffset := allMessages[count-1].Offset + 1
		if updateErr := repository.UpdateOffset(payload.ConsumerId, payload.PartitionId, newOffset); updateErr != nil {
			return nil, updateErr
		}
	}

	response := types.ConsumeMessageResponseEntity{Messages: allMessages}
	return &response, messageErr
}
