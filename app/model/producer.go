package model

import (
	"fmt"

	"github.com/Raman5837/kafka.go/base/model"
)

// Represents A Kafka Topic
type Topic struct {
	model.AbstractModel
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:256;uniqueIndex;not null"`
}

func (instance Topic) TableName() string {
	return "topics"
}

func (instance Topic) String() string {
	return fmt.Sprintf("Topic: --> %s", instance.Name)
}

// Represents A Kafka Partition
type Partition struct {
	model.AbstractModel
	ID          uint `gorm:"primaryKey"`
	PartitionId uint `gorm:"not null"`

	TopicId uint  `gorm:"not null"`
	Topic   Topic `gorm:"foreignKey:TopicID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (instance Partition) TableName() string {
	return "partitions"
}

func (instance Partition) String() string {
	return fmt.Sprintf("Partition Id: %d --> Of Topic %s", instance.PartitionId, instance.Topic.Name)
}

// Stores A Kafka Message In A Partition Of A Topic
type Message struct {
	model.AbstractModel
	ID uint `gorm:"primaryKey"`

	Offset int64       `gorm:"not null"`
	Value  interface{} `gorm:"type:json"`

	PartitionId uint      `gorm:"not null"`
	Partition   Partition `gorm:"foreignKey:PartitionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (instance Message) TableName() string {
	return "messages"
}

func (instance Message) String() string {
	return fmt.Sprintf("Message: %d --> Of Partition %d", instance.ID, instance.PartitionId)
}

// Stores The Last Assigned Partition To A Topic
type LastAssignedPartition struct {
	model.AbstractModel
	ID uint `gorm:"primaryKey"`

	TopicID uint  `gorm:"unique;not null"`
	Topic   Topic `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	PartitionId uint      `gorm:"not null"`
	Partition   Partition `gorm:"foreignKey:PartitionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (instance LastAssignedPartition) TableName() string {
	return "assignments"
}

func (instance LastAssignedPartition) String() string {
	return fmt.Sprintf("Topic: %s's Last Partition Was --> %d", instance.Topic.Name, instance.PartitionId)
}
