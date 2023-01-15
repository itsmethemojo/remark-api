package main

import (
	"os"
)

type EnvHelper struct {
}

func defaultEnvValues() map[string]string {
	return map[string]string{
		"ACCESS_CONTROL_ALLOW_CREDENTIALS": "true",
		"ACCESS_CONTROL_ALLOW_HEADERS":     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		"ACCESS_CONTROL_ALLOW_METHODS":     "POST,HEAD,PATCH, OPTIONS, GET, PUT",
		"ACCESS_CONTROL_ALLOW_ORIGIN":      "*",
		"API_PATH_PREFIX":                  "",
		"APP_DOMAIN":                       "localhost",
		"APP_PORT":                         "8080",
		"APP_SCHEMA":                       "http",
		"CORS_ENABLED":                     "1",
		"DATABASE_CONNECT_RETRY_COUNT":     "10",
		"DATABASE_CONNECT_WAIT_INTERVAL":   "5",
		"DATABASE_HOST":                    "database",
		"DATABASE_NAME":                    "remark",
		"DATABASE_PASSWORD":                "remarkpassword",
		"DATABASE_PORT":                    "5432",
		"DATABASE_SSLMODE":                 "disable",
		"DATABASE_TIMEZONE":                "UTC",
		"DATABASE_USERNAME":                "postgres",
		"DEMO_TOKENS":                      "LOCAL_TEST_TOKEN_1:1,LOCAL_TEST_TOKEN_2:2",
		"LOGIN_DATABASE_URL":               "root:rootpw@tcp(database:3306)/remark?charset=utf8mb4&parseTime=True&loc=Local",
		"LOGIN_PROVIDER":                   "DEX",
		"SWAGGER_PATH":                     "/swagger",
		"TEST_MODE":                        "false",
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
