package repositories

import (
	"github.com/stretchr/testify/assert"
	"news-hub-microservices_news-api/internal/databases"
	"testing"
)

func Test_NewNewsRepository(t *testing.T) {
	var relationalDatabase databases.RelationalDatabase

	assert.Implements(t, (*NewsRepository)(nil), NewNewsRepository(relationalDatabase))
}
