package main

import (
	"os"
)

type EnvHelper struct {
}

func defaultEnvValues() map[string]string {
	return map[string]string{
		"HOST":                             "localhost",
		"PORT":                             "8080",
		"SCHEMA":                           "http",
		"DATABASE_URL":                     "root:rootpw@tcp(database:3306)/remark?charset=utf8mb4&parseTime=True&loc=Local",
		"DATABASE_CONNECT_RETRY_COUNT":     "6",
		"DATABASE_CONNECT_WAIT_INTERVAL":   "2",
		"DEMO_TOKENS":                      "LOCAL_TEST_TOKEN_1:1,LOCAL_TEST_TOKEN_2:2",
		"LOGIN_PROVIDER":                   "DEMO_TOKEN",
		"CORS_ENABLED":                     "1",
		"ACCESS_CONTROL_ALLOW_ORIGIN":      "*",
		"ACCESS_CONTROL_ALLOW_CREDENTIALS": "true",
		"ACCESS_CONTROL_ALLOW_HEADERS":     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		"ACCESS_CONTROL_ALLOW_METHODS":     "POST,HEAD,PATCH, OPTIONS, GET, PUT",
		"LOGIN_DATABASE_URL":               "root:rootpw@tcp(database:3306)/remark?charset=utf8mb4&parseTime=True&loc=Local",
	}
}

func (this EnvHelper) Get(key string) string {
	envValue, exists := os.LookupEnv(key)
	if exists {
		return envValue
	}
	defaultEnvValue, defaultExists := defaultEnvValues()[key]
	if defaultExists {
		return defaultEnvValue
	}
	return ""
}
