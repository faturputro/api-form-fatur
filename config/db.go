package config

import (
	"dummy/entities"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func SetupConnection() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	var logLevel logger.LogLevel

	if viper.GetBool("APP_DEBUG") {
		logLevel = logger.Info
	} else {
		logLevel = logger.Silent
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	var dsn string

	if dbPass == "" {
		dsn = fmt.Sprintf("host=%s user=%s dbname=%s port=%s", dbHost, dbUser, dbName, dbPort)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPass, dbName, dbPort)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		AllowGlobalUpdate:      false,
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Succesfully connected to DB.")
	}
	db.AutoMigrate(
		&entities.Profile{},
		&entities.Education{},
		&entities.Employment{},
		&entities.Skill{},
		&entities.WorkExperience{},
	)
	sqlDb, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	return db
}

func CloseDBConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB connection closed.")
	}
	defer dbSQL.Close()
}
