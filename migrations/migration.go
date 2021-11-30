// migrate create -ext sql -dir migration/file -seq init_schema

package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func migrateSql() *migrate.Migrate {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations/file",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}

	return m
}

func RunMigrate() {

	migrator := migrateSql()
	defer func() {
		if r := recover(); r != nil {
			log.Println("Migration ", fmt.Sprintf("%v", r))
		}
	}()

	err := migrator.Up()
	if err != nil {
		v, d, _ := migrator.Version()
		if d {
			migrator.Force(int(v) - 1)
		}

		panic(err)
	}
}
