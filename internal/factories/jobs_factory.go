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
		fetchNewsCronJob: jobs.NewCronJob("FetchNews", "*/30 * * * *", layersFactory.GetNewsService().FetchNews),
	}
}

func (j jobsFactory) RunAll() {
	j.fetchNewsCronJob.Run()
}
