package weather

//Type for assinging type for current weather object
type Type uint8

//Data struct is dedicated for holding up the data for weather
type Data struct {
	Message string `xml:"message"`
	Type    Type
}

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
