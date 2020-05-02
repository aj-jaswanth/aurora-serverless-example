package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	log "github.com/sirupsen/logrus"
)

func addNewName(db *sql.DB) (string, error) {
	name := randomName()
	stmt, _ := db.Prepare("insert into fake_names values (?)")
	_, err := stmt.Exec(name)
	if err != nil {
		log.Errorf("Failed to add a new name: %s", err)
		return "", err
	}
	log.Infof("Added new name: %s", name)
	return name, nil
}

func listAllNames(db *sql.DB) ([]string, error) {
	stmt, _ := db.Prepare("select * from fake_names")
	rows, err := stmt.Query()
	if err != nil {
		log.Errorf("Failed to list all names: %s", err)
		return nil, err
	}
	names := []string{}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		names = append(names, name)
	}
	rows.Close()
	return names, nil
}

func migrate_db(db *sql.DB) error {
	log.Info("starting migrations")
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Errorf("failed to run migrations: %s", err)
		return err
	}
	m.Up()
	return nil
}
