package main

import (
	"errors"
	"log"
	"os"
	"strings"
)

type AuthentificationModel struct {
}

func (this AuthentificationModel) GetUserID(token string) (string, error) {

	//TODO add switch here (os.Getenv("LOGIN_PROVIDER")
	for _, tokenAndID := range strings.Split(os.Getenv("TOKENS"), ",") {
		tokenAndIDtokenAndID := strings.Split(tokenAndID, ":")
		//log.Println(tokenAndIDtokenAndID[0] + " " + token)
		if tokenAndIDtokenAndID[0] == token {
			return tokenAndIDtokenAndID[1], nil
		}
	}
	//TODO improve naming
	log.Printf("[INFO] token \"%v\" not authorized", token)
	return "", errors.New("unauthentificated")
}
