package initialize

import (
	"gin_skeleton/task"
	"github.com/robfig/cron/v3"
)

func InitTask() {

	// 支持秒定时器
	c := cron.New(cron.WithSeconds())
	// 注册定时器
	task.RegisterCron(c)
	// 启动定时器
	c.Start()

}
