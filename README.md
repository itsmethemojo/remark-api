# reMARK

TODO description

## howto to run it local in 2 minutes

```
docker build -t fpm-mysql-redis docker/fpm; \
docker-compose stop; \
docker-compose up -d --force-recreate; \
docker-compose ps; \
echo -e "\n\nopen: http://"$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker-compose ps | grep nginx_ | awk '{print $1}'))"/\n\n"
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
