package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kylechadha/omnia-app/app"
	"github.com/kylechadha/omnia-app/models"
	"gopkg.in/mgo.v2/bson"
)

// Days Controller type.
type daysController struct {
	databaseService app.IDatabaseService
}

// Days Controller constructor function.
func NewDaysController(app *app.Ioc) *daysController {
	return &daysController{}
}

// DayCreate handler.
func (c *daysController) DayCreate(w http.ResponseWriter, r *http.Request) (error, int) {

	day := models.Day{}
	day.Id = bson.NewObjectId()

	json.NewDecoder(r.Body).Decode(&day)

	c.databaseService.Create("days", day)

	return nil, http.StatusOK
}
