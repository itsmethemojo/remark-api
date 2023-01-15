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

2 refactor user mapping since login provider will turn username string now, username have to be added in user table only when calling remark and nothing is found (then query user and add if missing)

3 document available environment parameters with defaults

4 add github actions

5 return error class that includes the api response code

6 improve naming of AllBookmarkData

7 add debug logs

8 add dev/prod mode


