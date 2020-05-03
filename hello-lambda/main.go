package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

type Event struct {
	Action string `json:"action"`
}

type ListAllNamesResponse struct {
	Error error
	Names []string
}

type NameAddedResponse struct {
	Error error
	Name  string
}

type UnknownActionResponse struct {
	Error string
}

func HandleRequest(ctx context.Context, event Event) (interface{}, error) {
	db_user, _ := os.LookupEnv("DB_USER")
	db_pwd, _ := os.LookupEnv("DB_PWD")
	db_host, _ := os.LookupEnv("DB_HOST")
	db_name, _ := os.LookupEnv("DB_NAME")
	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true&timeout=30s", db_user, db_pwd, db_host, db_name))
	er := db.Ping()
	if er != nil {
		log.Errorf("Failed to connect to db: %s", er)
		return nil, er
	}
	err := migrate_db(db)
	if err != nil {
		return "", err
	}

	switch event.Action {
	case "ADD_NEW":
		name, err := addNewName(db)
		return NameAddedResponse{Name: name, Error: err}, nil
	case "LIST_ALL":
		names, err := listAllNames(db)
		return ListAllNamesResponse{Names: names, Error: err}, nil
	}
	db.Close()

	return UnknownActionResponse{Error: "Unknown action"}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
