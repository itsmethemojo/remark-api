#!/bin/bash

docker-compose exec -T fpm bash -c 'cd /app; vendor/bin/phpcbf --standard=PSR2 src public; vendor/bin/phpcs --standard=PSR2 src public'
