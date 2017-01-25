package locator

import (
	"fmt"
	"net"

	geoip2 "github.com/oschwald/geoip2-golang"
)

type IpResolver struct {
	DB ipCityDBInterface
}

//go:generate counterfeiter . ipCityDBInterface
type ipCityDBInterface interface {
	City(net.IP) (*geoip2.City, error)
}

func (r IpResolver) ResolveCity(ip string) (lat, long float64, err error) {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return 0.0, 0.0, fmt.Errorf("Unable to parse IP")
	}

	record, _ := r.DB.City(parsedIP)
	if record == nil {
		return 0.0, 0.0, fmt.Errorf("Unable to resolve IP")
	}

	return record.Location.Latitude, record.Location.Longitude, nil
}
