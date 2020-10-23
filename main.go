package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/ismailraqi/echoSQLboiler/cmd"
	_ "github.com/lib/pq"
)

func main() {
	cmd.Execute()
}
