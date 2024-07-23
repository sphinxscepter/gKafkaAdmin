package kafkaService

import (
	"gKafkaAdmin/internal/module/vo"
	"gKafkaAdmin/internal/zlog"
	"sort"

	"github.com/segmentio/kafka-go"
)

func ListAllTopic(topicName string) map[string]vo.KafkaTopic {
	conn, err := kafka.Dial("tcp", "172.22.33.102:29092")
	if err != nil {
		zlog.Info(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		zlog.Info(err.Error())
	}

	kafkaTopicInfo := make(map[string]vo.KafkaTopic)

	for _, p := range partitions {
		if p.Topic != "__consumer_offsets" {
			// 判断对应topic信息是否在map中存在
			topicData, exist := kafkaTopicInfo[p.Topic]
			// zlog.Info(p.Topic)
			// 如果topic存在，则更新partition信息
			if exist {
				topicData.PartitionCount += 1
				topicData.Partitions = append(topicData.Partitions, p.ID)
				sort.Ints(topicData.Partitions)
				kafkaTopicInfo[p.Topic] = topicData
			} else {
				// 如果topic存在，则初始化topic信息
				kafkaTopic := new(vo.KafkaTopic)
				kafkaTopic.Topic = p.Topic
				kafkaTopic.PartitionCount = 1
				partitions := []int{p.ID}
				kafkaTopic.Partitions = partitions
				kafkaTopicInfo[p.Topic] = *kafkaTopic
			}
		}
	}
	// zlog.Info(json.Marshal(kafkaTopicInfo))
	return kafkaTopicInfo
}
