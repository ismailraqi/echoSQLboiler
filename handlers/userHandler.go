package handlers

import (
	"net/http"

	"github.com/ismailraqi/echoSQLboiler/db"

	jwtModelsClaims "github.com/ismailraqi/echoSQLboiler/jwtmodelsclaims"
	"github.com/ismailraqi/echoSQLboiler/models"
	"github.com/labstack/echo"
)

// UserRegister ....
func UserRegister(c echo.Context) (err error) {
	user := new(models.User)
	defer c.Request().Body.Close()
	if err = c.Bind(user); err != nil {
		return c.JSON(http.StatusInternalServerError, "somthing wrong please try again")
	}
	// data validation using ozzo-validation
	u := models.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	er := u.Validate()
	if er != nil {
		return c.String(http.StatusBadRequest, "the length must be between 5 and 50")
	}
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
	if exist != true {
		return c.JSON(http.StatusUnauthorized, "something wrong")
	}

	t, err := jwtModelsClaims.JWTCreator(userinfo)
	if err != nil {
		return c.String(http.StatusInternalServerError, "something wrong")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})

}
