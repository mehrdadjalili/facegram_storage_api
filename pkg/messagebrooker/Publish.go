package messagebrooker

import (
	"encoding/json"
	"facegram_file_server/pkg/_crypto"
	"github.com/gofrs/uuid"
	"github.com/streadway/amqp"
	"time"
)

type MessagesStruct struct {
	Queue   string
	Message amqp.Publishing
}

func publishToQueue(data MessagesStruct) bool {
	co, err := GetRabbitMqConnection()
	if err != nil {
		return false
	}

	channel, err := co.Channel()
	if err != nil {
		return false
	}

	_, err = channel.QueueDeclare(data.Queue, true, false, false, true, nil)
	if err != nil {
		return false
	}

	err = channel.Publish("", data.Queue, false, false, data.Message)
	if err != nil {
		return false
	}

	return true
}

func messageBuilder(data string) (string, *amqp.Publishing, error) {

	enc, err := _crypto.AesEncryptInternalData(data)

	if err != nil {
		return "", nil, err
	}

	id, err := uuid.NewGen().NewV4()

	if err != nil {
		return "", nil, err
	}

	message := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Timestamp:    time.Now(),
		MessageId:    id.String(),
		Body:         []byte(enc),
	}

	return id.String(), &message, nil
}

func Publish(queue string, p *interface{}) bool {

	js, e2 := json.Marshal(p)
	if e2 != nil {
		return false
	}

	_, message, e3 := messageBuilder(string(js))
	if e3 != nil {
		return false
	}

	return publishToQueue(MessagesStruct{Queue: queue, Message: *message})
}
