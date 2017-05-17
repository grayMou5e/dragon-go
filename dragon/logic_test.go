package dragon

import (
	"testing"

	"github.com/grayMou5e/dragon-go/weather"
)

func Test_CreateDragon(t *testing.T) {
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

		if resultDragon.ClawSharpness != data.expectedDragon.ClawSharpness {
			t.Errorf("Expected Claw Sharpness of %d but received %d", data.expectedDragon.ClawSharpness, resultDragon.ClawSharpness)
		}
		if resultDragon.ScaleThickness != data.expectedDragon.ScaleThickness {
			t.Errorf("Expected Scale Thickness of %d but received %d", data.expectedDragon.ScaleThickness, resultDragon.ScaleThickness)
		}
		if resultDragon.FireBreath != data.expectedDragon.FireBreath {
			t.Errorf("Expected Fire Breath of %d but received %d", data.expectedDragon.FireBreath, resultDragon.FireBreath)
		}
		if resultDragon.WingStrength != data.expectedDragon.WingStrength {
			t.Errorf("Expected Wing Strength of %d but received %d", data.expectedDragon.WingStrength, resultDragon.WingStrength)
		}
		if resultDragon.Scared != data.expectedDragon.Scared {
			t.Errorf("Expected Scared value of %t but received %t", data.expectedDragon.Scared, resultDragon.Scared)
		}
	}
}

func Test_GetNormalDragon(t *testing.T) {
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

		if resultDragon.ClawSharpness != data.expectedDragon.ClawSharpness {
			t.Errorf("Expected Claw Sharpness of %d but received %d", data.expectedDragon.ClawSharpness, resultDragon.ClawSharpness)
		}
		if resultDragon.ScaleThickness != data.expectedDragon.ScaleThickness {
			t.Errorf("Expected Scale Thickness of %d but received %d", data.expectedDragon.ScaleThickness, resultDragon.ScaleThickness)
		}
		if resultDragon.FireBreath != data.expectedDragon.FireBreath {
			t.Errorf("Expected Fire Breath of %d but received %d", data.expectedDragon.FireBreath, resultDragon.FireBreath)
		}
		if resultDragon.WingStrength != data.expectedDragon.WingStrength {
			t.Errorf("Expected Wing Strength of %d but received %d", data.expectedDragon.WingStrength, resultDragon.WingStrength)
		}
	}
}
