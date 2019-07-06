package structure

import (
	"fmt"
)

type CountryInfo struct {
	IsoCode string `json:"isoCode"`
	Name    string `json:"Name"`
}

func (c *CountryInfo) String() string {
	return fmt.Sprintf("%+v", *c)
}
