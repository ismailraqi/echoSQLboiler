package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/ismailraqi/Golang-sqlboiler/routers"
	_ "github.com/lib/pq"
)

func main() {
	routers.StartRouters()
}
