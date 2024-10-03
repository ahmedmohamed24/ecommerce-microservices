package scheduler

import (
	"log"

	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/scheduler/tasks"
	"github.com/robfig/cron"
)

func Init() {
	scheduler := cron.New()
	err := scheduler.AddFunc("0 0 0 * * *", tasks.FlushExpiredRefreshTokens)
	if err != nil {
		log.Println(err)
	}
	scheduler.Start()
}
