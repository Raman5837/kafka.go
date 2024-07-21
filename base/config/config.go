package config

import (
	"sync"
	"time"
)

// Contains All Kafka Related Configurations
type KafkaConfiguration struct {
	Producer      ProducerConfig
	Consumer      ConsumerConfig
	ConsumerGroup ConsumerGroupConfig

	Polling          PollingConfig
	HeartBeat        HeartBeatConfig
	MessageRetention MessageRetentionConfig
}

// Configuration Related To Kafka Producers
type ProducerConfig struct {
	BatchSize    int    // Number Of Messages To Batch
	BufferMemory int64  // Buffer Memory Size In Bytes
	Compression  string // Compression Type (Eg. none, gzip)

	LingerMS      time.Duration // Time To Wait Before Sending A Batch Of Messages
	FlushInterval time.Duration // Interval To Flush The Messages
}

// Configuration Related To Kafka Consumers
type ConsumerConfig struct {
	FetchSize int // Number Of Messages To Fetch In One Request
	BatchSize int // Number Of Messages To Process In A Batch

	PollInterval   time.Duration // Time Between Polling Requests
	SessionTimeout time.Duration // Session Timeout For Consumers

	AutoCommit        bool          // Whether To Auto Commit Offsets
	FetchMinBytes     int           // Minimum Bytes Of Data To Fetch
	FetchMaxWaitMS    time.Duration // Maximum Wait Time To Fetch Data
	HeartBeatInterval time.Duration // Time Between Heartbeat Requests
}

// Configuration Related To Kafka Consumer Groups
type ConsumerGroupConfig struct {
	OffsetReset       string        // Offset Reset Policy (Eg. earliest, latest)
	RebalanceInterval time.Duration // Time Interval To Triggers Rebalancing
}

// Configuration Related To Kafka Message Polling
type PollingConfig struct {
	Interval time.Duration // Interval Between Polling Requests
}

// Configuration Related To Kafka HeartBeat
type HeartBeatConfig struct {
	Interval time.Duration // Heartbeat Interval
}

// Configuration Related To Kafka Message Retention
type MessageRetentionConfig struct {
	RetentionPeriod time.Duration // Retention Duration For Messages
}

var singleton sync.Once
var baseConfig KafkaConfiguration

// Initialize Base Kafka Configuration
func InitKafkaConfig() {

	singleton.Do(
		func() {
			baseConfig = KafkaConfiguration{
				Producer: ProducerConfig{
					BatchSize:     100,
					Compression:   "none",
					BufferMemory:  33554432,
					FlushInterval: 5 * time.Second,
					LingerMS:      500 * time.Millisecond,
				},
				Consumer: ConsumerConfig{
					FetchMinBytes:     1,
					FetchSize:         100,
					BatchSize:         10,
					AutoCommit:        true,
					SessionTimeout:    60 * time.Second,
					HeartBeatInterval: 10 * time.Second,
					PollInterval:      500 * time.Millisecond,
					FetchMaxWaitMS:    500 * time.Millisecond,
				},
				ConsumerGroup: ConsumerGroupConfig{
					OffsetReset:       "earliest",
					RebalanceInterval: 30 * time.Minute,
				},
				Polling: PollingConfig{
					Interval: 100 * time.Millisecond,
				},
				HeartBeat: HeartBeatConfig{
					Interval: 10 * time.Second,
				},
				MessageRetention: MessageRetentionConfig{
					RetentionPeriod: 24 * time.Hour,
				},
			}
		},
	)

}

// Returns Base Kafka Configuration
func GetKafkaConfig() *KafkaConfiguration {
	return &baseConfig
}
