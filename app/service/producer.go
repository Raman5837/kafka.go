package service

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
	"unsafe"

	interfaces "github.com/Raman5837/kafka.go/app/interface"
	"github.com/Raman5837/kafka.go/app/model"
	"github.com/Raman5837/kafka.go/app/repository"
	"github.com/Raman5837/kafka.go/app/types"
	"github.com/Raman5837/kafka.go/base/config"
	base "github.com/Raman5837/kafka.go/base/utils"
)

// Produce Service Object
type ProducerService struct {
	Config            config.ProducerConfig
	partitionAssigner interfaces.PartitionAssignerInterface

	BufferLock sync.Mutex      // Mutex For Thread Safety
	Buffer     []model.Message // Buffer Of Messages

	LastFlushedAt time.Time
	FlushTimer    *time.Timer   // Timer To Trigger Flush After Specific Delay(LingerMS)
	FlushSignal   chan struct{} // A Channel To Trigger Immediate Flushing Of Buffer
}

// Returns New Producer Service Instance
func NewProducerService(assigner interfaces.PartitionAssignerInterface, config config.ProducerConfig) *ProducerService {
	instance := &ProducerService{
		Config:            config,
		partitionAssigner: assigner,

		FlushSignal: make(chan struct{}),
		Buffer:      make([]model.Message, 0, config.BatchSize),
	}

	// Start Periodic Flushing Of Message
	go instance.PeriodicFlush()

	return instance
}

// A Helper Function To Calculate Estimated Size Of Message (For Buffer Management)
func (service *ProducerService) estimateMessageSize(message model.Message) int {

	var valueSize int

	// Estimate Size Of The Value Field
	switch value := message.Value.(type) {
	case string:
		valueSize = len(value)

	case []byte:
		valueSize = len(value)

	default:
		// If The Value Is Not A Recognized Type, Attempt To Serialize It To JSON
		data, parsingErr := json.Marshal(value)

		if parsingErr != nil {
			valueSize = 0
		} else {
			valueSize = len(data)
		}
	}

	fixedSize := int(unsafe.Sizeof(message.ID)) + int(unsafe.Sizeof(message.Offset)) + int(unsafe.Sizeof(message.PartitionId))

	return fixedSize + valueSize
}

// Add Or Produce New Message (Adds The Message In Buffer And Based On Configuration, Flush It)
func (service *ProducerService) AddNewMessage(payload *types.ProduceMessageRequestEntity) (instance *types.ProduceMessageResponseEntity, exception error) {

	var partitionNotFoundErr error
	var partition *types.GetPartition
	memoryLimitErr := "Buffer Memory Limit Exceeded"

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

	messagePayload := model.Message{Value: payload.Value, Offset: offsetNumber, PartitionId: partition.PartitionId}

	if estimatedMessageSize := service.estimateMessageSize(messagePayload); estimatedMessageSize > int(service.Config.BufferMemory) {
		return nil, errors.New(memoryLimitErr)
	}

	// Add Message To Buffer
	service.BufferLock.Lock()
	service.Buffer = append(service.Buffer, messagePayload)

	shouldFlush := len(service.Buffer) >= service.Config.BatchSize

	// Trigger Flush Immediately, If Batch Size Reached (As Set In Configuration)
	if shouldFlush {
		service.FlushSignal <- struct{}{}
	} else {

		// Reset The Flush Timer
		if service.FlushTimer != nil {
			service.FlushTimer.Stop()
		}

		// Trigger Flush, Once LingerMS Time Reach (As Set In Configuration)
		service.FlushTimer = time.AfterFunc(service.Config.LingerMS, func() {
			service.FlushSignal <- struct{}{}
		})
	}

	// Remove Lock
	service.BufferLock.Unlock()

	if shouldFlush {

		if flushErr := service.FlushBuffer(); flushErr != nil {
			return nil, flushErr
		}
	}

	response := &types.ProduceMessageResponseEntity{
		Value:       messagePayload.Value,
		Offset:      messagePayload.Offset,
		PartitionId: messagePayload.PartitionId,
		TopicId:     messagePayload.Partition.TopicId,
	}

	return response, nil

}

// Flush All The Buffers In The DB
func (service *ProducerService) FlushBuffer() (exception error) {

	service.BufferLock.Lock()
	defer service.BufferLock.Unlock()

	if len(service.Buffer) == 0 {
		return nil
	}

	// TODO: Handle Compression Logic

	allErrors := make([]error, 0)
	flushedBuffers := make([]*model.Message, 0)

	// TODO: Implement Retry Logic
	for _, message := range service.Buffer {

		messageCopy := message // To Avoid Implicit Memory Aliasing
		newMessage, queryErr := repository.AddNewMessage(&messageCopy)

		if queryErr != nil {
			allErrors = append(allErrors, queryErr)
		} else {
			flushedBuffers = append(flushedBuffers, newMessage)
		}
	}

	for _, queryErr := range allErrors {
		base.Logger.ErrorF(queryErr, "DB Error While Flushing Buffers")
	}

	for _, message := range flushedBuffers {
		base.Logger.InfoF("Added Message: %v In Partition: %v", message.Value, message.Partition.String())
	}

	// Reset Buffer
	service.Buffer = service.Buffer[:0]
	// Updated Last Flushed Time
	service.LastFlushedAt = time.Now()

	return nil
}

// Periodically Flush All Buffered Messages
func (service *ProducerService) PeriodicFlush() (exception error) {

	// TODO: Make It Efficient.
	ticker := time.NewTicker(service.Config.FlushInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Periodic Flush Based On Flush Interval
			if flushErr := service.FlushBuffer(); flushErr != nil {
				base.Logger.Error(flushErr, "Error While Periodically Flushing Buffer.")
			}

		case <-service.FlushSignal:
			// Immediate Flush Triggered By Signal
			if flushErr := service.FlushBuffer(); flushErr != nil {
				base.Logger.Error(flushErr, "Error While Flushing Buffer Due to Signal.")
			}
		}
	}
}

// Get Partition With Topic Id And Partition Id
func (service *ProducerService) GetPartition(topicId uint, partitionId uint) (Partition *types.GetPartition, exception error) {

	return repository.GetPartition(topicId, partitionId)
}
