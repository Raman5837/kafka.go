package service

import (
	"github.com/Raman5837/kafka.go/app/model"
	"github.com/Raman5837/kafka.go/app/repository"
	"github.com/Raman5837/kafka.go/app/types"
)

// Topic Service Object
type TopicService struct {
}

// Returns New Topic Service Instance
func NewTopicService() *TopicService {
	return &TopicService{}
}

// Get The Topic With Given Payload
func (service *TopicService) GetTopic(payload *types.GetTopicRequestEntity) (Topic *types.GetTopic, exception error) {

	return repository.GetTopic(payload.Name)

}

// Add New Topic With Given Payload
func (service *TopicService) AddNewTopic(payload *types.CreateTopicRequestEntity) (Topic *types.GetTopic, exception error) {

	instance, queryErr := repository.CreateTopic(&model.Topic{Name: payload.Name})
	newTopic := types.GetTopic{Id: instance.ID, Name: instance.Name, CreatedAt: instance.CreatedAt}
	return &newTopic, queryErr

}
