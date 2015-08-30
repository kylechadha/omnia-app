package main

import (
	"github.com/codegangsta/negroni"
	"github.com/kylechadha/omnia-app/app"
	"github.com/kylechadha/omnia-app/routes"
)

func main() {

	// Object Graph
	// ----------------------------
	ioc := app.Ioc{}

	// Router
	// ----------------------------
	router := routes.NewRouter(&ioc)
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)
	n.UseHandler(router)

	// Server
	// ----------------------------
	n.Run(":3000")

}
