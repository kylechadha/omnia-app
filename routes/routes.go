package routes

import (
	"errors"
	"net/http"

	"github.com/kylechadha/omnia-app/app"
	"github.com/kylechadha/omnia-app/utils"

	"github.com/gorilla/mux"
)

func NewRouter(app *app.Ioc) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	// Days routes.
	d := app.DaysController
	days := router.PathPrefix("/api/days").Subrouter()
	days.Handle("/", utils.AppHandler(d.DayCreate))

	// Static files.
	router.PathPrefix("/libs").Handler(utils.RestrictDir(http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/scripts").Handler(utils.RestrictDir(http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/styles").Handler(utils.RestrictDir(http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/views")))

	// Catch all route.
	// Actually we'll need a catch all route to send everything to the angular frontend to handle routing...
	router.NotFoundHandler = utils.AppHandler(func(w http.ResponseWriter, r *http.Request) (error, int) {
		return errors.New("Whoops, that's a 404."), http.StatusNotFound
	})

	return router
}
