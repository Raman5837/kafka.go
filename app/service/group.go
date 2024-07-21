package service

import (
	"github.com/Raman5837/kafka.go/app/model"
	"github.com/Raman5837/kafka.go/app/repository"
	"github.com/Raman5837/kafka.go/app/types"
)

// Consumer Service Object
type ConsumerGroupService struct {
}

// Returns New ConsumerGroup Service Instance
func NewConsumerGroupService() *ConsumerGroupService {
	return &ConsumerGroupService{}
}

// Add A New Consumer Group With Given Payload
func (service *ConsumerGroupService) AddNewConsumerGroup(payload *types.AddNewConsumerGroupRequestEntity) (ConsumerGroup *types.GetConsumerGroup, exception error) {

	instance, queryErr := repository.AddNewConsumerGroup(&model.ConsumerGroup{Name: payload.Name})
	if queryErr != nil {
		return nil, queryErr
	}

	newConsumer := types.GetConsumerGroup{Id: instance.ID, Name: instance.Name, CreatedAt: instance.CreatedAt}
	return &newConsumer, nil
}
