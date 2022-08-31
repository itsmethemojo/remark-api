package main

import (
	"errors"
	"log"
	"strings"
)

type AuthentificationModel struct {
}

func (this AuthentificationModel) GetUserID(token string) (string, error) {
	loginProvider := (EnvHelper).Get(EnvHelper{}, "LOGIN_PROVIDER")
	switch loginProvider {
	case "DEMO_TOKEN":
		for _, tokenAndID := range strings.Split((EnvHelper).Get(EnvHelper{}, "DEMO_TOKENS"), ",") {
			splittedTokenAndID := strings.Split(tokenAndID, ":")
			if splittedTokenAndID[0] == token {
				return splittedTokenAndID[1], nil
			}
		}
	case "DATABASE_TABLE":
		tokenRepository := TokenRepository{}
		isValid, userID := tokenRepository.tokenIsValid(token)
		log.Printf("[INFO] %v %v", isValid, userID)
		if isValid {
			return userID, nil
		}
	}

	//TODO improve naming
	log.Printf("[INFO] token \"%v\" not authorized", token)
	return "", errors.New("unauthentificated")
}
