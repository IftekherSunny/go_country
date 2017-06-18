package country

import "testing"

func Test_it_should_return_all_countries(t *testing.T) {

	// setup
	country := NewCountry()

	data := country.All()

	// asset
	assert(t, 245, len(data), "Total countries data should be 245")
}

func Test_it_should_return_a_country_name_and_dialing_code(t *testing.T) {

	// setup
	country := NewCountry()

	data, _ := country.Get("BD")

	details := data.(map[string]interface{})

	// asset
	assert(t, 2, len(details), "Length should be 2")

	assert(t, "Bangladesh", details["name"], "Country name should be Bangladesh")

	assert(t, "+880", details["code"], "Country dialing code should be +880")
}

func Test_it_should_return_multiple_countries_name_and_dialing_code(t *testing.T) {

	// setup
	country := NewCountry()

	data, _ := country.Get([]string{"BD", "US"})

	details := data.(map[string]interface{})

	BD := details["BD"].(map[string]interface{})
	US := details["US"].(map[string]interface{})

	// asset
	assert(t, 2, len(details), "Length should be 2")

	assert(t, "Bangladesh", BD["name"], "Country name should be Bangladesh")

	assert(t, "+880", BD["code"], "Country dialing code should be +880")

	assert(t, "United States", US["name"], "Country name should be United States")

	assert(t, "+1", US["code"], "Country dialing code should be +1")
}

func Test_it_should_return_country_code_validation_error_message(t *testing.T) {

	// setup
	country := NewCountry()

	_, err := country.Get("unknown")

	// asset
	assert(t, "The country ISO 3166-1 Alpha-2 code [ unknown ] does not exists.", err.Error(), "Country code should be unknown")
}

func Test_it_should_return_country_code_validation_error_message_when_trying_to_get_multiple_countries_name_and_dialing_code(t *testing.T) {

	// setup
	country := NewCountry()

	_, err := country.Get([]string{"BD", "unknown"})

	// asset
	assert(t, "The country ISO 3166-1 Alpha-2 code [ unknown ] does not exists.", err.Error(), "Country code should be unknown")
}

func Test_it_should_return_a_country_name(t *testing.T) {

	// setup
	country := NewCountry()

	countryName, _ := country.GetName("BD")

	// asset
	assert(t, "Bangladesh", countryName, "Country name should be Bangladesh")
}

func Test_it_should_return_a_country_dialing_code(t *testing.T) {

	// setup
	country := NewCountry()

	countryName, _ := country.GetDialingCode("BD")

	// asset
	assert(t, "+880", countryName, "Country dialing code should be +880")
}
