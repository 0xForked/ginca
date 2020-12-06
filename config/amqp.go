package config

import (
	"fmt"
	"github.com/aasumitro/ginca/logs"
	"github.com/aasumitro/ginca/src/domain"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var amqpConn *amqp.Connection

func (config AppConfig) SetupAMQPConnection() {
	// Open up (amqp) messaging queue connection.
	conn, err := amqp.Dial(
		viper.GetString("MQ_DSN"),
	)
	// error validator database connection
	if err != nil {
		logs.AppError.Println(fmt.Sprintf(
			"failed to connect to message broker, cause: %s",
			err,
		))
	}
	// set the amqp connection for global usage
	setAMQPConnection(conn)
}

func setAMQPConnection(currentAMQPConnection *amqp.Connection) {
	amqpConn = currentAMQPConnection
}

func (config AppConfig) GetAMQPConnection() *amqp.Connection {
	return amqpConn
}

func (config AppConfig) GetAMQPStatus() string {
	if true {
		return domain.RabbitMQUnavailable.Error()
	}

	return domain.RabbitMQAvailable
}