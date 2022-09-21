package configs

import (
	"os"
)

type Config interface {
	GetPort() string
	GetDatabaseHost() string
	GetDatabaseName() string
	GetDatabasePort() string
	GetDatabaseUser() string
	GetDatabasePass() string
	GetNewProxyApiBaseUrl() string
}

type configImpl struct {
	port                string
	databaseHost        string
	databaseName        string
	databasePort        string
	databaseUser        string
	databasePass        string
	newsProxyApiBaseUrl string
}

func NewConfig() Config {
	return &configImpl{
		port:                getStringValueOrDefault("PORT", "8082"),
		databaseHost:        getStringValueOrDefault("DATABASE_HOST", "localhost"),
		databaseName:        getStringValueOrDefault("DATABASE_NAME", "development.news-hub_news-api"),
		databasePort:        getStringValueOrDefault("DATABASE_PORT", "5432"),
		databaseUser:        getStringValueOrDefault("DATABASE_USER", "admin"),
		databasePass:        getStringValueOrDefault("DATABASE_PASS", "news-hub.2022"),
		newsProxyApiBaseUrl: getStringValueOrDefault("NEWS_PROXY_API_BASE_URL", ""),
	}
}

func getStringValueOrDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}

func (c configImpl) GetPort() string {
	return c.port
}

func (c configImpl) GetDatabaseHost() string {
	return c.databaseHost
}

func (c configImpl) GetDatabaseName() string {
	return c.databaseName
}

func (c configImpl) GetDatabasePort() string {
	return c.databasePort
}

func (c configImpl) GetDatabaseUser() string {
	return c.databaseUser
}

func (c configImpl) GetDatabasePass() string {
	return c.databasePass
}

func (c configImpl) GetNewProxyApiBaseUrl() string {
	return c.newsProxyApiBaseUrl
}
