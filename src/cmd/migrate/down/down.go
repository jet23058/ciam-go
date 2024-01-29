// cmd/migrate/down.go
package main

import (
	"ciam_exam/src/config"

	log "github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config.InitOs()

	m, err := migrate.New(
		"file://internal/dao/migrations",
		config.DbInfo(),
	)
	if err != nil {
		log.Warn("internal/dao/migrations - migrate.New")
		log.Fatal(err)
	}

	if err := m.Down(); err != nil {
		log.Warn("internal/dao/migrations - m.Down()")
		log.Fatal(err)
	}
}
