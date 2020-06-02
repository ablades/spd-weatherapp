// SPD 1.5 Weather APP
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	u "net/url"
	"os"
	"time"
)

// url to pull images http://openweathermap.org/img/wn/10d@2x.png

// struct that holds relevant weather data
type weatherStats struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
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
	Name string `json:"name"`
}

// Get current time a user submits an entry
func getTime() {
	fmt.Println(time.Now())
}

//Send request to weather api
func apiRequest(cityName string, apiKey string) []byte {

	// build url query string
	url, _ := u.Parse("api.openweathermap.org/data/2.5/weather")
	url.Scheme = "https"
	q := url.Query()
	q.Add("appid", apiKey)
	q.Add("q", cityName)
	url.RawQuery = q.Encode()

	// Request
	fmt.Println(url.String())
	r, err := http.Get(url.String())
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	content, _ := ioutil.ReadAll(r.Body)

	return content
}

func main() {
	args := os.Args
	apiRequest(args[1], args[2])
	getTime()
}
