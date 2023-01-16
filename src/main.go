package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	err1 := godotenv.Load("default.env")
	if err1 != nil {
		panic("default.env loading failed")
	}
	err2 := godotenv.Overload()
	if err2 != nil {
		log.Printf("[INFO] no .env file present, skip loading")
	}

	//TODO maybe this can be removed
	if os.Getenv("TEST_MODE") == "true" {
		for _, env := range os.Environ() {
			log.Printf("[DEBUG] \"%v\" into uint64", env)
		}
	}

	MigrateDatabase()
	RoutesRun()
}

//TODO not sure if thats the best place for this
func MigrateDatabase() {
	// in case the database needs time to start. lets wait and try again a couple of times
	retryCount, retryCountParseErr := strconv.ParseUint(os.Getenv("DATABASE_CONNECT_RETRY_COUNT"), 10, 32)
	if retryCountParseErr != nil {
		log.Printf("[ERROR] could not convert userID \"%v\" into uint64", os.Getenv("DATABASE_CONNECT_RETRY_COUNT"))
		panic("DATABASE_CONNECT_RETRY_COUNT is not an interger")
	}
	waitInterval, waitIntervalParseErr := strconv.ParseUint(os.Getenv("DATABASE_CONNECT_WAIT_INTERVAL"), 10, 32)
	if waitIntervalParseErr != nil {
		log.Printf("[ERROR] could not convert userID \"%v\" into uint64", os.Getenv("DATABASE_CONNECT_WAIT_INTERVAL"))
		panic("DATABASE_CONNECT_WAIT_INTERVAL is not an integer")
	}
	var connectError error
	var mDatabase *gorm.DB
	for i := 0; i < int(retryCount); i++ {
		dsn := "host=" + os.Getenv("DATABASE_HOST") +
			" user=" + os.Getenv("DATABASE_USERNAME") +
			" password=" + os.Getenv("DATABASE_PASSWORD") +
			" dbname=" + os.Getenv("DATABASE_NAME") +
			" port=" + os.Getenv("DATABASE_PORT") +
			" sslmode=" + os.Getenv("DATABASE_SSLMODE") +
			" TimeZone=" + os.Getenv("DATABASE_TIMEZONE")

		mDatabase, connectError = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if connectError == nil {
			break
		}
		log.Printf("[INFO] waiting another %v seconds for database to come up", waitInterval)
		time.Sleep(time.Duration(waitInterval) * time.Second)
	}
	bookmarkRepository := BookmarkRepository{mDatabase}
	initError := bookmarkRepository.InitializeDatabase()
	if initError != nil {
		panic("could not connect to database")
	}
}
