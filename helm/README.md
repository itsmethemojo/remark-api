# remark-api

![Version: 1.0.0](https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square)

A Helm chart to run remark-api in Kubernetes

## installation

```
helm upgrade --install demo . --set postgresql.auth.password=super-secret-password --set app.env.DATABASE_PASSWORD=super-secret-password
```

## Source Code

* <https://github.com/itsmethemojo/remark-api>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.bitnami.com/bitnami | postgresql | 12.1.3 |
| https://itsmethemojo.github.io/helm-charts/ | app(basic-web-app) | 1.1.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| app.containerPort | int | `8080` |  |
| app.env.ACCESS_CONTROL_ALLOW_CREDENTIALS | string | `""` | see [default.env](https://github.com/itsmethemojo/remark-api/blob/master/default.env) |
| app.env.ACCESS_CONTROL_ALLOW_HEADERS | string | `""` | see [default.env](https://github.com/itsmethemojo/remark-api/blob/master/default.env) |
| app.env.ACCESS_CONTROL_ALLOW_METHODS | string | `""` | see [default.env](https://github.com/itsmethemojo/remark-api/blob/master/default.env) |
| app.env.ACCESS_CONTROL_ALLOW_ORIGIN | string | `""` | see [default.env](https://github.com/itsmethemojo/remark-api/blob/master/default.env) |
| app.env.API_PATH_PREFIX | string | `""` | see [default.env](https://github.com/itsmethemojo/remark-api/blob/master/default.env) |
| app.env.APP_DOMAIN | string | `"localhost"` |  |
| app.env.APP_PORT | string | `"8080"` |  |
| app.env.APP_SCHEMA | string | `"http"` |  |
| app.env.CORS_ENABLED | string | `"1"` |  |
| app.env.DATABASE_CONNECT_RETRY_COUNT | string | `""` | see [default.env](https://github.com/itsmethemojo/remark-api/blob/master/default.env) |
| app.env.DATABASE_CONNECT_WAIT_INTERVAL | string | `""` | see [default.env](https://github.com/itsmethemojo/remark-api/blob/master/default.env) |
| app.env.DATABASE_HOST | string | `"remark-api-database"` |  |
| app.env.DATABASE_NAME | string | `"remark-api"` |  |
| app.env.DATABASE_PORT | string | `""` | see [default.env](https://github.com/itsmethemojo/remark-api/blob/master/default.env) |
| app.env.DATABASE_SSLMODE | string | `""` | see [default.env](https://github.com/itsmethemojo/remark-api/blob/master/default.env) |
| app.env.DATABASE_TIMEZONE | string | `""` | see [default.env](https://github.com/itsmethemojo/remark-api/blob/master/default.env) |
| app.env.DATABASE_USERNAME | string | `"remark-api"` |  |
| app.env.DEMO_TOKENS | string | `""` | see [default.env](https://github.com/itsmethemojo/remark-api/blob/master/default.env) |
| app.env.LOGIN_PROVIDER | string | `""` | see [default.env](https://github.com/itsmethemojo/remark-api/blob/master/default.env) |
| app.env.SWAGGER_PATH | string | `"/swagger"` |  |
| app.env.TEST_MODE | string | `"false"` |  |
| app.image.repository | string | `"ghcr.io/itsmethemojo/remark-api"` |  |
| app.image.tag | string | `"sha-c2b84a7"` |  |
| app.livenessProbe.httpGet.path | string | `"/health"` |  |
| app.nameOverride | string | `"remark-api"` |  |
| app.readinessProbe.httpGet.path | string | `"/health"` |  |
| postgresql.auth.database | string | `"remark-api"` |  |
| postgresql.auth.username | string | `"remark-api"` |  |
| postgresql.enabled | bool | `true` | to use a separate deployed database set to false here |
| postgresql.fullnameOverride | string | `"remark-api-database"` |  |
| postgresql.primary.persistence.enabled | bool | `false` | for production use persistence should be enabled |

## update docs

```
docker run --rm -v $(pwd):/app -w/app jnorwood/helm-docs -t helm-docs-template.gotmpl
```

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)