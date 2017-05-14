### Country

Country is the package that helps you to get country name and dialing code by the country ISO 3166-1 Alpha-2 code.

### Installation Process

```
 go get github.com/IftekherSunny/country-for-golang
```

### Basic Uses

##### Get all countries name and dialing code

```go
country := go_country.NewCountry()

countries := country.All()
```

##### Get a country name and dialing code

```go
country := go_country.NewCountry()

countryDetails, _ := country.Get("BD")
```

##### Get multiple countries name and dialing code

```go
country := go_country.NewCountry()

countries, _ := country.Get([]string{"BD", "US"})
```

##### Get a country name

```go
country := go_country.NewCountry()

name, _ := country.GetName("BD")
```

##### Get a country dialing code

```go
country := go_country.NewCountry()

dialingCode, _ := country.GetDialingCode("BD")
```

### Test

##### Run tests

```
go test -v
```

### License
This package is licensed under the [MIT License](https://github.com/iftekhersunny/country-for-golang/blob/master/LICENSE)
