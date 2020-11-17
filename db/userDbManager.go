package db

import (
	"context"

	jwtModelsClaims "github.com/ismailraqi/echoSQLboiler/jwtmodelsclaims"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

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
	u := &models.User{}
	if isDuplicated == false {
		user.Password = security.PassSHA256(user.Password)
		err = user.Insert(context.Background(), db, boil.Infer())
		u = &models.User{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password,
		}
		return false, *u
	}
	return true, *u

}

// CheckLogin is a func that return a bool value (true) if exist and (false) if not
func CheckLogin(Email, Password string) (bool, jwtModelsClaims.JwtCustomClaims) {
	db, err := DbConnection()
	if err != nil {
		panic(err)
	}
	// get the user based on given email
	logged, err := models.Users(qm.Where(models.UserColumns.Email+"= ?", Email)).One(context.Background(), db)
	if err != nil {
		panic(err)
	}
	// store password with security usin SHA256
	Password = security.PassSHA256(Password)
	claim := new(jwtModelsClaims.JwtCustomClaims)
	if err != nil {
		panic(err)
	}
	claim.Email = logged.Email
	claim.ID = logged.ID
	claim.Username = logged.Username
	if logged.Password == Password {
		return true, *claim
	}
	return false, *claim

}
