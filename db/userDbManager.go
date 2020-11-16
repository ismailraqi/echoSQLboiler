package db

import (
	"context"

	"github.com/ismailraqi/echoSQLboiler/security"

	"github.com/ismailraqi/echoSQLboiler/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

//InsertUser ....
func InsertUser(user models.User) (bool, models.User) {
	db, err := DbConnection()
	if err != nil {
		panic(err)
	}
	isDuplicated := DuplictedEmail(user.Email, db)
	if isDuplicated == false {
		user.Password = security.PassSHA256(user.Password)
		err = user.Insert(context.Background(), db, boil.Infer())
		u := &models.User{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		}
		return false, *u
	}
	return true, *u

}
