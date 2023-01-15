package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
)

type AuthentificationModel struct {
}

func (this AuthentificationModel) GetUserID(token string) (string, error) {
	loginProvider := os.Getenv("LOGIN_PROVIDER")
	switch loginProvider {
	case "DEMO_TOKEN":
		for _, tokenAndID := range strings.Split(os.Getenv("DEMO_TOKENS"), ",") {
			splittedTokenAndID := strings.Split(tokenAndID, ":")
			if splittedTokenAndID[0] == token {
				return splittedTokenAndID[1], nil
			}
		}
	case "DEX":
		provider, err := oidc.NewProvider(context.Background(), os.Getenv("DEX_URI"))
		if err != nil {
			panic(fmt.Sprintf("failed to get token: %v", err))
		}

		idTokenVerifier := provider.Verifier(&oidc.Config{ClientID: os.Getenv("DEX_CLIENT_ID")})
		// token should be Bearer xxxxxxx
		fields := strings.Split(token, " ")
		//check if first part is Bearer
		idToken, err := idTokenVerifier.Verify(context.Background(), fields[1])
		if err != nil {
			return "", fmt.Errorf("could not verify bearer token: %v", err)
		}
		// Extract custom claims.
		var claims struct {
			Email             string `json:"email"`
			Verified          bool   `json:"email_verified"`
			PreferredUsername string `json:"preferred_username"`
		}
		if err := idToken.Claims(&claims); err != nil {
			return "", fmt.Errorf("failed to parse claims: %v", err)
		}
		if !claims.Verified {
			return "", fmt.Errorf("email (%q) in returned claims was not verified", claims.Email)
		}
		return claims.PreferredUsername, nil
	}

	//TODO improve naming
	log.Printf("[INFO] token \"%v\" not authorized", token)
	return "", errors.New("unauthentificated")
}
