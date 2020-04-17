package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Country ...
type Country struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Country     string             `bson:"country,omitempty"`
	CountryInfo struct {
		ID   int64   `bson:"_id,omitempty"`
		Iso2 string  `bson:"iso2,omitempty"`
		Iso3 string  `bson:"iso3,omitempty"`
		Lat  float64 `bson:"lat,omitempty"`
		Long float64 `bson:"long,omitempty"`
		Flag string  `bson:"flag,omitempty"`
	}
	Updated             int64   `bson:"updated,omitempty"`
	Cases               int64   `bson:"cases,omitempty"`
	TodayCases          int64   `bson:"todayCases,omitempty"`
	Deaths              int64   `bson:"deaths,omitempty"`
	TodayDeaths         int64   `bson:"todayDeaths,omitempty"`
	Recovered           int64   `bson:"recovered,omitempty"`
	Active              int64   `bson:"active,omitempty"`
	Critical            int64   `bson:"critical,omitempty"`
	CasesPerOneMillion  float64 `bson:"casesPerOneMillion,omitempty"`
	DeathsPerOneMillion float64 `bson:"deathsPerOneMillion,omitempty"`
	Tests               int64   `bson:"tests,omitempty,omitempty"`
	TestsPerOneMillion  float64 `bson:"testsPerOneMillion,omitempty"`
	PopulationDensity   int64   `bson:"populationDensity"`
	FortnightCases      []int64 `bson:"fortnightCases"`
	FortnightAverage    float64	`bson:"fortnightAverage"`
	WeekAverage         float64 `bson:"weekAverage"`
	DropRate            float64	`bson:"dropRate"`
	OneWeekProjection   float64	  `bson:"oneWeekProjection"`
}

//CountryData ...
type CountryData struct {
	Country           string `json:"country"`
	PopulationDensity string `json:"popDensity"`
}

type HistoryData struct{
	Country string
	WeekData []float64
	fortnightData []float64
}

var Countries []map[string]interface{}

var Ctry []interface{}
