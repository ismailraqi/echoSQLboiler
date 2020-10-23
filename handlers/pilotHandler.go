package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/ismailraqi/echoSQLboiler/db"
	"github.com/ismailraqi/echoSQLboiler/models"
	"github.com/labstack/echo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type plt struct {
	ID   int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`
}

func (a plt) Validate() error {
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

//CreatePilot to create a new pilot
func CreatePilot(c echo.Context) error {
	defer c.Request().Body.Close()
	ID := c.FormValue("id")
	Name := c.FormValue("name")
	nPilot := new(models.Pilot)
	pid, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println(err)
	}
	p := plt{
		ID:   pid,
		Name: Name,
	}
	er := p.Validate()
	if er != nil {
		return c.String(http.StatusBadRequest, "the length must be between 5 and 50")
	}
	nPilot.ID = pid
	nPilot.Name = Name
	t := db.InsertPilot(*nPilot)
	if t != nil {
		return c.String(http.StatusBadRequest, "error detected please try again :/")
	}
	return c.JSON(http.StatusOK, nPilot)
}

//DeletePilot to delete a selected pilot
func DeletePilot(c echo.Context) error {
	ctx := context.Background()
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := models.Pilots(qm.Where("id=?", id)).DeleteAll(ctx, database)
	if err != nil {
		return c.String(http.StatusBadRequest, "error detected please try again :/")
	}
	return c.NoContent(http.StatusNoContent)
}

//UpdatePilot to delete a selected pilot
func UpdatePilot(c echo.Context) error {
	ID, _ := strconv.Atoi(c.FormValue("id"))
	Name := c.FormValue("name")
	nPilot := new(models.Pilot)
	nPilot.ID = ID
	nPilot.Name = Name
	p := plt{
		ID:   ID,
		Name: Name,
	}
	er := p.Validate()
	if er != nil {
		return c.String(http.StatusBadRequest, "the length must be between 5 and 50")
	}
	defer c.Request().Body.Close()
	errdb := db.UpdatePilot(*nPilot)
	if errdb != nil {
		return c.String(http.StatusBadRequest, "Something wrong with dbManager")
	}
	return c.JSON(http.StatusOK, "Updated successfully")
}
