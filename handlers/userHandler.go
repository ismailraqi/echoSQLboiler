package handlers

import (
	"github.com/ismailraqi/echoSQLboiler/db"
	"net/http"

	"github.com/ismailraqi/echoSQLboiler/models"
	"github.com/labstack/echo"
)

// UserRegister ....
func UserRegister(c echo.Context) (err error) {

	user := new(models.User)
	//close body after finish whit request
	defer c.Request().Body.Close()
	//decoding data based on the Content-Type header.
	if err = c.Bind(user); err != nil {
		return c.JSON(http.StatusInternalServerError, "somthing wrong please try again")
	}
	// data validation using validator
	err = user.Validate()
	if err != nil {
		return c.String(http.StatusBadRequest, "the length must be between 5 and 50")
	}
	// Insert user into db
	check,err := db.InsertUser(*user)
	if check == false {
		return c.JSON(http.StatusOK, us)
	}
	return c.JSON(http.StatusBadRequest, "you email aready exist")

}
