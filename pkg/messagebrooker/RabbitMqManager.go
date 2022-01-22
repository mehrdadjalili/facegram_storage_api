package messagebrooker

import (
	"facegram_file_server/config"
	"github.com/streadway/amqp"
)

func GetRabbitMqConnection() (*amqp.Connection, error) {

	cfg := config.GetRabbitMqtConfig()
	u := cfg.Username
	p := cfg.Password
	a := cfg.Address + cfg.Port

	url := "amqp://" + u + ":" + p + "@" + a + "/"

	co, err := amqp.Dial(url)

	if err != nil {
		return nil, err
	}

	return co, nil
}
