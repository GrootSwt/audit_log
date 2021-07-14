package mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"go_code/audit_log/common"
	"go_code/audit_log/convertor"
	"go_code/audit_log/dto"
	"go_code/audit_log/service"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetMsg() {
	mqModel := common.ConfigInformation.MQConfigInfo
	host := mqModel.Host
	port := mqModel.Port
	user := mqModel.User
	password := mqModel.Password
	queueName := mqModel.QueueName
	//	创建连接
	conn, err := amqp.Dial("amqp://"+user+":"+password+"@"+host+":"+port+"/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	//	创建连接通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	//	如果没有消息队列，创建消息队列
	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	//	从消息队列中获取消息
	msgList, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	//	创建协程获取消息队列
	go func() {
		for d := range msgList {
			//	json数据转换为logDto
			var logDto dto.LogDto
			err = json.Unmarshal(d.Body, &logDto)
			failOnError(err, "Failed to unmarshal json")
			//	logDto转换为logModel
			logModel := convertor.LogDtoTogModel(logDto)
			service.InsertLogService(logModel)
		}
	}()

	<-forever
}
