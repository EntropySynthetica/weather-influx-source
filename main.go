package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type OpenWeatherMapAPI struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`

	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`

	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`

	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`

	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`

	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func main() {

	// Get vars from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api_key := os.Getenv("API_KEY")
	location_id := os.Getenv("LOCATION_ID")
	units := os.Getenv("UNITS")
	poll_url := "https://api.openweathermap.org/data/2.5/weather?id=" + location_id + "&units=" + units + "&appid=" + api_key

	// Poll the API
	response, err := http.Get(poll_url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	response_data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Show the data we got
	fmt.Println("Raw Data")
	fmt.Println(string(response_data))
	fmt.Println("")

	var weather_data OpenWeatherMapAPI
	json.Unmarshal([]byte(response_data), &weather_data)

	fmt.Println(weather_data.Name)
	fmt.Printf("%+v\n", weather_data.Main)
	fmt.Println(weather_data.Main.Temp)

}
