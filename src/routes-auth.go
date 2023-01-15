package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

func addAuthRoutes(rg *gin.RouterGroup) {
	rg.GET("/start/", routeStartInit)
	rg.POST("/start/", routeStart)
	rg.GET("/callback/", routeCallback)
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func routeStartInit(c *gin.Context) {
	// name, value string, maxAge int, path, domain string, secure, httpOnly bool
	// 10 min
	c.SetCookie("auth_state", randSeq(30), 600, "/", (EnvHelper).Get(EnvHelper{}, "SWAGGER_HOST"), true, true)

	c.HTML(http.StatusOK, "start.tmpl", gin.H{
		"title": "start authentification",
	})
}

func routeStart(c *gin.Context) {
	auth_state, err := c.Cookie("auth_state")
	if err != nil {
		panic("auth_state is missing")
	}

	dex_uri := (EnvHelper).Get(EnvHelper{}, "DEX_URI")
	dex_client_id := (EnvHelper).Get(EnvHelper{}, "DEX_CLIENT_ID")
	dex_redirect_uri := (EnvHelper).Get(EnvHelper{}, "DEX_REDIRECT_URI")
	dex_connector_id := (EnvHelper).Get(EnvHelper{}, "DEX_CONNECTOR_ID")
	login_url := dex_uri + "/auth?client_id=" + url.QueryEscape(dex_client_id) + "&redirect_uri=" + url.QueryEscape(dex_redirect_uri) + "&connector_id=" + url.QueryEscape(dex_connector_id) + "&state=" + url.QueryEscape(auth_state) + "&response_type=code&scope=openid+profile+email"

	c.Redirect(http.StatusFound, login_url)
}

func routeCallback(c *gin.Context) {
	auth_state, state_err := c.Cookie("auth_state")
	if state_err != nil {
		panic("auth_state is missing")
	}

	var (
		err   error
		token *oauth2.Token
	)

	provider, err := oidc.NewProvider(context.Background(), (EnvHelper).Get(EnvHelper{}, "DEX_URI"))
	if err != nil {
		panic(fmt.Sprintf("failed to get token: %v", err))
	}

	oauth2Config := oauth2.Config{
		ClientID:     (EnvHelper).Get(EnvHelper{}, "DEX_CLIENT_ID"),
		ClientSecret: (EnvHelper).Get(EnvHelper{}, "DEX_CLIENT_SECRET"),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{"openid", "profile", "email"},
		RedirectURL:  (EnvHelper).Get(EnvHelper{}, "DEX_REDIRECT_URI"),
	}

	idTokenVerifier := provider.Verifier(&oidc.Config{ClientID: (EnvHelper).Get(EnvHelper{}, "DEX_CLIENT_ID")})

	code := c.DefaultQuery("code", "")
	if code == "" {
		c.String(http.StatusInternalServerError, "no code in request")
		return
	}
	if state := c.DefaultQuery("state", ""); state != auth_state {
		c.String(http.StatusInternalServerError, fmt.Sprintf("expected state %q got %q", auth_state, state))
		return
	}
	token, err = oauth2Config.Exchange(c, code)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get token: %v", err))
		return
	}
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		c.String(http.StatusInternalServerError, "no id_token in token response")
		return
	}

	idToken, err := idTokenVerifier.Verify(c, rawIDToken)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to verify ID token: %v", err))
		return
	}

	//accessToken, ok := token.Extra("access_token").(string)
	//if !ok {
	//	c.String(http.StatusInternalServerError, "no access_token in token response")
	//	return
	//}

	var claims json.RawMessage
	if err := idToken.Claims(&claims); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error decoding ID token claims: %v", err))
		return
	}

	// 30 days
	c.SetCookie("auth_state", "", 0, "/", (EnvHelper).Get(EnvHelper{}, "SWAGGER_HOST"), true, true)
	c.SetCookie("Authorization", "Bearer "+rawIDToken, 2592000, "/", (EnvHelper).Get(EnvHelper{}, "SWAGGER_HOST"), true, true)

	c.String(http.StatusOK, "Bearer "+rawIDToken)
	// redirect to FRONTEND_URI
}
