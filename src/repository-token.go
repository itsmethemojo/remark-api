package main

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

type Tokens struct {
	Token   string
	UserID  string
	Expires uint64
}

type TokenRepository struct {
}

func (this TokenRepository) getDB(databaseUrl string) (*gorm.DB, error) {
	db, connectError := gorm.Open(mysql.Open((EnvHelper).Get(EnvHelper{}, databaseUrl)), &gorm.Config{})
	if connectError != nil {
		return db, errors.New("could not connect to database")
	}
	return db, nil
}

func (this TokenRepository) tokenIsValid(token string) (bool, string) {
	loginDB, loginDBConnectError := this.getDB("LOGIN_DATABASE_URL")
	if loginDBConnectError != nil {
		log.Printf("[INFO] %v", loginDBConnectError)
		return false, ""
	}
	var tokenEntity Tokens
	initialSearchResult := loginDB.First(&tokenEntity, "token = ?", token)
	if errors.Is(initialSearchResult.Error, gorm.ErrRecordNotFound) {
		return false, ""
	}
	expiresTime := time.Unix(int64(tokenEntity.Expires), 0)
	today := time.Now()
	if today.After(expiresTime) {
		return false, ""
	}
	remarkDB, remarkDBConnectError := this.getDB("DATABASE_URL")
	if remarkDBConnectError != nil {
		log.Printf("[INFO] %v", remarkDBConnectError)
		return false, ""
	}
	var userEntity UserEntity
	userSearchResult := remarkDB.First(&userEntity, "name = ?", tokenEntity.UserID)
	if errors.Is(userSearchResult.Error, gorm.ErrRecordNotFound) {
		log.Printf("[INFO] user \"%v\" not found in user_entity table", tokenEntity.UserID)
		return false, ""
	}
	return true, strconv.Itoa(int(userEntity.ID))
}
