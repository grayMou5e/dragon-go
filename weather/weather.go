package weather

import "strings"

//Type for assinging type for current weather object
type weatherType uint8

const (
	//NormalWeather indicates normal weather
	NormalWeather weatherType = 1 << iota
	//DryWeather indicates that it's dry
	DryWeather weatherType = 1 << iota
	//FogWeather indicates that there is fog
	FogWeather weatherType = 1 << iota
	//StormWeather indicates that there is stroms
	StormWeather weatherType = 1 << iota
	//RainWeather indicates taht there is raining
	RainWeather weatherType = 1 << iota
)

//Weather struct is dedicated for holding up the data for weather
type Weather struct {
	Message string `xml:"message"`
	Type    weatherType
}

//AddType adds weather type to existing object
func AddType(weather *Weather) {
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
