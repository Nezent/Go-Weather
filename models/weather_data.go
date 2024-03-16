package models

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC          float64 `json:"temp_c"`
		LastTimeUpdate string  `json:"last_updated"`
		Condition      struct {
			Text string `json:"text"`
		} `json:"condition"`
		FeelsLikeC float64 `json:"feelslike_c"`
		Humidity   int     `json:"humidity"`
	} `json:"current"`
}