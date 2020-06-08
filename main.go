// SPD 1.5 Weather APP
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	u "net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
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

//TODO: Create landing search page
//TODO: Create template to hold weather data

// Get current time a user submits an entry
func getTime() {
	fmt.Println(time.Now())
}

func readFile(fileName string) {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(fileContents)
}

func writeToFile(fileName string, content string) {
	bytesToWrite := []byte(content)

	err := ioutil.WriteFile(fileName, bytesToWrite, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//Send request to weather api
func apiRequest(cityName string, apiKey string) weatherStats {

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

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weather = weatherStats{}

	json.Unmarshal(content, &weather)

	return weather
}

func main() {

	// Load Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get args
	args := os.Args

	// Send request
	w := apiRequest(args[1], os.Getenv("WEATHER_KEY"))

	fmt.Printf("%+v\n", w)

	getTime()
}
