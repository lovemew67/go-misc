package structure

import (
	"fmt"
)

type CityInfo struct {
	GeoNameId string `json:"geoNameId"`
	EnName    string `json:"enName"`
	Name      string `json:"name"`
}

func (c *CityInfo) String() string {
	return fmt.Sprintf("%+v", *c)
}

type CityInfoSet map[string]CityInfo

func NewCityInfoSet() *CityInfoSet {
	return &CityInfoSet{}
}

func (c *CityInfoSet) Len() int {
	return len(*c)
}

func (c *CityInfoSet) Add(city CityInfo) {
	(*c)[city.GeoNameId] = city
}

func (c *CityInfoSet) Contains(city CityInfo) bool {
	_, ok := (*c)[city.GeoNameId]
	return ok
}

func (c *CityInfoSet) Delete(city CityInfo) {
	delete(*c, city.GeoNameId)
}
