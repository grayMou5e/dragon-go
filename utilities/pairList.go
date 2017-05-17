package utilities

//Pair is a data struction for holding up key & value
type Pair struct {
	Key   string
	Value int8
}

//PairList is an array of Pair structs
type PairList []Pair

//ToMap converts PairList into map
func (p PairList) ToMap() map[string]int8 {
	m := make(map[string]int8, 5)
	for _, element := range p {
		m[element.Key] = element.Value
	}
	return m
}

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
