package weather

import "strings"

//Type for assinging type for current weather object
type Type uint8

const (
	//NormalWeather indicates normal weather
	NormalWeather Type = 1 << iota
	//DryWeather indicates that it's dry
	DryWeather Type = 1 << iota
	//FogWeather indicates that there is fog
	FogWeather Type = 1 << iota
	//StormWeather indicates that there is stroms
	StormWeather Type = 1 << iota
	//RainWeather indicates taht there is raining
	RainWeather Type = 1 << iota
)

//Data struct is dedicated for holding up the data for weather
type Data struct {
	Message string `xml:"message"`
	Type    Type
}

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
