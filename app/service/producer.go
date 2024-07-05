package service

import (
	interfaces "github.com/Raman5837/kafka.go/app/interface"
	"github.com/Raman5837/kafka.go/app/model"
	"github.com/Raman5837/kafka.go/app/repository"
	"github.com/Raman5837/kafka.go/app/types"
)

// Produce Service Object
type ProducerService struct {
	partitionAssigner interfaces.PartitionAssignerInterface
}

// Returns New Producer Service Instance
func NewProducerService(assigner interfaces.PartitionAssignerInterface) *ProducerService {
	return &ProducerService{partitionAssigner: assigner}
}

// Add Or Produce New Message
func (service *ProducerService) AddNewMessage(payload *types.ProduceMessageRequestEntity) (instance *types.ProduceMessageResponseEntity, exception error) {

	var partition *types.GetPartition
	var partitionNotFoundErr error

	if payload.PartitionId == nil {
		partition, partitionNotFoundErr = service.partitionAssigner.Next(payload.TopicId)
	} else {
		partition, partitionNotFoundErr = service.GetPartition(payload.TopicId, *payload.PartitionId)
	}

	if partitionNotFoundErr != nil {
		return nil, partitionNotFoundErr
	}

	offsetNumber, queryErr := repository.GetAllMessageCount()

	if queryErr != nil {
		return nil, queryErr
	}

	messagePayload := model.Message{Value: payload.Value, Offset: offsetNumber, PartitionID: partition.PartitionId}
	newMessage, queryErr := repository.AddNewMessage(&messagePayload)
	if queryErr != nil {
		return nil, queryErr
	}

	response := &types.ProduceMessageResponseEntity{
		Value:       newMessage.Value,
		Offset:      newMessage.Offset,
		PartitionId: &newMessage.PartitionID,
		TopicId:     newMessage.Partition.TopicID,
	}

	return response, nil

}

// Get Partition With Topic Id And Partition Id
func (service *ProducerService) GetPartition(topicId uint64, partitionId uint64) (Partition *types.GetPartition, exception error) {

	return repository.GetPartition(topicId, partitionId)
}
