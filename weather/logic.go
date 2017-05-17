package weather

import "strings"

//AddType adds weather type to existing object
func AddType(weather *Data) {
	toCompare := strings.ToLower(weather.Message)
	if strings.Contains(toCompare, "regular normal weather") {
		weather.Type = NormalWeather
	} else if strings.Contains(toCompare, "frog rain") {
		weather.Type = StormWeather
	} else if strings.Contains(toCompare, "flood") {
		weather.Type = RainWeather
	} else if strings.Contains(toCompare, "long dry") {
		weather.Type = DryWeather
	} else if strings.Contains(toCompare, "fog") {
		weather.Type = FogWeather
	}
}
