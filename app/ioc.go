package app

import "net/http"

type Ioc struct {
	UsersController IUsersController
	ItemController  IItemController
}

type IUsersController interface {
	UserCreate(w http.ResponseWriter, r *http.Request) (error, int)
}

type IItemController interface {
	ItemCreate(w http.ResponseWriter, r *http.Request) (error, int)
}
