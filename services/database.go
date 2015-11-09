package services

import (
	"fmt"

	"github.com/kylechadha/omnia-app/app"
	"gopkg.in/mgo.v2"
)

// Database Service type.
type databaseService struct {
	session *mgo.Session
}

// Database Service constructor function.
func NewDatabaseService(app *app.Ioc) *databaseService {
	mongoUrl, err := app.ConfigService.GetConfig("mongourl")
	if err != nil {
		fmt.Println("Config file does not include a 'mongourl'.")
		panic(err)
	}

	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		fmt.Println("Unable to connect to the Mongo DB.")
		panic(err)
	}

	return &databaseService{session}
}

func (d *databaseService) Create(collection string, data interface{}) error {

	// Write the user to mongo.
	err := d.session.DB("omnia_app").C(collection).Insert(data)
	if err != nil {
		return err
	}

	return nil
}
