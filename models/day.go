package models

type Day struct {
	DayId            int    `json:"DayId"`
	DayOfTheWeek     string `json:"DayOfTheWeek"`
	DayOfTheWeekType string `json:"DayOfTheWeekType"`
	SuccessfulWakeUp bool   `json:"SuccessfulWakeUp"`
	MorningWork      bool   `json:"MorningWork"`
	MorningWorkType  string `json:"MorningWorkType"`
	WorkedOut        bool   `json:"WorkedOut"`
	WorkedOutType    string `json:"WorkedOutType"`
	PlannedNextDay   bool   `json:"PlannedNextDay"`
}
