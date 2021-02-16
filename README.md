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

1 use dotenv to retrieve database string
2 return error class that includes the api response code
3 see if import naming can be improved `repository "../repositories/bookmark"`
