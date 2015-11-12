package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	return &daysController{app.DatabaseService}
}

// curl -XPOST -H 'Content-Type: application/json' -d '{"dayOfTheWeek": "Monday", "dayOfTheWeekType": "Weekday", "successfulWakeUp": true, "morningWork": true, "morningWorkType": "omnia app", "workedOut": true, "workedOutType": "swam", "plannedNextDay": true}' http://localhost:3000/api/day
// DayCreate handler.
func (c *daysController) DayCreate(w http.ResponseWriter, r *http.Request) (error, int) {

	// Create a new Day struct and set the ObjectId.
	day := models.Day{}
	day.Id = bson.NewObjectId()

	// Decode the JSON onto the struct.
	json.NewDecoder(r.Body).Decode(&day)

	// Create the Day via the Database Service.
	err := c.databaseService.Create("days", day)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusCreated
}

func (c *daysController) DayFind(w http.ResponseWriter, r *http.Request) (error, int) {

	log.Println("In day find")

	// Get the day ID from the params.
	vars := mux.Vars(r)
	dayId := vars["id"]

	// Verify the ID is a valid ObjectId.
	if !bson.IsObjectIdHex(dayId) {
		return errors.New("Invalid ID format."), http.StatusInternalServerError
	}

	// Get the Day via the Database Service.
	dayOId := bson.ObjectIdHex(dayId)
	log.Println(dayOId)
	day, err := c.databaseService.Find("days", dayOId, models.Day{})
	if err != nil {
		return err, http.StatusNotFound
	}

	// Encode the Account struct as JSON.
	json, err := json.Marshal(day)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	// Write the JSON to the response.
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

	return nil, http.StatusOK
}
