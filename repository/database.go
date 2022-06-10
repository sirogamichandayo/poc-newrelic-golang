package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/dijsilva/golang-api-newrelic/config"
	"github.com/dijsilva/golang-api-newrelic/entities"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpgx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func ConnectDatabase() {
	databaseHost := config.Configuration.PostgresHost
	user := config.Configuration.PostgresUser
	password := config.Configuration.PostgresPassword
	databaseName := config.Configuration.PostgresDatabse
	port := config.Configuration.PostgresPort

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", databaseHost, user, password, databaseName, port)

	conn, _ := sql.Open("nrpgx", dsn)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{})

	dbSql, err := db.DB()

	dbSql.SetMaxOpenConns(10)
	dbSql.SetMaxIdleConns(10)
	dbSql.SetConnMaxLifetime(time.Second * time.Duration(10))

	if err != nil {
		log.Fatalf("Fail to initialize database %s", err.Error())
	}

	db.AutoMigrate(&entities.User{})

	database = db
}
