package main

import (
	"TimedTasks/handleFunc"
	"TimedTasks/task"
	"github.com/robfig/cron/v3"
)

func Start() {
	cronTask := task.New("auto_check_compare",
		cron.WithChain(cron.Recover(cron.DefaultLogger),
			cron.DelayIfStillRunning(cron.DefaultLogger),
		),
		cron.WithSeconds(), // 秒级别

	)
	// 每五分钟处理一次
	// 获取未处理的图片
	cronTask.AddJob("*/5 * * * * *", handleFunc.GenCompareToken())
	// 处理一个对比组中只有一个checkKey的第一张图片
	cronTask.AddJob("*/5 * * * * *", handleFunc.GenFirstCompareToken())
	// 将当前图片与上一张图片的信息存储到结构中
	cronTask.AddJob("*/5 * * * * *", handleFunc.GenProcessJob())
	// 进行比较判断图片是否作弊
	cronTask.AddJob("*/5 * * * * *", handleFunc.ComProcessJob())
	cronTask.Start()
	cronTask.Wait()
}

func main() {
	Start()
}
