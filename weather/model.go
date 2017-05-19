package weather

//Type for assinging type for current weather object
type Type uint8

//Data struct is dedicated for holding up the data for weather
type Data struct {
	Message string `xml:"message"`
	Code    string `xml:"code"`
	Type    Type   `xml:"-"`
}

const (
	//NormalWeather indicates normal weather
	NormalWeather Type = iota
	//DryWeather indicates that it's dry
	DryWeather Type = iota
	//FogWeather indicates that there is fog
	FogWeather Type = iota
	//StormWeather indicates that there is stroms
	StormWeather Type = iota
	//RainWeather indicates taht there is raining
	RainWeather Type = iota
)
