package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Nezent/Go-Weather/models"
)

func main() {
	res, err := http.Get("https://api.weatherapi.com/v1/current.json?key=affdacaf8348487eb5c132555241503&q=Dhaka&aqi=yes")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Can't Get Weather Forcast!")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather models.Weather
	err = json.Unmarshal(body,&weather)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Temparature Update:\nCountry: %s, %s\n",weather.Location.Name,weather.Location.Country)
	fmt.Printf("Time: %s\nTemparature: %0.1f, Feels Like %0.1f\nHumidity: %d\nCondition: %s\n",
	weather.Current.LastTimeUpdate,weather.Current.TempC,
	weather.Current.FeelsLikeC,weather.Current.Humidity,
	weather.Current.Condition.Text)
}