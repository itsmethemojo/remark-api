package main

import (
	"log"
	"strconv"
)

type BookmarkModel struct {
}

func (this BookmarkModel) ListAll(userID string) (AllBookmarkData, error) {
	parsedUserId, parseErr := strconv.ParseUint(userID, 10, 32)
	if parseErr != nil {
		log.Printf("[ERROR] could not convert userID \"%v\" into uint64", userID)
		emptyData := AllBookmarkData{}
		return emptyData, parseErr
	}
	bookmarkRepository := BookmarkRepository{}
	bookmarkRepositoryData, repositoryError := bookmarkRepository.ListAll(parsedUserId)
	return bookmarkRepositoryData, repositoryError
}

func (this BookmarkModel) Remark(userID string, url string) error {
	log.Printf("[INFO] url \"%v\"", url) //TODO remove
	parsedUserId, parseErr := strconv.ParseUint(userID, 10, 32)
	if parseErr != nil {
		log.Printf("[ERROR] could not convert userID \"%v\" into uint64", userID)
		return parseErr
	}
	bookmarkRepository := BookmarkRepository{}
	repositoryError := bookmarkRepository.Remark(parsedUserId, url)
	return repositoryError
}

func (this BookmarkModel) Click(userID string, id string) error {
	parsedUserID, parsedUserIdError := strconv.ParseUint(userID, 10, 32)
	parsedID, parsedIDError := strconv.ParseUint(id, 10, 32)
	if parsedUserIdError != nil {
		log.Printf("[ERROR] could not convert userID \"%v\" into uint64", userID)
		return parsedUserIdError
	}
	if parsedIDError != nil {
		log.Printf("[ERROR] could not convert bookmark id \"%v\" into uint64", userID)
		return parsedIDError
	}
	bookmarkRepository := BookmarkRepository{}
	repositoryError := bookmarkRepository.Click(parsedUserID, parsedID)
	return repositoryError
}

func (this BookmarkModel) Edit(userID string, id string, title string) error {
	parsedUserID, parsedUserIdError := strconv.ParseUint(userID, 10, 32)
	parsedID, parsedIDError := strconv.ParseUint(id, 10, 32)
	if parsedUserIdError != nil {
		log.Printf("[ERROR] could not convert userID \"%v\" into uint64", userID)
		return parsedUserIdError
	}
	if parsedIDError != nil {
		log.Printf("[ERROR] could not convert bookmark id \"%v\" into uint64", userID)
		return parsedIDError
	}
	bookmarkRepository := BookmarkRepository{}
	repositoryError := bookmarkRepository.Edit(parsedUserID, parsedID, title)
	return repositoryError
}

func (this BookmarkModel) Delete(userID string, id string) error {
	parsedUserID, parsedUserIdError := strconv.ParseUint(userID, 10, 32)
	parsedID, parsedIDError := strconv.ParseUint(id, 10, 32)
	if parsedUserIdError != nil {
		log.Printf("[ERROR] could not convert userID \"%v\" into uint64", userID)
		return parsedUserIdError
	}
	if parsedIDError != nil {
		log.Printf("[ERROR] could not convert bookmark id \"%v\" into uint64", userID)
		return parsedIDError
	}
	bookmarkRepository := BookmarkRepository{}
	repositoryError := bookmarkRepository.Delete(parsedUserID, parsedID)
	return repositoryError
}
