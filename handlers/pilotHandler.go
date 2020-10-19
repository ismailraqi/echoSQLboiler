package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/ismailraqi/Golang-sqlboiler/db"
	"github.com/ismailraqi/Golang-sqlboiler/models"
	"github.com/labstack/echo"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type pilot struct {
	ID   int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`
}

func (a pilot) Validate() error {
	return validation.ValidateStruct(&a,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Name, validation.Required, validation.Length(5, 50)),
	)
}

var database *sql.DB

func init() {
	db, err := db.DbConnection()
	if err != nil {
		fmt.Println(err)
	}
	database = db
}

//GetAllPilots to retreive all pilots
func GetAllPilots(c echo.Context) error {
	ctx := context.Background()
	ps, _ := models.Pilots().All(ctx, database)
	for _, p := range ps {
		fmt.Printf("Pilot Info: %+v\n", p)
	}
	return c.JSON(http.StatusOK, ps)
}

//GetOnePilots to retreive one pilot
func GetOnePilots(c echo.Context) error {
	ctx := context.Background()
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	ps, _ := models.Pilots(qm.Where("id=?", pID)).One(ctx, database)
	if ps == nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	return c.JSON(http.StatusOK, ps)
}
func insertPilot(pilot models.Pilot) error {
	db, err := db.DbConnection()
	if err != nil {
		panic(err)
	}
	err = pilot.Insert(context.Background(), db, boil.Infer())
	return err

}

//CreatePilot
func CreatePilot(c echo.Context) error {
	// var body pilot
	// err := body.Validate()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // bind data given fro user & check if there is an error
	// if err := c.Bind(&body); err != nil {
	// 	return err
	// }
	// // validates a structs exposed fields & check if there is an error
	// if err := c.Validate(body); err != nil {
	// 	return err
	// }
	defer c.Request().Body.Close()
	ID := c.FormValue("id")
	Name := c.FormValue("name")
	nPilot := new(models.Pilot)
	pid, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println(err)
	}
	nPilot.ID = pid
	nPilot.Name = Name
	t := insertPilot(*nPilot)
	if t != nil {
		return c.String(http.StatusBadRequest, "please try again somthing wrong")
	}
	return c.JSON(http.StatusOK, nPilot)
}
func DeletePilot() {

}
