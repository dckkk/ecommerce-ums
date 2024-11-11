package main

import (
	"ecommerce-ums/cmd"
	"ecommerce-ums/helpers"
)

func main() {
	// load config
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	// load db
	helpers.SetupPostgreSQL()

	// load redis
	// helpers.SetupRedis()

	// run kafka consumer
	// cmd.ServeKafkaConsumer()

	// run http
	cmd.ServeHTTP()
}
