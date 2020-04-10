package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"sort"

	"github.com/ObieWalker/covid19-analytics/helper"
	"github.com/ObieWalker/covid19-analytics/models"
	log "github.com/sirupsen/logrus"
)

func main() {

	resp, err := http.Get("https://corona.lmao.ninja/v2/historical")

	if err != nil {
		log.Fatal(err)
	}

	var data []map[string]interface{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	obj := string(body)
	dataError := json.Unmarshal([]byte(obj), &data)

	if dataError != nil {
		log.Fatal(dataError)
	}

	type HistoryData struct {
		Country       string
		WeekData      []float64
		fortnightData []float64
	}

	collection := helper.GetCollection()
	hd := HistoryData{}
	fmt.Println("Adding daily update to DB...")
	for _, e := range data {
		hd.Country = e["country"].(string)
		cases := e["timeline"].(map[string]interface{})["cases"]

		var s []float64

		for _, element := range cases.(map[string]interface{}) {
			value, _ := element.(float64)
			s = append(s, value)
		}
		sort.Float64s(s)
		fortnight := s[len(s)-15:]
		week := s[len(s)-8:]

		models.UpdateWeekCases(collection, hd.Country, getDailyDifference(fortnight), getDailyDifference(week))
	}
	fmt.Println("Done.")
}

func getDailyDifference(sli []float64) []float64 {
	var diffSlice []float64
	for i := 0; i < len(sli)-1; i++ {
		val := math.Abs(sli[i] - sli[i+1])
		diffSlice = append(diffSlice, val)
	}
	return diffSlice
}
