package main

import (
	"os"

	"github.com/kylechadha/omnia-app/app"
	"github.com/kylechadha/omnia-app/controllers"
	"github.com/kylechadha/omnia-app/routes"
	"github.com/kylechadha/omnia-app/services"

	"github.com/codegangsta/negroni"
)

func main() {

	// Config
	// ----------------------------

	// Set the port.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// ? Should this be a database service? Yes, move out soon
	// Database
	// ----------------------------
	// mongoUrl, mongoUrlOk := config["mongourl"]

	// if mongoUrlOk {
	// 	sessions, err := mgo.Dial(mongoUrl)
	// 	if err != nil {
	// 		fmt.Println("Unable to connect to the Mongo DB.")
	// 		os.Exit(1)
	// 	}
	// } else {
	// 	fmt.Println("Config file does not include a 'mongourl'.")
	// 	os.Exit(1)
	// }

	// Object Graph
	// ----------------------------
	app := app.Ioc{}
	app.ConfigService = services.NewConfigService()
	app.DaysController = controllers.NewDaysController(&app)

	// Router
	// ----------------------------
	router := routes.NewRouter(&app)

	// Do we want to use negroni again? Meh ... what about another way to do recovery
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)
	n.UseHandler(router)

	// Server
	// ----------------------------
	n.Run(":" + port)

}
