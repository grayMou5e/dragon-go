package dragon

import (
	"testing"

	"github.com/grayMou5e/dragon-go/weather"
	"github.com/stretchr/testify/assert"
)

func Test_CreateDragon(t *testing.T) {
	assert := assert.New(t)
	testingData := []struct {
		atack          int8
		armor          int8
		agility        int8
		endurance      int8
		weather        weather.Type
		expectedDragon Data
	}{
		{8, 7, 3, 2, weather.NormalWeather, Data{ScaleThickness: 10, ClawSharpness: 6, WingStrength: 2, FireBreath: 2, Scared: false}},
		{8, 7, 3, 2, weather.StormWeather, Data{ScaleThickness: 0, ClawSharpness: 0, WingStrength: 0, FireBreath: 0, Scared: true}},
		{8, 7, 3, 2, weather.DryWeather, Data{ScaleThickness: 5, ClawSharpness: 5, WingStrength: 5, FireBreath: 5, Scared: false}},
		{8, 7, 3, 2, weather.RainWeather, Data{ScaleThickness: 5, ClawSharpness: 10, WingStrength: 5, FireBreath: 0, Scared: false}},
		{8, 7, 3, 2, weather.FogWeather, Data{ScaleThickness: 5, ClawSharpness: 10, WingStrength: 1, FireBreath: 4, Scared: false}},
	}

	for _, data := range testingData {
		resultDragon := CreateDragon(data.atack, data.armor, data.agility, data.endurance, data.weather)

		assert.Equal(data.expectedDragon.ClawSharpness, resultDragon.ClawSharpness)
		assert.Equal(data.expectedDragon.ScaleThickness, resultDragon.ScaleThickness)
		assert.Equal(data.expectedDragon.FireBreath, resultDragon.FireBreath)
		assert.Equal(data.expectedDragon.WingStrength, resultDragon.WingStrength)
		assert.Equal(data.expectedDragon.Scared, resultDragon.Scared)
	}
}

func Test_GetNormalDragon(t *testing.T) {
	assert := assert.New(t)
	testingData := []struct {
		atack          int8
		armor          int8
		agility        int8
		endurance      int8
		expectedDragon Data
	}{
		{8, 7, 3, 2, Data{ScaleThickness: 10, ClawSharpness: 6, WingStrength: 2, FireBreath: 2}},
		{7, 8, 6, 5, Data{ScaleThickness: 6, ClawSharpness: 10, WingStrength: 5, FireBreath: 5}},
		{3, 6, 7, 4, Data{ScaleThickness: 3, ClawSharpness: 5, WingStrength: 9, FireBreath: 3}},
	}

	for _, data := range testingData {
		resultDragon := getNormalDragon(data.atack, data.armor, data.agility, data.endurance)

		assert.Equal(data.expectedDragon.ClawSharpness, resultDragon.ClawSharpness)
		assert.Equal(data.expectedDragon.ScaleThickness, resultDragon.ScaleThickness)
		assert.Equal(data.expectedDragon.FireBreath, resultDragon.FireBreath)
		assert.Equal(data.expectedDragon.WingStrength, resultDragon.WingStrength)
		assert.Equal(data.expectedDragon.Scared, resultDragon.Scared)
	}
}
