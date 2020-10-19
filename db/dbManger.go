package db

import "database/sql"

//DbConnection is a func that's allow you ro open a connection with Database
func DbConnection() (*sql.DB, error) {
	//open db (regular sql open call) : Thanks to iRELG <3
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/pilots")
	return db, err

}
