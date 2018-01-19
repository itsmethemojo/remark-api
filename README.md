# reMARK

TODO description

## howto to run it local in 2 minutes

start redis and mysql docker container
```
docker run --name local-redis -d redis ;\
docker run --name local-mariadb -e MYSQL_ROOT_PASSWORD=root -d -v $PWD/installation:/data mariadb
```

create the default config files for your app
```
installation/createLocalConfigs.sh
```

put the tables in the local database
```
docker exec -t local-mariadb bash -c "mysql -uroot -proot < /data/database-structure.sql"
```

start the app itself
```
docker run --rm --interactive --tty --volume $PWD:/app composer install ;\
docker build -t remark-api . ;\
docker run -td -p 80:8080 --name remark-api -v $(pwd):/var/www remark-api ;\
echo -e "\n\n   open this url: http://"$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' remark-api)":8080/status\n\n" ;\
docker exec -t remark-api php -S 0.0.0.0:8080 index.php
```

## what's next?

### configure the application

[more...](documentation/config.md)

### available API routes

[more...](documentation/routes.md)

### howto use the development tools with the container

[more...](documentation/tools.md)

### howto configure your webserver for it

[more...](https://www.slimframework.com/docs/start/web-servers.html)
