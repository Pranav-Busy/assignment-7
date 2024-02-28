package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"regexp"
)

type Weatherreport struct {
	Name    string   `json:"name"`
	Weather string   `json:"weather"`
	Status  []string `json:"status"`
}

type GetWeatherFromJson struct {
	Page       uint            `json:"page"`
	Per_page   uint            `json:"per_page"`
	Total      uint            `json:"Total"`
	Total_page uint            `json:"total_pages"`
	Data       []Weatherreport `json:"data"`
}

var Result [][]string

//var Listofcities []Result;

func fetchData(apiURL, cityName string) {

	page := 1
	for {
		params := url.Values{}
		params.Add("name", cityName)
		params.Add("page", strconv.Itoa(page))
		fullURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())

		response, err := http.Get(fullURL)
		if err != nil {
			return
		}
		defer response.Body.Close()
		
	var weatherData GetWeatherFromJson

	
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&weatherData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
     
	 for _,val:=range weatherData.Data{
        var helper [] string
		re := regexp.MustCompile(`\D+`)
		val.Weather = re.ReplaceAllString(val.Weather, "")
		val.Status[0] = re.ReplaceAllString(val.Status[0], "")
		val.Status[1] = re.ReplaceAllString(val.Status[1], "")
		
		helper=append(helper,val.Name)
		helper=append(helper,val.Weather)
		helper=append(helper,val.Status...)
		Result=append(Result, helper)
	 }

	 if page >= int(weatherData.Total_page) {
		break
	}
	 page++

	}

}

func main() {
	apiURL := "https://jsonmock.hackerrank.com/api/weather/search"
	var cityName string
	log.Println("Enter the city name for which you want to search the weather")
	fmt.Scanf("%v", &cityName)
	fetchData(apiURL, cityName)
	
	for _,val:=range Result{

		fmt.Println(val)
	}
	fmt.Println(len(Result))
}
