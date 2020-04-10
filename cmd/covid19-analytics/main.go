package main

import (
  "time"
  "github.com/robfig/cron/v3"
  "github.com/ObieWalker/covid19-analytics/services"
  log "github.com/sirupsen/logrus"
)

func main() {

  // nyc, _ := time.LoadLocation("America/New_York")
  // c := cron.New(cron.WithLocation(nyc))
  // c.AddFunc("0 19 * * ?", func() {
  c := cron.New()
    c.AddFunc("@every 30s", func() {
      log.Infof("Cron Job Running...")
    services.UpdateCountriesData() 
  })

  c.Start()

  time.Sleep(1 * time.Minute)
}
