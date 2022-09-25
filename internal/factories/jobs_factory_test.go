package factories

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"news-hub-microservices_news-api/internal/jobs"
	mocks "news-hub-microservices_news-api/test/mocks/factories"
	mocks2 "news-hub-microservices_news-api/test/mocks/services"
	"testing"
)

func TestNewJobsFactory(t *testing.T) {
	newsServiceMock := mocks2.NewsService{}
	newsServiceMock.On("Fetch").Return(nil)

	layersFactoryMock := &mocks.LayersFactory{}
	layersFactoryMock.On("GetNewsService").Return(&newsServiceMock)

	jobsFactory := NewJobsFactory(layersFactoryMock)

	fetchNewsCronJobExpected := jobs.NewCronJob("Fetch", "*/30 * * * *", layersFactoryMock.GetNewsService().Fetch)

	assert.Equal(t, jobsFactory.fetchNewsCronJob.GetName(), fetchNewsCronJobExpected.GetName())
	assert.Equal(t, jobsFactory.fetchNewsCronJob.GetCron(), fetchNewsCronJobExpected.GetCron())
}

func Test_jobsFactory_RunAll(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprintf("the test should not panic: %v", r)
			t.Errorf(str)
		}
	}()

	newsServiceMock := mocks2.NewsService{}
	newsServiceMock.On("Fetch").Return(nil)

	layersFactoryMock := &mocks.LayersFactory{}
	layersFactoryMock.On("GetNewsService").Return(&newsServiceMock)

	jobsFactory := NewJobsFactory(layersFactoryMock)

	jobsFactory.RunAll()
}
