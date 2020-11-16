package db

import (
	"context"

	"database/sql"

	"github.com/ismailraqi/echoSQLboiler/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

//DuplictedEmail ...
func DuplictedEmail(email string, db *sql.DB) bool {
	exist, err := models.Users(qm.Where(models.UserColumns.Email+"=?", email)).Exists(context.Background(), db)
	if err != nil {
		panic(err)
	}
	if exist == true {
		return true
	}
	return false
}
