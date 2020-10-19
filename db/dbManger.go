package db

import (
	"context"
	"database/sql"

	"github.com/ismailraqi/Golang-sqlboiler/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

//DbConnection is a func that's allow you ro open a connection with Database
func DbConnection() (*sql.DB, error) {
	//open db (regular sql open call) : Thanks to iRELG <3
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/pilots")
	return db, err

}

//InsertPilot is a func that's allow you to insert into db
func InsertPilot(pilot models.Pilot) error {
	db, err := DbConnection()
	if err != nil {
		panic(err)
	}
	err = pilot.Insert(context.Background(), db, boil.Infer())
	return err

}
