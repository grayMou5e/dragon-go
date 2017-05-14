package dragon

//Dragon model struct for storing dragon information
type Dragon struct {
	ScaleThickness int8 `json:"scaleThickness"`
	ClawSharpness  int8 `json:"clawSharpness"`
	WingStrength   int8 `json:"wingStrength"`
	FireBreath     int8 `json:"fireBreath"`
}
