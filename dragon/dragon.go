package dragon

import (
	"fmt"
	"sort"
)

//Dragon model struct for storing dragon information
type Dragon struct {
	ScaleThickness int8 `json:"scaleThickness"`
	ClawSharpness  int8 `json:"clawSharpness"`
	WingStrength   int8 `json:"wingStrength"`
	FireBreath     int8 `json:"fireBreath"`
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
	knightEndurance int8) *Dragon {
	powers := pairList{
		pair{powerAttack, knightAttack},
		pair{powerArmor, knightArmor},
		pair{powerAgility, knightAgility},
		pair{powerEndurance, knightEndurance},
	}
	sort.Sort(sort.Reverse(powers))

	strongestStat := powers[0]
	strongestStatDif := int(10 - strongestStat.Value)

	fmt.Println(powers)

	for i := 1; i <= strongestStatDif && i < powers.Len()-1; i++ {
		powers[0].Value++
		powers[i].Value--
	}

	dragonPower := powers.ToMap()

	return &Dragon{
		ClawSharpness:  dragonPower[powerArmor],
		ScaleThickness: dragonPower[powerAttack],
		WingStrength:   dragonPower[powerAgility],
		FireBreath:     dragonPower[powerEndurance],
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
