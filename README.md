# CIAM-golang

## Description

Basic gin

## Environment

- Docker: v20.10.7
- Go: v1.16.5
- golangci-lint: v1.41.0

## install
```bash
# install makefile
brew install make

# install migrate (參考:https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
brew install golang-migrate
```


## Command

```bash
# Launch go server
make start

# Set runtime enviroment
make setEnv

# Run unit test with cache
make test

# Run unit test without cache
make testAll

# Run golangci-lint
make lint

# Run swagger doc
make doc

# Run migrate
make migrate

# Run migrate rollback
make migrate_rollback
```

## Postman

- domain: http://localhost
For local testing