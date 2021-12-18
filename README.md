# reMARK

this is a refactor try with go framwork gin-gonic of the formar php applciation

## prequisites

[go-task](https://github.com/go-task/task)
[docker](https://www.docker.com/get-started)

## installation

`task dz:init`

## start server local

`task build && docker-compose up -d --build`

## see swagger api documentation

[open](http://localhost:8080/swagger/index.html)

## interesting task targets

`task lint`

`task build`

## TODO

1 return error class that includes the api response code
2 improve naming of AllBookmarkData
3 add debug logs
4 include login-with-twitter with .env feature toggle
5 add dev/prod mode
