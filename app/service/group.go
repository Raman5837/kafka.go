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

// Consumer Service Object
type ConsumerGroupService struct {
	Config config.ConsumerGroupConfig

	RebalanceLock    sync.Mutex // Mutex To Synchronize Access To Shared State
	LastRebalancedAt time.Time  // Time At Which Last Re-balancing Was Done.
	// RebalancingTimer  *time.Timer   // Timer To Trigger Re-balancing After Specific Delay(RebalanceInterval)
	RebalancingSignal chan struct{} // A Channel To Signal Re-balancing Operation
}

// Returns New ConsumerGroup Service Instance
func NewConsumerGroupService(config config.ConsumerGroupConfig) *ConsumerGroupService {
	instance := &ConsumerGroupService{
		Config:            config,
		RebalancingSignal: make(chan struct{}, 1), // Buffering Channel To Avoid Blocking
	}

	// Start Periodic-Rebalancing In Separate Goroutines
	go instance.PeriodicRebalancing()

	return instance
}

// Add A New Consumer Group With Given Payload
func (service *ConsumerGroupService) AddNewConsumerGroup(payload *types.AddNewConsumerGroupRequestEntity) (ConsumerGroup *types.GetConsumerGroup, exception error) {

	instance, queryErr := repository.AddNewConsumerGroup(&model.ConsumerGroup{Name: payload.Name})
	if queryErr != nil {
		return nil, queryErr
	}

	newConsumer := types.GetConsumerGroup{Id: instance.ID, Name: instance.Name, CreatedAt: instance.CreatedAt}

	service.RebalanceLock.Lock()
	shouldRebalance := time.Since(service.LastRebalancedAt) >= service.Config.RebalanceInterval
	service.RebalanceLock.Unlock()

	// Trigger The Signal For Rebalancing
	if shouldRebalance {
		service.RebalancingSignal <- struct{}{}
	}

	return &newConsumer, nil
}

// Performs Periodic Re-balancing Based On Set Consumer Group Configuration
func (service *ConsumerGroupService) PeriodicRebalancing() {

	ticker := time.NewTicker(service.Config.RebalanceInterval)
	defer ticker.Stop()

	for {
		select {
		case <-service.RebalancingSignal:
			service.ExecuteRebalancing()

		case <-ticker.C:
			service.ExecuteRebalancing()
		}
	}
}

// Rebalance Consumer Groups
func (service *ConsumerGroupService) ExecuteRebalancing() {

	service.RebalanceLock.Lock()
	defer service.RebalanceLock.Unlock()

	// Fetch All Consumers Of The Groups, Create A Map Of Group Id And Topic Ids
	assignedConsumers, queryErr := repository.GetAllAssignedConsumers(nil)

	if queryErr != nil {
		base.Logger.Error(queryErr, "[CG Rebalancing]: Something Went Wrong While Fetching All Consumer Groups")
		return
	}

	allConsumers := *assignedConsumers
	payload := make(map[uint]map[uint]struct{}, len(allConsumers))

	for _, consumer := range allConsumers {

		if _, exists := payload[consumer.GroupId]; !exists {
			payload[consumer.GroupId] = make(map[uint]struct{})
		}

		// Add TopicId To The Set For The Corresponding Group Id
		payload[consumer.GroupId][consumer.TopicId] = struct{}{}

	}

	mapper := make(map[uint][]uint)
	for groupId, topicMap := range payload {

		topicsSlice := make([]uint, 0, len(topicMap))

		for topicId := range topicMap {
			topicsSlice = append(topicsSlice, topicId)
		}
		mapper[groupId] = topicsSlice
	}

	var waitGroup *sync.WaitGroup
	errChannels := make(chan error, len(mapper)*len(allConsumers))

	instance := utils.NewRebalanceService()

	for groupId, topicIds := range mapper {

		for _, topicId := range topicIds {

			waitGroup.Add(1)
			go func(groupId uint, topicId uint) {

				defer waitGroup.Done()

				if rebalanceErr := instance.RebalanceGroup(groupId, topicId); rebalanceErr != nil {
					errChannels <- rebalanceErr
				}

			}(groupId, topicId)

		}
	}

	waitGroup.Wait()
	close(errChannels)

	for someErr := range errChannels {
		base.Logger.Error(someErr, someErr.Error())
	}

	// Update LastRebalancedAt Time
	service.LastRebalancedAt = time.Now()

}
