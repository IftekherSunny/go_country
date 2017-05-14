### Country

Country is the package that helps you to get country name and dialing code by the country ISO 3166-1 Alpha-2 code.

### Installation Process

```
 go get github.com/IftekherSunny/country-for-golang
```

### Basic Uses

##### Get all countries name and dialing code

```go
country := country_for_golang.NewCountry()

countries := country.All()
```

##### Get a country name and dialing code

```go
country := country_for_golang.NewCountry()

countryDetails, _ := country.Get("BD")
```

##### Get multiple countries name and dialing code

```go
country := country_for_golang.NewCountry()

countriesDetails, _ := country.Get([]string{"BD", "US"})
```

##### Get a country name

```go
country := country_for_golang.NewCountry()

countryDetails, _ := country.GetName("BD")
```

##### Get a country dialing code

```go
country := country_for_golang.NewCountry()

countryDetails, _ := country.GetDialingCode("BD")
```

### Test

##### Run tests

```
go test -v
```

### License
This package is licensed under the [MIT License](https://github.com/iftekhersunny/country-for-golang/blob/master/LICENSE)
