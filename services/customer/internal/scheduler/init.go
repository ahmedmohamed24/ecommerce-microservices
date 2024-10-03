package scheduler

import (
	"log"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/types"
	"github.com/robfig/cron/v3"
)

func Schedule() {
	c := cron.New()
	_, err := c.AddJob("0 0 * * *", &types.RefreshToken{})
	if err != nil {
		log.Println(err)
	}
	c.Start()

}
