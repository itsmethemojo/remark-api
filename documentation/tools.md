[back to README](../README.md)

## use PHP Codesniffer as pre-commit

install a pre-commit hook that runs a PHP codesniffer
```
hooks/install-hooks.sh
```

you can run the codesniffer by yourself by using this command

it will run  **phpcbf** and **phpcs** in the php-fpm container
```
hooks/codesniffer.sh
```


## use phpunit
run phpunit in the php-fpm container
```
docker-compose exec -T fpm bash -c 'cd /app; vendor/bin/phpunit .'
```

## use Composer
you can easily do composer actions with the composer container like this
```
docker run --rm --interactive --tty --volume $PWD:/app composer update slim/slim
docker run --rm --interactive --tty --volume $PWD:/app composer require whatever
```
