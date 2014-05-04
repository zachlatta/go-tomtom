package tomtom

import (
	"fmt"
	"net/url"
)

type GeocodeService struct {
	client *Client
}

type reverseGeoResponseWrapper struct {
	ReverseGeoResponse reverseGeoResponse `json:"reverseGeoResponse"`
}

type reverseGeoResponse struct {
	ReverseGeoResults []ReverseGeoResult `json:"reverseGeoResult"`
}

// ReverseGeoResult is a result as returned by the TomTom API.
type ReverseGeoResult struct {
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	HouseNumber      string  `json:"houseNumber"`
	Street           string  `json:"street"`
	City             string  `json:"city"`
	State            string  `json:"state"`
	Country          string  `json:"country"`
	CountryISO3      string  `json:"countryISO3"`
	FormattedAddress string  `json:"formattedAddress"`
}

func (s *GeocodeService) ReverseGeocode(lat, long float64) ([]ReverseGeoResult,
	error) {
	url, err := url.Parse("lbs/services/reverseGeocode/3/json")
	if err != nil {
		return nil, err
	}

	params := url.Query()
	params.Set("point", fmt.Sprintf("%f,%f", lat, long))
	url.RawQuery = params.Encode()

	req, err := s.client.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	respStruct := new(reverseGeoResponseWrapper)
	err = s.client.Do(req, &respStruct)
	if err != nil {
		return nil, err
	}

	return respStruct.ReverseGeoResponse.ReverseGeoResults, nil
}
