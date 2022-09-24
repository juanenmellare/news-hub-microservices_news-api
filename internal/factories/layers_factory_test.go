package factories

import (
	"github.com/stretchr/testify/assert"
	"news-hub-microservices_news-api/configs"
	"news-hub-microservices_news-api/internal/controllers"
	"news-hub-microservices_news-api/internal/services"
	"testing"
)

func buildDomainLayersFactory() LayersFactory {
	return NewLayersFactory(nil, configs.NewConfig())
}

func TestNewDomainLayersFactory(t *testing.T) {
	domainLayersFactory := buildDomainLayersFactory()

	assert.Implements(t, (*LayersFactory)(nil), domainLayersFactory)
}

func Test_domainLayersFactory_GetHealthChecksController(t *testing.T) {
	domainLayersFactory := buildDomainLayersFactory()

	assert.Implements(t, (*controllers.HealthChecksController)(nil), domainLayersFactory.GetHealthChecksController())
}

func Test_domainLayersFactory_GetNewsService(t *testing.T) {
	domainLayersFactory := buildDomainLayersFactory()

	assert.Implements(t, (*services.NewsService)(nil), domainLayersFactory.GetNewsService())
}

func Test_domainLayersFactory_GetNewsController(t *testing.T) {
	domainLayersFactory := buildDomainLayersFactory()

	assert.Implements(t, (*controllers.NewsController)(nil), domainLayersFactory.GetNewsController())
}
