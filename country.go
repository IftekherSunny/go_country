package country_for_golang

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"runtime"
)

// Country struct
type Country struct {
	data map[string]interface{}
}

// Create a new instance of country-for-golang struct
func New() *Country {
	instance := &Country{}

	instance = instance.readCountriesDataFile()

	return instance
}

// Read countries data file
func (country *Country) readCountriesDataFile() *Country {

	if len(country.data) > 0 {
		return country
	}

	_, filename, _, _ := runtime.Caller(0)

	dataPath := path.Join(path.Dir(filename), "data/countries.json")

	file, _ := ioutil.ReadFile(dataPath)

	json.Unmarshal([]byte(file), &country.data)

	return country
}

// Get all countries name and dialing code
func (country *Country) All() map[string]interface{} {
	return country.readCountriesDataFile().data
}

// Get a single country-for-golang by the given country-for-golang name
func (country *Country) getCountry(name string) (interface{}, error) {
	if details := country.readCountriesDataFile().data[name]; details == nil {
		return nil, NewValidationError(name)
	} else {
		return details, nil
	}
}

// Get a country-for-golang name and dialing code by the given country-for-golang name
func (country *Country) Get(name interface{}) (interface{}, error) {

	switch name.(type) {
	case string:
		return country.getCountry(name.(string))
	case []string:
		data, err := country.gets(name.([]string))

		return data, err
	default:
		return nil, nil
	}
}

// Get a country-for-golang name by the given country-for-golang name
func (country *Country) GetName(name string) (interface{}, error) {
	data, err := country.getCountry(name)

	if err != nil {
		return nil, err
	}

	details, _ := data.(map[string]interface{})

	return details["name"], nil
}

// Get a country-for-golang dialing code by the given country-for-golang name
func (country *Country) GetDialingCode(name string) (interface{}, error) {
	data, err := country.getCountry(name)

	if err != nil {
		return nil, err
	}

	details, _ := data.(map[string]interface{})

	return details["code"], nil
}

// Get countries name and dialing code by the given countries name
func (country *Country) gets(names []string) (map[string]interface{}, error) {
	countries := make(map[string]interface{})

	for _, name := range names {
		details, err := country.getCountry(name)

		if err != nil {
			return nil, err
		} else {
			countries[name] = details
		}
	}

	return countries, nil
}
