package types

// Get Topic API Request Payload
type GetTopicRequestEntity struct {
	Name string `json:"name"`
}

// Create New Topic API Request Payload
type CreateTopicRequestEntity struct {
	Name string `json:"name"`
}

// Add New Consumer API Request Payload
type AddNewConsumerRequestEntity struct {
	GroupID uint64 `json:"group_id"`
	TopicId uint64 `json:"topic_id"`
}

// Produce Message API Request Payload
type ProduceMessageRequestEntity struct {
	Value       interface{} `json:"value"`
	TopicId     uint64      `json:"topic_id"`
	PartitionId *uint64     `json:"partition_id"`
}

// Get Message To Consumer API Request Payload
type GetMessageToConsumeRequestEntity struct {
	Value       interface{} `json:"value"`
	Offset      uint64      `json:"offset"`
	GroupID     uint64      `json:"group_id"`
	ConsumerId  uint64      `json:"consumer_id"`
	PartitionId uint64      `json:"partition_id"`
}
