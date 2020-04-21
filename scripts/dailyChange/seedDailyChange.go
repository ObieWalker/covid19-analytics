package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/ObieWalker/covid19-analytics/helper"
	"github.com/ObieWalker/covid19-analytics/dao"
	log "github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}

	resp, err := http.Get(os.Getenv("HISTORY_URL"))

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
		fortnightData []float64
	}

	// collection := helper.GetCollection()
	collection := helper.ConnectDB()
	hd := HistoryData{}
	fmt.Println("Adding daily update to DB...")
	for _, e := range data {
		if e["province"] == nil{
			hd.Country = e["country"].(string)
			cases := e["timeline"].(map[string]interface{})["cases"]
	
			var s []float64
	
			for _, element := range cases.(map[string]interface{}) {
				value, _ := element.(float64)
				s = append(s, value)
			}
			sort.Float64s(s)
			fortnight := s[len(s)-15:]
	
			dao.UpdateWeekCases(collection, hd.Country, getDailyDifference(fortnight))
		} else {
			hd.Country = strings.Title(e["province"].(string))
			cases := e["timeline"].(map[string]interface{})["cases"]
	
			var s []float64
	
			for _, element := range cases.(map[string]interface{}) {
				value, _ := element.(float64)
				s = append(s, value)
			}
			sort.Float64s(s)
			fortnight := s[len(s)-15:]
	
			dao.UpdateWeekCases(collection, hd.Country, getDailyDifference(fortnight))
		}
	}
		m := map[string][]float64{
		"Côte d'Ivoire": {62, 26, 35, 60, 36, 53, 52, 12, 16, 0, 34, 113, 46},
		"Saint Pierre Miquelon": {0,0,0,0,0,0,0,0,0,0,0,0, 0,0},
		"Turks and Caicos Islands": {3,0,0,0,0,3,0,1, 0,0,1, 0,0,0},
		"Lao People's Democratic Republic": {1,2,1,1,0,2,1,0,0,0,0,0,0,0},
		"Isle of Man": {12,11,8,32,11,25,2,14,12, 2,28,7,6,1},
		"Curaçao": {2,0,1,0,0,0,0,0,0,0,0,0,0,0},
		"Holy See (Vatican City State)": {0,0,1,0,0,0,0,0,0,0,0,0,0,0},
		"Saint Martin": {0,0,0,0,0,0,0,0,0,0,3,0,0,2},
		"St. Barth": {0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		"Réunion": {9,4,0,206,1,2,0,0,0,3,8,5,1},
		}

		for country, fortnight := range m {
			dao.UpdateWeekCases(collection, country, fortnight)
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
