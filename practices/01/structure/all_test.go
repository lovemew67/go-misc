package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	cityInfo    = CityInfo{
		GeoNameId: "test",
	}
	countryInfo = CountryInfo{}
	subInfo     = SubInfo{
		IsoCode: "test",
	}
)

func TestInfo(t *testing.T) {
	assert.NotNil(t, cityInfo.String())
	assert.NotNil(t, countryInfo.String())
	assert.NotNil(t, subInfo.String())
}

func TestSet(t *testing.T) {
	testCitySet := NewCityInfoSet()
	testCitySet.Len()
	testCitySet.Add(cityInfo)
	testCitySet.Contains(cityInfo)
	testCitySet.Delete(cityInfo)

	testSubInfoSet := NewSubInfoSet()
	testSubInfoSet.Len()
	testSubInfoSet.Add(subInfo)
	testSubInfoSet.Contains(subInfo)
	testSubInfoSet.Delete(subInfo)
}