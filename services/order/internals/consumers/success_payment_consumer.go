package consumers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ahmedmohamed24/ecommerce-microservices/order/internals/handlers"
	"github.com/ahmedmohamed24/ecommerce-microservices/order/internals/types"
	order_protos_v1 "github.com/ahmedmohamed24/ecommerce-microservices/order/protos/v1/gen"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SuccessPaymentConsumer(h *handlers.Server) error {
	// start the consumer of the RabbitMQ
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v:%v/", h.Config.Services.AMQP.User, h.Config.Services.AMQP.Password, h.Config.Services.AMQP.Host, h.Config.Services.AMQP.Port))
	if err != nil {
		return err
	}
	h.Closers = append(h.Closers, conn)

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	h.Closers = append(h.Closers, ch)

	q, err := ch.QueueDeclare("success_payment", false, false, false, false, nil)
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for d := range msgs {
			log.Println("message received from the consumer")
			msg := &order_protos_v1.PaymentDoneSuccessfully{}
			err := json.Unmarshal(d.Body, msg)
			if err != nil {
				log.Println(err)
				continue
			}
			order := &types.Order{}
			if err := h.DB.First(order, msg.OrderId).Error; err != nil {
				log.Println(err)
				continue
			}
			order.StatusID = types.OrderStatusCompleted
			if err := h.DB.Save(order).Error; err != nil {
				log.Println(err)
				continue
			}
			fmt.Println("Order status updated successfully")
		}
	}()
	return nil

}
