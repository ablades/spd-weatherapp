// SPD 1.5 Weather APP
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	u "net/url"
	"os"
)

func apiRequest(cityName string, apiKey string) {

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

	fmt.Println(string(content))
}

func main() {
	args := os.Args
	apiRequest(args[1], args[2])

}
