package weather

//Type for assinging type for current weather object
type weatherType uint8

const (
	//NormalWeather indicates normal weather
	NormalWeather weatherType = iota
	//DryWeather indicates that it's dry
	DryWeather weatherType = iota
	//FogWeather indicates that there is fog
	FogWeather weatherType = iota
	//StormWeather indicates that there is stroms
	StormWeather weatherType = iota
	//RainWeather indicates taht there is raining
	RainWeather weatherType = iota
)

//Weather struct is dedicated for holding up the data for weather
type Weather struct {
	Message string `xml:"message"`
	Type    weatherType
}

// //AddType adds weather type to existing object
// func AddType(weather *Weather) {
// 	if strings.Contains(weather.Message,""
// }
