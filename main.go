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
	name, country, temp := weather.Location.Name,weather.Location.Country,weather.Current.TempC
	fmt.Printf("Name: %s  Country: %s  Temparature: %.1f\n",name,country,temp)
}