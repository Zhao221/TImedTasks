package task

import (
	"gTest/TaskTest/logger"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"syscall"
)

type TaskManager struct {
	*cron.Cron
	Name string
}

func New(name string, opts ...cron.Option) *TaskManager {
	return &TaskManager{
		Cron: cron.New(opts...),
		Name: name,
	}
}

func (task *TaskManager) Wait() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
	select {
	case sig := <-ch:
		logger.Zap().Infow("task receive signal", "taskName", task.Name, "signal", sig)
		ctx := task.Cron.Stop()
		logger.Zap().Infow("task wait exit", "taskName", task.Name, "signal", sig)
		<-ctx.Done()
		logger.Zap().Infow("task success exit", "taskName", task.Name, "signal", sig)

	}
}
