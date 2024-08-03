package interfaces

import "github.com/Raman5837/kafka.go/app/types"

// PartitionAssigner Interface
type PartitionAssignerInterface interface {
	Next(topicID uint) (*types.GetPartition, error)
}
