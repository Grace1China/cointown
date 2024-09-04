package initialize

import (
	"fmt"

	"github.com/Grace1China/cointown/server/task"

	"github.com/robfig/cron/v3"

	"github.com/Grace1China/cointown/server/global"
)

func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 清理DB定时任务
		_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := task.ClearTable(global.GVA_DB) // 定时任务方法定在task文件包中
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "定时清理数据库【日志，黑名单】内容", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 其他定时任务定在这里 参考上方使用方法

		_, err1 := global.GVA_Timer.AddTaskByFunc("get Price", "0 */2 * * * *", func() {
			fmt.Println("go 1 min call GetPrice()")
			err := task.GetPrice() // 定时任务方法定在task文件包中
			if err != nil {
				fmt.Println("timer error:", err)
			}

		}, "go 1 min", option...)
		if err1 != nil {
			fmt.Println("add timer error:", err1)
		}
	}()
}
