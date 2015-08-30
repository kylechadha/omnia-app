package app

import "net/http"

type Ioc struct {
	UserController IUserController
	PostController IPostController
}

type IUserController interface {
	UserCreate(w http.ResponseWriter, r *http.Request) (error, int)
}

type IPostController interface {
	PostCreate(w http.ResponseWriter, r *http.Request) (error, int)
}
