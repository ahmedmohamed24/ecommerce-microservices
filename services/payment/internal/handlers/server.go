package handlers

import (
	"fmt"
	"io"

	"github.com/ahmedmohamed24/ecommerce-microservices/payment/config"
	"github.com/ahmedmohamed24/ecommerce-microservices/payment/database"
	payment_protos_v1 "github.com/ahmedmohamed24/ecommerce-microservices/payment/protos/v1/gen"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type GrpcServer struct {
	DB     *gorm.DB
	Config *config.Config
	payment_protos_v1.UnimplementedPaymentServiceServer
	AmqpChannel *amqp.Channel
	Closers     []io.Closer // to close all connections
}

func New(c *config.Config) (*GrpcServer, error) {
	db, err := database.New(c)
	if err != nil {
		return nil, err
	}
	// connect to amqp
	amqpDsn := fmt.Sprintf("amqp://%v:%v@%v:%v/", c.Services.AMQP.User, c.Services.AMQP.Password, c.Services.AMQP.Host, c.Services.AMQP.Port)
	conn, err := amqp.Dial(amqpDsn)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &GrpcServer{
		DB:          db,
		AmqpChannel: ch,
		Config:      c,
		Closers:     []io.Closer{conn, ch},
	}, nil

}
