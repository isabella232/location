package locator

import "math"

type OfficeLocator struct {
	IPResolver ipResolverInterface
	Offices    []Office
}

//go:generate counterfeiter . ipResolverInterface
type ipResolverInterface interface {
	ResolveCity(ip string) (lat, long float64, err error)
}

func (l OfficeLocator) Nearest(ipAddress string) (string, float64, error) {
	userLat, userLong, err := l.IPResolver.ResolveCity(ipAddress)
	if err != nil {
		return "", 0.0, err
	}

	nearestDistance := math.MaxFloat64
	var nearestOffice Office

	for _, o := range l.Offices {
		dist := calcDistance(userLat, userLong, o.Lat, o.Long)
		if dist < nearestDistance {
			nearestDistance = dist
			nearestOffice = o
		}
	}
	return nearestOffice.Slug, nearestDistance, nil
}

// https://github.com/njj/go-haversine/blob/master/haversine.go
func calcDistance(startLat, startLong, endLat, endLong float64) float64 {
	km := float64(6371)

	dLat := toRadians(endLat - startLat)
	dLon := toRadians(endLong - startLong)

	lat1 := toRadians(startLat)
	lat2 := toRadians(endLat)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*
			math.Cos(lat1)*math.Cos(lat2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return km * c
}

func toRadians(num float64) float64 {
	return num * math.Pi / 180
}
