package command

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/lovemew67/go-misc/01/structure"
	"github.com/lovemew67/go-misc/01/util"
)

var (
	jsonCountriesFile string
	jsonSubdivisionsFile string
	jsonCitiesFile string
	cdnCountriesFile string
	cdnFolder string
	jsonFolder string
	locale string

	geoIpLocaion = "./file/GeoLite2-City-Locations-en.csv"

	cdnSubdivisionsFolderTemplate = "%s/%s"
	cdnSubdivisionsFileTemplate   = "%s/%s/subdivisions-%s.json"
	cdnCitiesFolderTemplate       = "%s/%s/%s"
	cdnCitiesFileTemplate         = "%s/%s/%s/cities-%s.json"
)

func NewParseCommand() *cobra.Command {
	var (
		root string
		locale string
		command string
		location string
		indent = false
	)

	var parseCmd = &cobra.Command{
		Use:   "parse",
		Short: "parse geo ip list to cdn usage",
		Long:  `parse geo ip list to cdn usage`,
		Run: func(cmd *cobra.Command, args []string) {

			geoIpLocaion = location
			cdnFolder  = root+"/docs/cdn"
			jsonFolder = root+"/docs/json"

			jsonCountriesFile    = fmt.Sprintf("%s/countries-%s.json", jsonFolder, locale)
			jsonSubdivisionsFile = fmt.Sprintf("%s/subdivisions-%s.json", jsonFolder, locale)
			jsonCitiesFile       = fmt.Sprintf("%s/cities-%s.json", jsonFolder, locale)
			cdnCountriesFile     = fmt.Sprintf("%s/countries-%s.json", cdnFolder, locale)

			switch command {
			case "generate":
				log.Println("received command: generate")
				generate(indent)
			case "validate":
				log.Println("received command: validate")
				validate()
			default:
				log.Println("unknown command: ", command)
			}
		},
	}
	parseCmd.Flags().BoolVarP(&indent, "indent", "i", false, "to indent json file or not.")
	parseCmd.Flags().StringVarP(&root, "root", "r", ".", "file folder path.")
	parseCmd.Flags().StringVarP(&command, "locale", "a", "en", "to indicate the locale.")
	parseCmd.Flags().StringVarP(&command, "command", "c", "", "to indicate what to do.")
	parseCmd.Flags().StringVarP(&location, "location", "b", geoIpLocaion, "to indicate where is the source file.")
	return parseCmd
}

func generate(indent bool) {
	// read file
	csvFile, err := os.Open(geoIpLocaion)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	// skip title
	csvReader := csv.NewReader(csvFile)
	_, err = csvReader.Read()
	if err != nil {
		panic(err)
	}

	// init ds
	countryMap := map[string]structure.CountryInfo{}
	subMap := map[string]structure.SubInfoSet{}
	cityMap := map[string]map[string]structure.CityInfoSet{}

	// rows is of type [][]string
	rows, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, row := range rows {

		// geoname_id,             row[0]
		// country_iso_code,       row[4]
		// country_name,           row[5]
		// subdivision_1_iso_code, row[6]
		// subdivision_1_name,     row[7]
		// subdivision_2_iso_code, row[8]
		// subdivision_2_name,     row[9]
		// city_name,              row[10]

		countryIso := row[4]
		countryName := row[5]
		subdivisionIso := row[6]
		subdivisionName := row[7]
		geonameID := row[0]
		cityName := row[10]

		if len(countryName) != 0 && len(cityName) != 0 {

			// country map
			if _, ok := countryMap[countryIso]; !ok {
				countryMap[countryIso] = structure.CountryInfo{
					IsoCode: countryIso,
					Name:    countryName,
				}
			}

			// sub map
			if len(subdivisionIso) == 0 {
				subdivisionIso = row[4]
				subdivisionName = row[5]
			}
			subValue, ok := subMap[countryIso]
			if !ok {
				newSubSet := structure.NewSubInfoSet()
				newSubInfo := structure.SubInfo{
					IsoCode: subdivisionIso,
					Name:    subdivisionName,
				}
				newSubSet.Add(newSubInfo)
				subMap[countryIso] = *newSubSet
			} else {
				newSubInfo := structure.SubInfo{
					IsoCode: subdivisionIso,
					Name:    subdivisionName,
				}
				if !subValue.Contains(newSubInfo) {
					subValue.Add(newSubInfo)
					subMap[countryIso] = subValue
				}
			}

			// city map
			cityValue, ok := cityMap[countryIso]
			newCityInfo := structure.CityInfo{
				GeoNameId: geonameID,
				EnName:    cityName,
				Name:      cityName,
			}
			if !ok {
				tempCitySet := structure.NewCityInfoSet()
				tempCitySet.Add(newCityInfo)
				cityMap[countryIso] = map[string]structure.CityInfoSet{}
				cityMap[countryIso][subdivisionIso] = *tempCitySet
			} else {
				subValue, ok := cityValue[subdivisionIso]
				if !ok {
					tempCitySet := structure.NewCityInfoSet()
					tempCitySet.Add(newCityInfo)
					cityValue[subdivisionIso] = *tempCitySet
				} else {
					subValue.Add(newCityInfo)
					cityValue[subdivisionIso] = subValue
				}
			}
		}
	}

	// output map to json
	util.CreateDirIfNotExist(jsonFolder)
	util.MarshalJsonAndWriteFile(countryMap, jsonCountriesFile, indent)
	util.MarshalJsonAndWriteFile(subMap, jsonSubdivisionsFile, indent)
	util.MarshalJsonAndWriteFile(cityMap, jsonCitiesFile, indent)

	// convert to array and output to cdn
	// country
	util.CreateDirIfNotExist(cdnFolder)
	countryArray := []structure.CountryInfo{}
	for _, countryInfo := range countryMap {
		countryArray = append(countryArray, countryInfo)
	}
	util.MarshalJsonAndWriteFile(countryArray, cdnCountriesFile, indent)

	// subdivision
	for countryIso, countrySubSet := range subMap {
		util.CreateDirIfNotExist(fmt.Sprintf(cdnSubdivisionsFolderTemplate, cdnFolder, countryIso))
		subArray := []structure.SubInfo{}
		for _, sub := range countrySubSet {
			subArray = append(subArray, sub)
		}
		util.MarshalJsonAndWriteFile(subArray, fmt.Sprintf(cdnSubdivisionsFileTemplate, cdnFolder, countryIso, locale), indent)
	}
	// city
	for countryIso, countrySubMap := range cityMap {
		for subIso, countrySubCitySet := range countrySubMap {
			util.CreateDirIfNotExist(fmt.Sprintf(cdnCitiesFolderTemplate, cdnFolder, countryIso, subIso))
			cityArray := []structure.CityInfo{}
			for _, city := range countrySubCitySet {
				cityArray = append(cityArray, city)
			}
			util.MarshalJsonAndWriteFile(cityArray, fmt.Sprintf(cdnCitiesFileTemplate, cdnFolder, countryIso, subIso, locale), indent)
		}
	}
}

func validate() {
	if _, errExist := os.Stat(cdnFolder); os.IsNotExist(errExist) {
		panic(errExist)
	}
	countryArray := []structure.CountryInfo{}
	countriesFile, errReadFile := ioutil.ReadFile(cdnCountriesFile)
	if errReadFile != nil {
		panic(errReadFile)
	}
	errJson := json.Unmarshal([]byte(countriesFile), &countryArray)
	if errJson != nil {
		panic(errJson)
	}
	folderNumber := 0
	countryFolders, errReadDir := ioutil.ReadDir(cdnFolder)
	if errReadDir != nil {
		panic(errReadDir)
	}
	for _, f := range countryFolders {
		if f.IsDir() {
			folderNumber++
		}
	}
	if len(countryArray) != folderNumber {
		panic(fmt.Sprintf("no aligned folder number between: %s and numbers in: %s", cdnFolder, cdnCountriesFile))
	}
	for _, currentCountry := range countryArray {
		if len(currentCountry.IsoCode) == 0 || len(currentCountry.Name) == 0 {
			panic(fmt.Sprintf("invalid cdn countries file: %s", cdnCountriesFile))
		}
		subFolderName := fmt.Sprintf(cdnSubdivisionsFolderTemplate, cdnFolder, currentCountry.IsoCode)
		subFolderFile := fmt.Sprintf(cdnSubdivisionsFileTemplate, cdnFolder, currentCountry.IsoCode, locale)
		if _, errExist := os.Stat(subFolderName); os.IsNotExist(errExist) {
			panic(errExist)
		}
		subArray := []structure.SubInfo{}
		subdivisionFile, errReadFile := ioutil.ReadFile(subFolderFile)
		if errReadFile != nil {
			panic(errReadFile)
		}
		errJson := json.Unmarshal([]byte(subdivisionFile), &subArray)
		if errJson != nil {
			panic(errJson)
		}
		folderNumber = 0
		subFolders, errReadDir := ioutil.ReadDir(subFolderName)
		if errReadDir != nil {
			panic(errReadDir)
		}
		for _, f := range subFolders {
			if f.IsDir() {
				folderNumber++
			}
		}
		if len(subArray) != folderNumber {
			panic(fmt.Sprintf("no aligned folder number between: %s and numbers in: %s", subFolderName, subFolderFile))
		}
		for _, currentSub := range subArray {
			if len(currentSub.IsoCode) == 0 || len(currentSub.Name) == 0 {
				panic(fmt.Sprintf("invalid subdivision file: %s", subFolderFile))
			}
			cityFolderName := fmt.Sprintf(cdnCitiesFolderTemplate, cdnFolder, currentCountry.IsoCode, currentSub.IsoCode)
			cityFolderFile := fmt.Sprintf(cdnCitiesFileTemplate, cdnFolder, currentCountry.IsoCode, currentSub.IsoCode, locale)
			if _, errExist := os.Stat(cityFolderName); os.IsNotExist(errExist) {
				panic(errExist)
			}
			cityArray := []structure.CityInfo{}
			cityFile, errReadFile := ioutil.ReadFile(cityFolderFile)
			if errReadFile != nil {
				panic(errReadFile)
			}
			errJson := json.Unmarshal([]byte(cityFile), &cityArray)
			if errJson != nil {
				panic(errJson)
			}
			for _, currentCity := range cityArray {
				if len(currentCity.GeoNameId) == 0 || len(currentCity.EnName) == 0 || len(currentCity.Name) == 0 {
					panic(fmt.Sprintf("invalid city file: %s", cityFolderFile))
				}
			}
			log.Println("verified file: ", cityFolderFile)
		}
		log.Println("verified file: ", subFolderFile)
	}
	log.Println("verified file: ", cdnCountriesFile)
}
