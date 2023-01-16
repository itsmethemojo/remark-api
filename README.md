# reMARK

api to store bookmarks and sort/rank them over time

the frontend can be found [here](https://github.com/itsmethemojo/remark)

the initial version was written in php [see](https://github.com/itsmethemojo/remark-api/tree/21f1ad861827053c76f328d39174b97c225cc5d4)

## prequisites

- [go-task](https://github.com/go-task/task)
- [docker](https://www.docker.com/get-started)

## installation

`task dz:init`

## start server local

`docker-compose up -d --build`

## see swagger api documentation

[open](http://localhost:8080/swagger/index.html)

## available environment parameters

| Name | Default | Description |
|------|---------|-------------|
| ACCESS_CONTROL_ALLOW_CREDENTIALS | `true` | CORS Header |
| ACCESS_CONTROL_ALLOW_HEADERS | `Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With` | CORS Header |
| ACCESS_CONTROL_ALLOW_METHODS | `POST,HEAD,PATCH, OPTIONS, GET, PUT` | CORS Header |
| ACCESS_CONTROL_ALLOW_ORIGIN | `true` | CORS Header |
| API_PATH_PREFIX | `""` | path prefix to host the api on a path like `/you/custom/path/v1/bookmarks` |
| APP_DOMAIN | `localhost` | web domain |
| APP_PORT | `8080` | port of the webserver |
| APP_SCHEMA | `http` | might be https or http |
| CORS_ENABLED | `1` | toggle CORS Headers |
| DATABASE_CONNECT_RETRY_COUNT | `10` | for local testing if database boots slow, the app will retry later X times |
| DATABASE_CONNECT_WAIT_INTERVAL | `5` | intervall between retries |
| DATABASE_HOST | `database` | - |
| DATABASE_NAME | `remark` | - |
| DATABASE_PASSWORD | `remarkpassword` | - |
| DATABASE_PORT | `5432` | - |
| DATABASE_SSLMODE | `disable` | - |
| DATABASE_TIMEZONE | `UTC` | - |
| DATABASE_USERNAME | `postgres` | - |
| DEMO_TOKENS | `LOCAL_TEST_TOKEN_1:user1,LOCAL_TEST_TOKEN_2:user2` | for local testing this tokens can be used as Authorization Header |
| DEX_CLIENT_ID | - | [see](https://dexidp.io/docs/connectors/github/#configuration) |
| DEX_CLIENT_SECRET | - | [see](https://dexidp.io/docs/connectors/github/#configuration) |
| DEX_CONNECTOR_ID | - | [see](https://dexidp.io/docs/connectors/github/#configuration) |
| DEX_URI | - | URI of the dex instance |
| LOGIN_PROVIDER | `DEMO_TOKEN` | can be `DEX` or `DEMO_TOKEN` |
| SWAGGER_PATH | `/swagger` | path where the swagger ui will be available |
| TEST_MODE | `false` | if `true` an additional api route `DELETE /bookmark/` to clear the database is available, also env parameters will be dumped in startup log |

## lint code

`task lint`

## run tests

```
docker-compose -f docker-compose.yml -f docker-compose-testing.yml up -d --build
./test.sh
```

## TODO

1. improve test suite
2. add github actions for tests
3. check conneciton pooling https://gorm.io/docs/generic_interface.html#Connection-Pool
4. improve naming of AllBookmarkData
5. add debug logs
6. add dev/prod mode

## extract available env parameters

```
 cat src/*.go | tr ' ' '\n' | grep 'os.G' | cut -d '"' -f 2 | awk '{print $1"="}' |sort -u > /tmp/all-used-env.ini
 awk -F= 'NR==FNR{a[$1]=$0;next}($1 in a){$0=a[$1]}1' default.env /tmp/all-used-env.ini
```
