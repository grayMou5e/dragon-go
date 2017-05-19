package weather

import (
	"strings"
)

//AddType adds weather type to existing object
func (weather *Data) AddType() {
	toCompare := strings.ToLower(weather.Code)
	if toCompare == "nmr" {
		weather.Type = NormalWeather
	} else if toCompare == "sro" {
		weather.Type = StormWeather
	} else if toCompare == "hva" {
		weather.Type = RainWeather
	} else if toCompare == "t e" {
		weather.Type = DryWeather
	} else if toCompare == "fundefinedg" {
		weather.Type = FogWeather
	}
}
