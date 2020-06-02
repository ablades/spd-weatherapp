// SPD 1.5 Weather APP
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	u "net/url"
)

func apiRequest(cityName string, apiKey string) {

	// build url query string
	url, _ := u.Parse("api.openweathermap.org/data/2.5/weather")
	q := url.Query()
	q.Add("q", cityName)
	q.Add("appid", apiKey)
	url.RawQuery = q.Encode()

	// Request
	r, err := http.Get(url.String())
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	content, _ := ioutil.ReadAll(r.Body)

	fmt.Println(content)
}

func main() {

}
