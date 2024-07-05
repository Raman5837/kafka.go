package types

// GetTopic API Request Payload
type GetTopicRequestEntity struct {
	Name string `json:"name"`
}

// CreateTopic API Request Payload
type CreateTopicRequestEntity struct {
	Name string `json:"name"`
}

// ProduceMessage API Request Payload
type ProduceMessageRequestEntity struct {
	Value       interface{} `json:"value"`
	TopicId     uint64      `json:"topic_id"`
	PartitionId *uint64     `json:"partition_id"`
}
