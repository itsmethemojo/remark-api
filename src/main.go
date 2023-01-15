package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	err1 := godotenv.Load("default.env")
	err2 := godotenv.Load()
	if err1 != nil || err2 != nil{
		panic("dot env loading failed")
	}
	//TODO init connection and database migration might be done lazy calling the first route
	MigrateDatabase()
	RoutesRun()
}

//TODO not sure if thats the best place for this
func MigrateDatabase() {
	bookmarkRepository := BookmarkRepository{}
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
	for i := 0; i < int(retryCount); i++ {
		error := bookmarkRepository.InitializeDatabase()
		if error == nil {
			break
		}
		log.Printf("[INFO] waiting another %v seconds for database to come up", waitInterval)
		time.Sleep(time.Duration(waitInterval) * time.Second)
	}
}