package dragon

//Data model struct for storing dragon information
type Data struct {
	ScaleThickness int8 `json:"scaleThickness"`
	ClawSharpness  int8 `json:"clawSharpness"`
	WingStrength   int8 `json:"wingStrength"`
	FireBreath     int8 `json:"fireBreath"`
	//Scared indicates i;f dragon can go to fight or not
	Scared bool
}
