package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kylechadha/omnia-app/app"
	"github.com/kylechadha/omnia-app/routes"

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

	// ? Should this be a config service ? ... probably...
	// Read the config data.
	rawConfig, err := ioutil.ReadFile("app/config.json")
	if err != nil {
		fmt.Println("Unable to read 'app/config.json'. Is the filepath correct?")
		os.Exit(1)
	}

	// Unmarshal the data into a JSON wrapper.
	var config map[string]string
	err = json.Unmarshal(rawConfig, &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
	n.Run(":" + port)

}
