package app

import "net/http"

type Ioc struct {
	ConfigService   IConfigService
	DatabaseService IDatabaseService
	DaysController  IDaysController
}

type IConfigService interface {
	GetConfig(key string) (string, error)
}

type IDatabaseService interface {
	Create(collection string, data interface{}) error
}

type IDaysController interface {
	DayCreate(w http.ResponseWriter, r *http.Request) (error, int)
}
