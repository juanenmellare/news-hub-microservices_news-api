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
	GetNewsProxyApiBaseUrl() string
	GetNewsProxyApiUsername() string
	GetNewsProxyApiPassword() string
	GetUsersApiBaseUrl() string
	GetUsersApiUsername() string
	GetUsersApiPassword() string
	GetBasicAuthUsername() string
	GetBasicAuthPassword() string
}

type config struct {
	port                 string
	databaseHost         string
	databaseName         string
	databasePort         string
	databaseUser         string
	databasePass         string
	newsProxyApiBaseUrl  string
	newsProxyApiUsername string
	newsProxyApiPassword string
	usersApiBaseUrl      string
	usersApiUsername     string
	usersApiPassword     string
	basicAuthUsername    string
	basicAuthPassword    string
}

func NewConfig() Config {
	return &config{
		port:                 getStringValueOrDefault("PORT", "8082"),
		databaseHost:         getStringValueOrDefault("DATABASE_HOST", "localhost"),
		databaseName:         getStringValueOrDefault("DATABASE_NAME", "development.news-hub_news-api"),
		databasePort:         getStringValueOrDefault("DATABASE_PORT", "5432"),
		databaseUser:         getStringValueOrDefault("DATABASE_USER", "admin"),
		databasePass:         getStringValueOrDefault("DATABASE_PASS", ""),
		newsProxyApiBaseUrl:  getStringValueOrDefault("NEWS_PROXY_API_BASE_URL", "https://news-hub-microservices-np-api.herokuapp.com"),
		newsProxyApiUsername: getStringValueOrDefault("NEWS_PROXY_API_USERNAME", "admin"),
		newsProxyApiPassword: getStringValueOrDefault("NEWS_PROXY_API_PASSWORD", "password"),
		usersApiBaseUrl:      getStringValueOrDefault("USERS_API_BASE_URL", "https://news-hub-microservices-u-api.herokuapp.com"),
		usersApiUsername:     getStringValueOrDefault("USERS_API_USERNAME", "admin"),
		usersApiPassword:     getStringValueOrDefault("USERS_API_PASSWORD", "password"),
		basicAuthUsername:    getStringValueOrDefault("BASIC_AUTH_USERNAME", "admin"),
		basicAuthPassword:    getStringValueOrDefault("BASIC_AUTH_PASSWORD", "password"),
	}
}

func getStringValueOrDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}

func (c config) GetPort() string {
	return c.port
}

func (c config) GetDatabaseHost() string {
	return c.databaseHost
}

func (c config) GetDatabaseName() string {
	return c.databaseName
}

func (c config) GetDatabasePort() string {
	return c.databasePort
}

func (c config) GetDatabaseUser() string {
	return c.databaseUser
}

func (c config) GetDatabasePass() string {
	return c.databasePass
}

func (c config) GetNewsProxyApiBaseUrl() string {
	return c.newsProxyApiBaseUrl
}

func (c config) GetNewsProxyApiUsername() string {
	return c.newsProxyApiUsername
}

func (c config) GetNewsProxyApiPassword() string {
	return c.newsProxyApiPassword
}

func (c config) GetUsersApiBaseUrl() string {
	return c.usersApiBaseUrl
}

func (c config) GetUsersApiUsername() string {
	return c.usersApiUsername
}

func (c config) GetUsersApiPassword() string {
	return c.usersApiPassword
}

func (c config) GetBasicAuthUsername() string {
	return c.basicAuthUsername
}

func (c config) GetBasicAuthPassword() string {
	return c.basicAuthPassword
}
