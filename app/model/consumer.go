package model

import (
	"fmt"

	"github.com/Raman5837/kafka.go/base/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Represents Kafka Consumer Groups
type ConsumerGroup struct {
	model.AbstractModel
	ID uint `gorm:"primaryKey"`

	IsActive bool   `gorm:"not null; default:true"`
	Name     string `gorm:"size:256;uniqueIndex;not null"`
}

func (instance ConsumerGroup) TableName() string {
	return "groups"
}

func (instance ConsumerGroup) String() string {
	return fmt.Sprintf("Consumer Group: --> %s", instance.Name)
}

// Represents Kafka Consumers
type Consumer struct {
	model.AbstractModel
	ID uint `gorm:"primaryKey"`

	ConsumerId uuid.UUID `gorm:"type:uuid;unique;not null"`

	GroupID uint          `gorm:"not null"`
	Group   ConsumerGroup `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// Add Default Value For Column `ConsumerId`
func (instance *Consumer) BeforeCreate(transaction *gorm.DB) error {
	instance.ConsumerId = uuid.New()
	return nil
}

func (instance Consumer) TableName() string {
	return "consumers"
}

func (instance Consumer) String() string {
	return fmt.Sprintf("Consumer: %s Of Group: %s", instance.ConsumerId, instance.Group.Name)
}

// Stores Current Offset Of The Consumer In The Associated Partition
type Offset struct {
	model.AbstractModel
	ID uint `gorm:"primaryKey"`

	Number     uint64   `gorm:"not null"`
	ConsumerId uint     `gorm:"not null"`
	Consumer   Consumer `gorm:"foreignKey:ConsumerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	PartitionId uint      `gorm:"not null"`
	Partition   Partition `gorm:"foreignKey:PartitionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (instance Offset) TableName() string {
	return "offsets"
}

func (instance Offset) String() string {
	return fmt.Sprintf("Offset Having Number: %d For Consumer: %d Of Partition: %d", instance.Number, instance.ConsumerId, instance.PartitionId)
}

// To Manage Consumer Rebalancing, Partitions Assignments And All
type ConsumerAssignment struct {
	model.AbstractModel
	ID uint `gorm:"primaryKey"`

	ConsumerId uint     `gorm:"not null"`
	Consumer   Consumer `gorm:"foreignKey:ConsumerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	PartitionId uint      `gorm:"not null"`
	Partition   Partition `gorm:"foreignKey:PartitionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (instance ConsumerAssignment) TableName() string {
	return "rebalancing"
}

func (instance ConsumerAssignment) String() string {
	return fmt.Sprintf("Management Of Consumer: %d Of Group: %d And Partition: %d", instance.ConsumerId, instance.Consumer.GroupID, instance.PartitionId)
}
