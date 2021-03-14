# reMARK

this is a refactor try with go framwork gin-gonic of the formar php applciation

## prequisites

[go-task](https://github.com/go-task/task)
[docker](https://www.docker.com/get-started)

## installation

`task dz:init`

## start server local

`./local-server.sh`

## see swagger api documentation

[open](http://localhost:8080/swagger/index.html)

## interesting task targets

`task lint`

`task build`

## TODO

1 return error class that includes the api response code
2 improve naming of AllBookmarkData
3 add debug logs
4 use docker compose to start up local database and init it
5 include login-with-twitter with .env feature toggle
5a describe token via swagger https://swagger.io/docs/specification/authentication/cookie-authentication/ https://swagger.io/docs/specification/authentication/api-keys/
https://github.com/swaggo/swag/issues/519#issuecomment-742368554
5b add another token based authentification provider in .env.default to test authentification without twitter
5c elimintate userid from calls if there is always one auth provider
```
TOKENS[]=bladsfdsdsdfsdf:<userid1>
TOKENS[]=blubbbsdsadasda:<userid2>
```
