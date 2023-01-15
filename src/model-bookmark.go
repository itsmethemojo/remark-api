package main

import (
	"log"
	"strconv"
)

type BookmarkModel struct {
}

func (this BookmarkModel) DeleteAllData() error {
	bookmarkRepository := BookmarkRepository{}
	repositoryError := bookmarkRepository.DeleteAllData()
	return repositoryError
}

func (this BookmarkModel) ListAll(username string) (AllBookmarkData, error) {
	bookmarkRepository := BookmarkRepository{}
	bookmarkRepositoryData, repositoryError := bookmarkRepository.ListAll(username)
	return bookmarkRepositoryData, repositoryError
}

func (this BookmarkModel) Remark(username string, url string) error {
	bookmarkRepository := BookmarkRepository{}
	repositoryError := bookmarkRepository.Remark(username, url)
	return repositoryError
}

func (this BookmarkModel) Click(username string, id string) error {
	parsedID, parsedIDError := strconv.ParseUint(id, 10, 32)
	if parsedIDError != nil {
		log.Printf("[ERROR] could not convert bookmark id \"%v\" into uint64", username)
		return parsedIDError
	}
	bookmarkRepository := BookmarkRepository{}
	repositoryError := bookmarkRepository.Click(username, parsedID)
	return repositoryError
}

func (this BookmarkModel) Edit(username string, id string, title string) error {
	parsedID, parsedIDError := strconv.ParseUint(id, 10, 32)
	if parsedIDError != nil {
		log.Printf("[ERROR] could not convert bookmark id \"%v\" into uint64", username)
		return parsedIDError
	}
	bookmarkRepository := BookmarkRepository{}
	repositoryError := bookmarkRepository.Edit(username, parsedID, title)
	return repositoryError
}

func (this BookmarkModel) Delete(username string, id string) error {
	parsedID, parsedIDError := strconv.ParseUint(id, 10, 32)
	if parsedIDError != nil {
		log.Printf("[ERROR] could not convert bookmark id \"%v\" into uint64", username)
		return parsedIDError
	}
	bookmarkRepository := BookmarkRepository{}
	repositoryError := bookmarkRepository.Delete(username, parsedID)
	return repositoryError
}
