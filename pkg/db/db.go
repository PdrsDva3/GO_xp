package db

import (
	"GO_xp/pkg/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func ConnectDB() *sqlx.DB {
	conn := fmt.Sprintf(
		"user=%v password=%v host=%v port=%v dbname=%v sslmode=disable",
		viper.GetString(config.DBUser),
		viper.GetString(config.DBPassword),
		viper.GetString(config.DBHost),
		viper.GetInt(config.DBPort),
		viper.GetString(config.DBName),
	)
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		panic(fmt.Sprintf("Error while connecting to database: %v", err.Error()))
	}
	return db
}
