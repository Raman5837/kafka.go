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
	GroupId uint `json:"group_id"`
	TopicId uint `json:"topic_id"`
}

// Add New Consumer Group API Request Payload
type AddNewConsumerGroupRequestEntity struct {
	Name string `json:"name"`
}

// Produce Message API Request Payload
type ProduceMessageRequestEntity struct {
	Value       interface{} `json:"value"`
	TopicId     uint        `json:"topic_id"`
	PartitionId *uint       `json:"partition_id"`
}

// Get Message To Consumer API Request Payload
type GetMessageToConsumeRequestEntity struct {
	Value       interface{} `json:"value"`
	GroupId     uint        `json:"group_id"`
	ConsumerId  uint        `json:"consumer_id"`
	PartitionId uint        `json:"partition_id"`
}
