package types

import (
	"time"

	"github.com/google/uuid"
)

// To Store Topic Object From Repository
type GetTopic struct {
	Id        uint      `gorm:"column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

// To Store Partition Object From Repository
type GetPartition struct {
	Id        uint      `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`

	TopicID     uint `gorm:"column:topic_id" json:"topic_id"`
	PartitionId uint `gorm:"column:partition_id" json:"partition_id"`
}

// To Store Message Object From Repository
type GetMessage struct {
	Id        uint      `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`

	Value       interface{} `gorm:"column:value" json:"value"`
	Offset      uint64      `gorm:"column:offset" json:"offset"`
	PartitionId uint        `gorm:"column:partition_id" json:"partition_id"`
}

// To Store LastAssignedPartition Object From Repository
type GetLastAssignedPartition struct {
	Id        uint      `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`

	TopicId     uint `gorm:"column:topic_id" json:"topic_id"`
	PartitionId uint `gorm:"column:partition_id" json:"partition_id"`
}

// To Store Consumer Object From Repository
type GetConsumer struct {
	Id        uint      `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`

	GroupId    uint      `gorm:"column:group_id" json:"group_id"`
	ConsumerId uuid.UUID `gorm:"column:consumer_id" json:"consumer_id"`
}

// To Store ConsumerGroup Object From Repository
type GetConsumerGroup struct {
	Id        uint      `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`

	Name     string `gorm:"column:name" json:"name"`
	IsActive bool   `gorm:"column:is_active" json:"is_active"`
}

// To Store Offset Object From Repository
type GetOffset struct {
	Id        uint      `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`

	Number      uint64 `gorm:"column:number" json:"number"`
	ConsumerId  uint   `gorm:"column:consumer_id" json:"consumer_id"`
	PartitionId uint   `gorm:"column:partition_id" json:"partition_id"`
}

// To Store ConsumerAssignment Object From Repository
type GetConsumerAssignment struct {
	Id        uint      `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`

	GroupId     uint `gorm:"column:group_id" json:"group_id"`
	TopicId     uint `gorm:"column:topic_id" json:"topic_id"`
	ConsumerId  uint `gorm:"column:consumer_id" json:"consumer_id"`
	PartitionId uint `gorm:"column:partition_id" json:"partition_id"`
}
