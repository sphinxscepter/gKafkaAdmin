package vo

type KafkaTopic struct {
	Topic          string `json:"topic"`
	PartitionCount int    `json:"partitionCount"`
	Partitions     []int  `json:"partitions"`
}
