# reMARK

api to store bookmarks and sort/rank them over time

the frontend can be found [here](https://github.com/itsmethemojo/remark)

the initial version was wriiten in php [see](https://github.com/itsmethemojo/remark-api/tree/21f1ad861827053c76f328d39174b97c225cc5d4)

## prequisites

[go-task](https://github.com/go-task/task)
[docker](https://www.docker.com/get-started)

## installation

`task dz:init`

## start server local

`docker-compose up -d --build`

## see swagger api documentation

[open](http://localhost:8080/swagger/index.html)

## interesting task targets

`task lint`

## run tests

```
docker-compose -f docker-compose.yml -f docker-compose-testing.yml up -d --build
./test.sh
```

## TODO

1 improve test suite

3 document available environment parameters with defaults

4 add github actions

5 check conneciton pooling https://gorm.io/docs/generic_interface.html#Connection-Pool

6 improve naming of AllBookmarkData

7 add debug logs

8 add dev/prod mode


