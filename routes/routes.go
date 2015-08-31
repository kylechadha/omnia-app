package routes

import (
	"errors"
	"net/http"

	"github.com/kylechadha/omnia-app/app"
	"github.com/kylechadha/omnia-app/utils"

	"github.com/gorilla/mux"
)

func NewRouter(ioc *app.Ioc) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	// Static files.
	// NOTE: Folders will be served too -- dir tree not restricted
	router.PathPrefix("/assets/").Handler(http.FileServer(http.Dir(".")))

	// User routes.
	// u := ioc.UsersController
	// users := router.PathPrefix("/users").Subrouter()
	// users.Handle("/", utils.AppHandler(u.UserCreate))

	// Catch all route.
	router.NotFoundHandler = utils.AppHandler(func(w http.ResponseWriter, r *http.Request) (error, int) {
		return errors.New("Whoops, that's a 404."), http.StatusNotFound
	})

	return router
}
