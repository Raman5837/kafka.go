package service

import (
	"sync"
	"time"

	"github.com/Raman5837/kafka.go/app/model"
	"github.com/Raman5837/kafka.go/app/repository"
	"github.com/Raman5837/kafka.go/app/types"
	"github.com/Raman5837/kafka.go/app/utils"
	"github.com/Raman5837/kafka.go/base/config"
	base "github.com/Raman5837/kafka.go/base/utils"
)

// Kafka Consumer Service
type ConsumerService struct {
	LastHeartBeatAt time.Time
	LastPollTime    time.Time
	OffsetMutex     sync.Mutex
	Config          config.ConsumerConfig
}

// Returns New Consumer Service Instance
func NewConsumerService(config config.ConsumerConfig) *ConsumerService {
	return &ConsumerService{Config: config, LastHeartBeatAt: time.Now(), LastPollTime: time.Now()}
}

// Add A New Consumer
func (service *ConsumerService) AddNewConsumer(payload *types.AddNewConsumerRequestEntity) (Consumer *types.GetConsumer, exception error) {

	instance, queryErr := repository.AddNewConsumer(&model.Consumer{GroupID: payload.GroupId})

	if queryErr != nil {
		return nil, queryErr
	}

	newConsumer := types.GetConsumer{GroupId: instance.Group.ID, ConsumerId: instance.ConsumerId}

	// Let's Perform Re-balancing
	rebalance := utils.NewRebalanceService()
	if rebalanceErr := rebalance.RebalanceGroup(payload.GroupId, payload.TopicId); rebalanceErr != nil {
		return &newConsumer, rebalanceErr
	}

	return &newConsumer, queryErr
}

// Get Committed Offset For Give Consumer And Partition
func (service *ConsumerService) GetCommittedOffset(consumerId uint, partitionId uint) (Offset *types.GetOffset, exception error) {

	return repository.GetConsumerOffset(consumerId, partitionId)
}

// Get Message To Consume For Requested Consumer
func (service *ConsumerService) GetMessages(payload *types.GetMessageToConsumeRequestEntity) (Messages *types.ConsumeMessageResponseEntity, exception error) {

	// Validate Consumer
	if _, consumerErr := repository.GetConsumer(payload.ConsumerId, payload.GroupId); consumerErr != nil {
		return nil, consumerErr
	}

	committedOffset, queryErr := service.GetCommittedOffset(payload.ConsumerId, payload.PartitionId)

	nextOffset := func() uint64 {
		// Let's Consume From Beginning If `committedOffset` Is nil Or Something Breaks While Fetching The Committed Offset.
		if queryErr != nil || committedOffset == nil {
			return 0
		}
		return committedOffset.Number
	}()

	baseConfig := config.GetKafkaConfig()

	// Fetch New Message
	messages, messageErr := repository.GetMessages(payload.PartitionId, nextOffset, baseConfig.Consumer.FetchSize)
	if messageErr != nil {
		return nil, messageErr
	}

	allMessages := *messages
	count := len(allMessages)

	// Update Offset
	if count > 0 {
		service.OffsetMutex.Lock()
		newOffset := allMessages[count-1].Offset + 1
		if updateErr := repository.UpdateOffset(payload.ConsumerId, payload.PartitionId, newOffset); updateErr != nil {
			service.OffsetMutex.Unlock()
			return nil, updateErr
		}
		service.OffsetMutex.Unlock()
	}

	response := types.ConsumeMessageResponseEntity{Messages: allMessages}
	return &response, messageErr
}

// Start Polling Messages Periodically Based On Set Configuration
func (service *ConsumerService) PollMessage(payload *types.GetMessageToConsumeRequestEntity) (Message *types.ConsumeMessageResponseEntity, exception error) {

	// TODO: Make If Efficient (Let's Avoid Using Ticker)
	ticker := time.NewTicker(service.Config.PollInterval)
	defer ticker.Stop()

	for range ticker.C {

		service.LastPollTime = time.Now()
		message, queryErr := service.GetMessages(payload)

		if queryErr != nil {
			return nil, queryErr
		}

		if len(message.Messages) > 0 {
			return message, nil
		}

	}

	return nil, nil

}

// Performs HearBeat Operation
func (service *ConsumerService) CheckHeartBeat() (exception error) {

	// TODO: Make If Efficient (Let's Avoid Using Ticker)
	ticker := time.NewTicker(service.Config.PollInterval)
	defer ticker.Stop()

	for range ticker.C {
		service.LastHeartBeatAt = time.Now()
		// TODO: Implement HeartBeat Logic
		base.Logger.InfoF("[Consumer HeartBeat]: LastHeartBeat Checked Was At: %v", service.LastHeartBeatAt)

	}

	return nil
}
