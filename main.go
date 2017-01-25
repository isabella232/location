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
		fmt.Printf("Unable to read GeoLite DB")
		os.Exit(1)
	}
	defer db.Close()

	offices, err := locator.OfficeRepo{}.LoadOffices("data/offices.yaml")
	if err != nil {
		fmt.Printf("Unable to read offices YAML")
		os.Exit(1)
	}

	ol := locator.OfficeLocator{
		IPResolver: locator.IpResolver{DB: db},
		Offices:    offices,
	}
	web.GetMainEngine(ol).Run(":8080")
}
