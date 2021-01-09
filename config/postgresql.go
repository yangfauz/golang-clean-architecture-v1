package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreDatabase(configuration Config) *gorm.DB {

	dsn := "host= " + configuration.Get("PSQL_HOST") + " user= " + configuration.Get("PSQL_USER") + " password= " + configuration.Get("PSQL_PASS") + " dbname= " + configuration.Get("PSQL_DB") + " port= " + configuration.Get("PSQL_PORT") + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
