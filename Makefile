TEST_COMMAND= GIN_MODE=test go test -v -cover ./src/tests/...

start:
	DB_HOST=localhost go run ./src/main.go
setEnv:
	cp .env.example .env
test:
	$(TEST_COMMAND)
testAll:
	go clean -testcache && $(TEST_COMMAND)
doc:
	swag init --pd -d src/
migrate:
	migrate -database 'mysql://admin:admin@(localhost:3309)/member?charset=utf8&parseTime=true&loc=Local' -path ./src/db/migrations up
migrate_rollback:
	migrate -database 'mysql://admin:admin@(localhost:3309)/member?charset=utf8&parseTime=true&loc=Local' -path ./src/db/migrations down
