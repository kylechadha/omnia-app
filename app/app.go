package app

import "net/http"

type Ioc struct {
	ConfigService  IConfigService
	DaysController IDaysController
}

type IConfigService interface {
	GetConfig(key string) (string, error)
}

type IDaysController interface {
	DayCreate(w http.ResponseWriter, r *http.Request) (error, int)
}
