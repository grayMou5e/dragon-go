package dragon

import (
	"sort"

	"github.com/grayMou5e/dragon-go/utilities"
	"github.com/grayMou5e/dragon-go/weather"
)

const (
	powerAttack    = "attack"
	powerArmor     = "armor"
	powerAgility   = "agility"
	powerEndurance = "endurance"
)

//CreateDragon function for creating dragon based on knight stats & weather type
func CreateDragon(knightAttack int8,
	knightArmor int8,
	knightAgility int8,
	knightEndurance int8,
	weatherType weather.Type) *Data {
	switch weatherType {
	case weather.DryWeather:
		return &Data{ClawSharpness: 5, FireBreath: 5, ScaleThickness: 5, WingStrength: 5, Scared: false}
	case weather.StormWeather:
		return &Data{ClawSharpness: 0, FireBreath: 0, ScaleThickness: 0, WingStrength: 0, Scared: true}
	case weather.RainWeather:
		return &Data{ClawSharpness: 10, WingStrength: 5, ScaleThickness: 5, FireBreath: 0, Scared: false}
	case weather.FogWeather:
		return &Data{ClawSharpness: 10, WingStrength: 1, ScaleThickness: 5, FireBreath: 4, Scared: false}
	default:
		return getNormalDragon(knightAttack, knightArmor, knightAgility, knightEndurance)
	}
}

func getNormalDragon(knightAttack int8,
	knightArmor int8,
	knightAgility int8,
	knightEndurance int8) *Data {
	powers := utilities.PairList{
		utilities.Pair{Key: powerAttack, Value: knightAttack},
		utilities.Pair{Key: powerArmor, Value: knightArmor},
		utilities.Pair{Key: powerAgility, Value: knightAgility},
		utilities.Pair{Key: powerEndurance, Value: knightEndurance},
	}
	sort.Sort(sort.Reverse(powers))

	strongestStat := powers[0]
	strongestStatDif := int(10 - strongestStat.Value)

	for i := 1; i <= strongestStatDif && i < powers.Len()-1; i++ {
		powers[0].Value++
		powers[i].Value--
	}

	dragonPower := powers.ToMap()

	return &Data{
		ClawSharpness:  dragonPower[powerArmor],
		ScaleThickness: dragonPower[powerAttack],
		WingStrength:   dragonPower[powerAgility],
		FireBreath:     dragonPower[powerEndurance],
		Scared:         false,
	}
}
