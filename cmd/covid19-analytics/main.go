package main

import (
  "time"
	"net/http"
	"os"

	"github.com/ObieWalker/covid19-analytics/helper"
	"github.com/ObieWalker/covid19-analytics/routes"
	"github.com/ObieWalker/covid19-analytics/services"
	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Initialize database
	helper.ConnectDB()

	nyc, _ := time.LoadLocation("America/New_York")
	c := cron.New(cron.WithLocation(nyc))
	c.AddFunc("15 18 * * ?", func() {
    log.Infof("Cron Job Running...")
		services.UpdateCountriesData()
	})

  c.Start()

	router := mux.NewRouter()
	routes.UseRoutes(router)

	http.ListenAndServe(":" + os.Getenv("PORT"), router)
}
