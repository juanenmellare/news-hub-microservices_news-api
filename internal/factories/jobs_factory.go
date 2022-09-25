package factories

import (
	"news-hub-microservices_news-api/configs"
	"news-hub-microservices_news-api/internal/jobs"
)

type JobsFactory interface {
	RunAll(config configs.Config)
}

type jobsFactory struct {
	fetchNewsCronJob jobs.CronJob
}

func NewJobsFactory(layersFactory LayersFactory) *jobsFactory {
	return &jobsFactory{
		fetchNewsCronJob: jobs.NewCronJob("Fetch", "*/30 * * * *", layersFactory.GetNewsService().Fetch),
	}
}

func (j jobsFactory) RunAll() {
	j.fetchNewsCronJob.Run()
}
