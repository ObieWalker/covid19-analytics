package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Country struct {
	ID     primitive.ObjectID  `bson:"_id,omitempty"`
	Country string			       `bson:"country,omitempty"`
	CountryInfo struct {
		ID int32			           `bson:"_id,omitempty"`
		Iso2 string			         `bson:"iso2,omitempty"`
		Iso3 string			         `bson:"iso3,omitempty"`
		Lat float64			         `bson:"lat,omitempty"`
		Long float64		         `bson:"long,omitempty"`
		Flag string			         `bson:"flag,omitempty"`
	}
	Updated int64			         `bson:"updated,omitempty"`
	Cases int32				         `bson:"cases,omitempty"`
	TodayCases int32		       `bson:"todayCases,omitempty"`
	Deaths int32			         `bson:"deaths,omitempty"`
	TodayDeaths int32	    	   `bson:"todayDeaths,omitempty"`
	Recovered int32			       `bson:"recovered,omitempty"`
	Active int32			         `bson:"active,omitempty"`
	Critical int32			       `bson:"critical,omitempty"`
	CasesPerOneMillion float32 `bson:"casesPerOneMillion,omitempty"`
	DeathsPerOneMillion float32`bson:"deathsPerOneMillion,omitempty"`
	Tests int32			 	         `bson:"tests,omitempty,omitempty"`
	TestsPerOneMillion float32 `bson:"testsPerOneMillion,omitempty"`
	PopulationDensity int32	   `bson:"populationDensity"`
	LatestWeekCases []int32  	 `bson:"latestWeekCases"`
	WeekIncreaseRate float32	 `bson:"weekIncreaseRate"`
	ThreeDayRate int32    		 `bson:"threeDayRate"`
}