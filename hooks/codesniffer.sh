#!/bin/bash

docker exec -t $(docker-compose ps | grep fpm_ | awk '{print $1}') bash -c 'cd /app; vendor/bin/phpcbf --standard=PSR2 src public; vendor/bin/phpcs --standard=PSR2 src public'
