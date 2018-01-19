#!/bin/bash

configPath=$(readlink -f $(dirname $(readlink -f $0))"/../config")

echo "" > $configPath"/remark.ini"
echo "DEBUG_MODE=true" >> $configPath"/remark.ini"
echo "MYSQL_USERNAME=root" >> $configPath"/remark.ini"
echo "MYSQL_PASSWORD=root" >> $configPath"/remark.ini"
echo "MYSQL_DATABASE=remark" >> $configPath"/remark.ini"
echo "MYSQL_HOST=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' local-mariadb)" >> $configPath"/remark.ini"
echo "REDIS_HOST=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' local-redis)" >> $configPath"/remark.ini"
echo "REDIS_PREFIX=remark_" >> $configPath"/remark.ini"

echo "" > $configPath"/login.ini"
echo "DUMMY_MODE=true" >> $configPath"/login.ini"
echo "LIFETIME=1000" >> $configPath"/login.ini"
