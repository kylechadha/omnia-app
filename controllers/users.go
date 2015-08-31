package controllers

import (
	"net/http"

	"print.io/printio-widget-go/app"
)

// Users Controller type.
type usersController struct {
}

// Users Controller constructor function.
func NewUsersController(ioc *app.Ioc, templatePath string, templateName string) *usersController {

	return &usersController{}
}

// UserCreate handler.
func (c *usersController) UserCreate(w http.ResponseWriter, r *http.Request) (error, int) {

	// If we do use angular -- this won't write to templates, it'll just return the data itself
	// ie: the goapp is just an API

	return nil, http.StatusOK
}
