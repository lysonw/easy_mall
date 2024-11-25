package script

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func HotProductCron() {
	go func() {
		c := cron.New()
		// 添加任务
		_, err := c.AddFunc("@every 1s", func() { fmt.Println("Every hour, starting an hour from now") })
		if err != nil {

		}
		c.Start()

		select {}
	}()
}

func HotProductHandle() {

}
