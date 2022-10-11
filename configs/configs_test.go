package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Config_GetPort(t *testing.T) {
	expectedValue := "0000"
	_ = os.Setenv("PORT", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetPort())
}

func Test_Config_GetPort_default(t *testing.T) {
	_ = os.Unsetenv("PORT")

	config := NewConfig()

	assert.Equal(t, "8082", config.GetPort())
}

func Test_Config_GetDatabasePort(t *testing.T) {
	expectedValue := "5431"
	_ = os.Setenv("DATABASE_PORT", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetDatabasePort())
}

func Test_Config_GetDatabasePort_default(t *testing.T) {
	_ = os.Unsetenv("DATABASE_PORT")

	config := NewConfig()

	assert.Equal(t, "5432", config.GetDatabasePort())
}

func TestConfig_GetDatabaseHost(t *testing.T) {
	expectedValue := "foo-host"
	_ = os.Setenv("DATABASE_HOST", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetDatabaseHost())
}

func TestConfig_GetDatabaseHost_default(t *testing.T) {
	_ = os.Unsetenv("DATABASE_HOST")

	config := NewConfig()

	assert.Equal(t, "localhost", config.GetDatabaseHost())
}

func TestConfig_GetDatabaseName(t *testing.T) {
	expectedValue := "foo-name"
	_ = os.Setenv("DATABASE_NAME", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetDatabaseName())
}

func TestConfig_GetDatabaseName_default(t *testing.T) {
	_ = os.Unsetenv("DATABASE_NAME")

	config := NewConfig()

	assert.Equal(t, "development.news-hub_news-api", config.GetDatabaseName())
}

func TestConfig_GetDatabaseUser(t *testing.T) {
	expectedValue := "foo-user"
	_ = os.Setenv("DATABASE_USER", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetDatabaseUser())
}

func TestConfig_GetDatabaseUser_default(t *testing.T) {
	_ = os.Unsetenv("DATABASE_USER")

	config := NewConfig()

	assert.Equal(t, "admin", config.GetDatabaseUser())
}

func TestConfig_GetDatabasePass(t *testing.T) {
	expectedValue := "foo-pass"
	_ = os.Setenv("DATABASE_PASS", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetDatabasePass())
}

func TestConfig_GetDatabasePass_default(t *testing.T) {
	_ = os.Unsetenv("DATABASE_PASS")

	config := NewConfig()

	assert.Equal(t, "", config.GetDatabasePass())
}

func TestConfig_GetNewsProxyApiBaseUrl(t *testing.T) {
	expectedValue := "news-api"
	_ = os.Setenv("NEWS_PROXY_API_BASE_URL", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetNewsProxyApiBaseUrl())
}

func TestConfig_GetNewsProxyApiBaseUrl_default(t *testing.T) {
	_ = os.Unsetenv("NEWS_PROXY_API_BASE_URL")

	config := NewConfig()

	assert.Equal(t, "https://news-hub-microservices-np-api.herokuapp.com", config.GetNewsProxyApiBaseUrl())
}

func TestConfig_GetNewsProxyApiUsername(t *testing.T) {
	expectedValue := "root"
	_ = os.Setenv("NEWS_PROXY_API_USERNAME", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetNewsProxyApiUsername())
}

func TestConfig_GetNewsProxyApiUsername_default(t *testing.T) {
	_ = os.Unsetenv("NEWS_PROXY_API_USERNAME")

	config := NewConfig()

	assert.Equal(t, "admin", config.GetNewsProxyApiUsername())
}

func TestConfig_GetNewsProxyApiPassword(t *testing.T) {
	expectedValue := "password-2"
	_ = os.Setenv("NEWS_PROXY_API_PASSWORD", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetNewsProxyApiPassword())
}

func TestConfig_GetNewsProxyApiPassword_default(t *testing.T) {
	_ = os.Unsetenv("NEWS_PROXY_API_PASSWORD")

	config := NewConfig()

	assert.Equal(t, "password", config.GetNewsProxyApiPassword())
}

func TestConfig_GetUsersApiBaseUrl(t *testing.T) {
	expectedValue := "users-api"
	_ = os.Setenv("USERS_API_BASE_URL", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetUsersApiBaseUrl())
}

func TestConfig_GetUsersApiBaseUrl_default(t *testing.T) {
	_ = os.Unsetenv("USERS_API_BASE_URL")

	config := NewConfig()

	assert.Equal(t, "https://news-hub-microservices-u-api.herokuapp.com", config.GetUsersApiBaseUrl())
}

func TestConfig_GetUsersApiUsername(t *testing.T) {
	expectedValue := "root"
	_ = os.Setenv("USERS_API_USERNAME", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetUsersApiUsername())
}

func TestConfig_GetUsersApiUsername_default(t *testing.T) {
	_ = os.Unsetenv("USERS_API_USERNAME")

	config := NewConfig()

	assert.Equal(t, "admin", config.GetUsersApiUsername())
}

func TestConfig_GetUsersApiPassword(t *testing.T) {
	expectedValue := "password-2"
	_ = os.Setenv("USERS_API_PASSWORD", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetUsersApiPassword())
}

func TestConfig_GetUsersApiPassword_default(t *testing.T) {
	_ = os.Unsetenv("USERS_API_PASSWORD")

	config := NewConfig()

	assert.Equal(t, "password", config.GetUsersApiPassword())
}

func Test_config_GetBasicAuthUsername(t *testing.T) {
	_ = os.Setenv("BASIC_AUTH_USERNAME", "root")

	config := NewConfig()

	assert.Equal(t, "root", config.GetBasicAuthUsername())
}

func Test_config_GetBasicAuthUsername_default(t *testing.T) {
	_ = os.Unsetenv("BASIC_AUTH_USERNAME")

	config := NewConfig()

	assert.Equal(t, "admin", config.GetBasicAuthUsername())
}

func Test_config_GetBasicAuthPassword(t *testing.T) {
	_ = os.Setenv("BASIC_AUTH_PASSWORD", "foo")

	config := NewConfig()

	assert.Equal(t, "foo", config.GetBasicAuthPassword())
}

func Test_config_GetBasicAuthPassword_default(t *testing.T) {
	_ = os.Unsetenv("BASIC_AUTH_PASSWORD")

	config := NewConfig()

	assert.Equal(t, "password", config.GetBasicAuthPassword())
}
