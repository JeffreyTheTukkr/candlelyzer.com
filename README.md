# candlelyzer.com

Repository for the [candlelyzer.com](https://candlelyzer.com) project.

**Requirements**

```
golang          >= v1.22
golangci-lint   >= v1.59
timescaledb     >= v16.3
```

**Quick commands**

```
run    : go run main.go
test   : go test ./... [-v] [-cover]
lint   : golangci-lint run ./... [--fix]
```

## development

Documentation regarding the (local) development of this project. 

**Dotenv file**

```
# general
APP_ENV=development

# postgresql database
PG_HOST=localhost
PG_PORT=5432
PG_USER=__USERNAME_
PG_PASS=__PASSWORD__
PG_NAME=candlelyzer
```
