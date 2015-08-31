package app

import "net/http"

type Ioc struct {
	UsersController IUsersController
	PostsController IPostsController
}

type IUsersController interface {
	UserCreate(w http.ResponseWriter, r *http.Request) (error, int)
}

type IPostsController interface {
	PostCreate(w http.ResponseWriter, r *http.Request) (error, int)
}
