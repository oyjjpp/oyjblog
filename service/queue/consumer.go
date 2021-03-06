package queue

import (
	"context"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	ctx context.Context
}

// 校验接口是否被实现
var _ sarama.ConsumerGroupHandler = &Consumer{}

func (consumer *Consumer) Setup(s sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) Cleanup(s sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case data, ok := <-claim.Messages():
			log.Printf("ConsumeClaim data:%v,ok:%v \n", data, ok)
			if !ok {
				log.Printf("ConsumeClaim fail:%v \n", data)
				time.Sleep(10 * time.Second)
				continue
			}
			message := string(data.Value)
			consumer.handle(message)

			// 提交消息 标记消息已经被消费
			session.MarkMessage(data, "")
		case <-consumer.ctx.Done():
			log.Printf("ConsumeClaim consume_message exit")
			return nil
		}
	}
}

// 处理消息
func (consumer *Consumer) handle(message string) {
	log.Printf("consumer message:%v \n", message)
}

// 拦截器
type DelayInterceptor struct {
	ctx context.Context
}

var _ sarama.ConsumerInterceptor = &DelayInterceptor{}

func (delay *DelayInterceptor) OnConsume(message *sarama.ConsumerMessage) {
	// message.Timestamp
	// message.Headers
}

// kafka消费组初始化
func ConsumerGroupInit(ctx context.Context) {
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_0
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// 注册拦截器
	delay := new(DelayInterceptor)
	config.Consumer.Interceptors = []sarama.ConsumerInterceptor{delay}

	// 创建一个消费组
	client, err := sarama.NewConsumerGroup(brokerList[:], "consumer_topic-study", config)
	if err != nil {
		log.Printf("ConsumerInit fail : %v\n", err)
		panic(err)
	}

	// 资源回收
	// defer func() { _ = client.Close() }()

	// 打印异常
	go func() {
		for err := range client.Errors() {
			log.Printf("ConsumerInit Consumer Group Err : %v\n", err)
		}
	}()

	consumer := Consumer{ctx}
	go func() {
		for {
			err := client.Consume(ctx, []string{"topic-study"}, &consumer)
			if err != nil {
				log.Printf("ConsumerInit Consume fail:%v \n", err)
				time.Sleep(time.Second * 5)
			}

			log.Printf("ConsumerInit Consume")
		}
	}()
}

// kafka消费组初始化
func ConsumerInit(ctx context.Context) {
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_0
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewConsumer(brokerList[:], config)
	if err != nil {
		panic(err)
	}

	// 指定分区消费组
	consume, err := client.ConsumePartition("topic-study", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case message := <-consume.Messages():
			log.Println(message, message.Offset)
		case <-ctx.Done():
			return
		}
	}
}
