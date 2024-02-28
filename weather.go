package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

// Please obtain a free api key from api.weatherapi.com and place it here.
const apiKey = "API_KEY_GOES_HERE"

type WeatherResponse struct {
	Location struct {
		Name    string  `json:"name"`
		Region  string  `json:"region"`
		Country string  `json:"country"`
		Lat     float64 `json:"lat"`
		Lon     float64 `json:"lon"`
		TzID    string  `json:"tz_id"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		UV         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: weather <city>")
		os.Exit(1)
	}

	city := os.Args[1]

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, city)
	fmt.Println("API Request URL:", url)

	client := resty.New()
	response, err := client.R().Get(url)
	if err != nil {
		fmt.Println("Error fetching weather data:", err)
		os.Exit(1)
	}

	if response.StatusCode() != 200 {
		fmt.Println("Error:", response.Status())
		os.Exit(1)
	}

	weather := WeatherResponse{}
	if err := json.Unmarshal(response.Body(), &weather); err != nil {
		fmt.Println("Error parsing weather data:", err)
		os.Exit(1)
	}

	fmt.Printf("%s:\n", weather.Location.Country)
	fmt.Printf("Weather in %s:\n", weather.Location.Name)
	fmt.Printf("Latitude: %.1f°\n", weather.Location.Lat)
	fmt.Printf("Longitude: %.1f°\n", weather.Location.Lon)
	fmt.Printf("Temperature: %.1f°C\n", weather.Current.TempC)
	fmt.Printf("Temperature: %.1f°F\n", weather.Current.TempF)
	fmt.Printf("Humidity: %d%%\n", weather.Current.Humidity)
	fmt.Printf("Wind Speed: %.1f km/h\n", weather.Current.WindKph)
}
