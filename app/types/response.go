package types

// ProduceMessage Service Response
type ProduceMessageResponseEntity struct {
	Offset      int64       `json:"offset"`
	Value       interface{} `json:"value"`
	TopicId     uint        `json:"topic_id"`
	PartitionId uint        `json:"partition_id"`
}

// Consumer Message Response Entity
type ConsumeMessageResponseEntity struct {
	Messages []GetMessage `json:"messages"`
}
