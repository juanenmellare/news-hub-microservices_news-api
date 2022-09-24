package jobs

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

type CronJob interface {
	Run()
	GetName() string
	GetCron() string
}

type cronJob struct {
	name     string
	cron     string
	function func()
}

func (j cronJob) Run() {
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Cron(j.cron).StartImmediately().Do(func() {
		fmt.Println(fmt.Sprintf("%s - Started at %s", j.name, time.Now().UTC().String()))
		j.function()
		fmt.Println(fmt.Sprintf("%s - Finished at %s", j.name, time.Now().UTC().String()))
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("Job Error while formatting the cron: %s", j.name))
	}
	scheduler.StartAsync()
}

func (j cronJob) GetName() string {
	return j.name
}

func (j cronJob) GetCron() string {
	return j.cron
}

func NewCronJob(name, cron string, function func()) CronJob {
	return &cronJob{
		name,
		cron,
		function,
	}
}
