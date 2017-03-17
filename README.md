# Location

[![Build Status](https://travis-ci.com/thoughtbot/location.svg?token=47cp3CiWHmDqjYJKGejt&branch=master)](https://travis-ci.com/thoughtbot/location)

Live service: [https://location.thoughtbot.com/v1/nearest](https://location.thoughtbot.com/v1/nearest)

## Endpoints

### `/v1/nearest`

Returns information about the nearest thoughtbot office based on your IP:

```json
{
  "meta": {
    "distanceKmToUser": 1.1977565711120122
  },
  "name": "London",
  "slug": "london"
}
```

## Developing

- You need to have [Go](https://golang.org/) 1.7+ setup on your machine
- Set your `$GOPATH` to a value such as: `~/workspace/go`
- Clone the repo within: `$GOPATH/src/github.com/thoughtbot/location`
- Run `bin/setup` to install dependencies
- Run `bin/test` to run tests
- Run `bin/run` to run the service locally

### Resources

Finding the nearest city to the user (based on the user's IP address) uses the
free [GeoLite2 City database](http://dev.maxmind.com/geoip/geoip2/geolite2/).
This data is updated monthly.

The data does not include street-address level IP mappings, the mapping is only
to the nearest city, which is adequate to locate the nearest office.
