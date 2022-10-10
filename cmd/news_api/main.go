package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"news-hub-microservices_news-api/api"
	"news-hub-microservices_news-api/configs"
	"news-hub-microservices_news-api/internal/databases"
	"news-hub-microservices_news-api/internal/factories"
)

func main() {
	logger := log.Default()
	config := configs.NewConfig()

	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		config.GetDatabaseUser(), config.GetDatabasePass(), config.GetDatabaseHost(),
		config.GetDatabasePort(), config.GetDatabaseName())

	relationalDatabase := databases.NewConnection(gorm.Open(postgres.Open(connectionString), &gorm.Config{}))
	relationalDatabase.DoMigration()

	domainLayersFactory := factories.NewLayersFactory(relationalDatabase, config)

	factories.NewJobsFactory(domainLayersFactory).RunAll()

	port := ":" + config.GetPort()
	if err := api.NewRouter(domainLayersFactory, config).Run(port); err != nil {
		logger.Fatalf("Error while trying to create the router: " + err.Error())
	}
}
