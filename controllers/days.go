package controllers

import (
	"net/http"

	"github.com/kylechadha/omnia-app/app"
)

// Days Controller type.
type daysController struct {
}

// Days Controller constructor function.
func NewDaysController(app *app.Ioc) *daysController {
	return &daysController{}
}

// DayCreate handler.
func (c *daysController) DayCreate(w http.ResponseWriter, r *http.Request) (error, int) {
	return nil, http.StatusOK
}
