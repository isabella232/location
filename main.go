package main

import (
	"fmt"
	"os"

	geoip2 "github.com/oschwald/geoip2-golang"
	"github.com/thoughtbot/location/locator"
	"github.com/thoughtbot/location/web"
)

func main() {
	db, err := geoip2.Open("data/GeoLite2-City.mmdb")
	if err != nil {
		fmt.Println("Unable to read GeoLite DB")
		os.Exit(1)
	}
	defer db.Close()

	offices, err := locator.OfficeRepo{}.LoadOffices("data/offices.yaml")
	if err != nil {
		fmt.Println("Unable to read offices YAML")
		os.Exit(1)
	}

	ol := locator.OfficeLocator{
		IPResolver: locator.IpResolver{DB: db},
		Offices:    offices,
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		fmt.Println("Please specify $PORT")
		os.Exit(1)
	}
	web.GetMainEngine(ol).Run(":" + port)
}
