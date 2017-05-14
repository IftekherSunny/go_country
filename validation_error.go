package country_for_golang

// validation error struct
type ValidationError struct {
	message string
}

// implementing error method
func (v ValidationError) Error() string {
	return v.message
}

// create a new validation error class
func NewValidationError(countryName string) error {
	ve := &ValidationError{}
	ve.message = "Country name [ " + countryName + " ] doesn't exist"

	return ve
}
