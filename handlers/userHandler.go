package handlers

import (
	"net/http"

	"github.com/ismailraqi/echoSQLboiler/db"

	"github.com/ismailraqi/echoSQLboiler/models"
	"github.com/labstack/echo"
	jwtModelsClaims "github.com/ismailraqi/echoSQLboiler/jwtmodelsclaims"

)

// UserRegister ....
func UserRegister(c echo.Context) (err error) {
	user := new(models.User)
	defer c.Request().Body.Close()
	if err = c.Bind(user); err != nil {
		return c.JSON(http.StatusInternalServerError, "somthing wrong please try again")
	}
	// data validation using ozzo-validation

	// err = user.Validate()
	// if err != nil {
	// 	return c.String(http.StatusBadRequest, "the length must be between 5 and 50")
	// }
	// Insert user into db
	check, us := db.InsertUser(*user)
	if check == false {
		return c.JSON(http.StatusOK, us)
	}
	return c.JSON(http.StatusBadRequest, "you email aready exist")

}

// UserLogin ....
func UserLogin(c echo.Context) error {
	useremail := c.FormValue("Email")
	userpassword := c.FormValue("Password")
	exist, userinfo := db.CheckLogin(useremail, userpassword)
	defer c.Request().Body.Close()
	// Throws unauthorized error
	if exist != true {
		return c.JSON(http.StatusUnauthorized, "please try again or contact support")
	}

	t, err := jwtModelsClaims.JWTCreator(userinfo)
	if err != nil {
		return c.String(http.StatusInternalServerError, "something wrong")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})

}
