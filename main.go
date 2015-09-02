package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kylechadha/omnia-app/app"
	"github.com/kylechadha/omnia-app/routes"
	"gopkg.in/mgo.v2"

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

	// Read the config data.
	rawConfig, err := ioutil.ReadFile("app/config.json")
	if err != nil {
		fmt.Println("Unable to read 'app/config.json'. Are you sure it exists?")
		os.Exit(1)
	}

	// Unmarshal the data into a JSON wrapper.
	var config map[string]string
	err = json.Unmarshal(rawConfig, &config)
	if err != nil {
		panic(err)
	}

	// Database
	// ----------------------------
	mongoUrl, mongoUrlOk := config["mongourl"]

	if mongoUrlOk {
		_, err := mgo.Dial(mongoUrl)
		if err != nil {
			panic(err)
		}

	}

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
