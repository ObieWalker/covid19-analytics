package main

import (
	"time"

	"context"

	"github.com/ObieWalker/covid19-analytics/helper"
	"github.com/ObieWalker/covid19-analytics/routes"
	"github.com/ObieWalker/covid19-analytics/services"
	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
)

func main() {

	nyc, _ := time.LoadLocation("America/New_York")
	c := cron.New(cron.WithLocation(nyc))
	c.AddFunc("0 19 * * ?", func() {
		log.Infof("Cron Job Running...")
		services.UpdateCountriesData()
	})

	c.Start()

	helper.ConnectDB()

	router := mux.NewRouter()
	routes.UseRoutes(router)

	http.ListenAndServe(":8000", router)

}
