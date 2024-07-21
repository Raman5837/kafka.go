package types

// ProduceMessage Service Response
type ProduceMessageResponseEntity struct {
	Offset      int64       `json:"offset"`
	Value       interface{} `json:"value"`
	TopicId     uint64      `json:"topic_id"`
	PartitionId uint64      `json:"partition_id"`
}

// Consumer Message Response Entity
type ConsumeMessageResponseEntity struct {
	Messages []GetMessage `json:"messages"`
}
