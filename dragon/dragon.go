package dragon

import (
	"sort"

	"github.com/grayMou5e/dragon-go/weather"
)

//Data model struct for storing dragon information
type Data struct {
	ScaleThickness int8 `json:"scaleThickness"`
	ClawSharpness  int8 `json:"clawSharpness"`
	WingStrength   int8 `json:"wingStrength"`
	FireBreath     int8 `json:"fireBreath"`
	//Scared indicates if dragon can go to fight or not
	Scared bool
}

const (
	powerAttack    = "attack"
	powerArmor     = "armor"
	powerAgility   = "agility"
	powerEndurance = "endurance"
)

//CreateDragon function for creating dragon based on knight stats
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
		return getNormalsDragon(knightAttack, knightArmor, knightAgility, knightEndurance)
	}
}

func getNormalsDragon(knightAttack int8,
	knightArmor int8,
	knightAgility int8,
	knightEndurance int8) *Data {
	powers := pairList{
		pair{powerAttack, knightAttack},
		pair{powerArmor, knightArmor},
		pair{powerAgility, knightAgility},
		pair{powerEndurance, knightEndurance},
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

type pair struct {
	Key   string
	Value int8
}

type pairList []pair

func (p pairList) ToMap() map[string]int8 {
	m := map[string]int8{
		powerAttack:    0,
		powerArmor:     0,
		powerAgility:   0,
		powerEndurance: 0,
	}
	for _, element := range p {
		m[element.Key] = element.Value
	}
	return m
}
func (p pairList) Len() int           { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
