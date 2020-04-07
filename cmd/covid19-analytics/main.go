package main

import (
  "github.com/robfig/cron/v3"
  "github.com/ObieWalker/covid19-analytics/services"
  log "github.com/sirupsen/logrus"
)

func main() {

  c := cron.New()
  c.AddFunc("@every 12h", func() {
    log.Infof("Cron Job Running...")
    services.GetCountriesData() 
  })

  c.Start()

 }