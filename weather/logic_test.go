package weather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddType(t *testing.T) {
	assert := assert.New(t)
	testData := []struct {
		data            Data
		expectedWeather Type
	}{
		{Data{Code: "NMR"}, NormalWeather},
		{Data{Code: "nmr"}, NormalWeather},
		{Data{Code: "SRO"}, StormWeather},
		{Data{Code: "sro"}, StormWeather},
		{Data{Code: "HVA"}, RainWeather},
		{Data{Code: "hva"}, RainWeather},
		{Data{Code: "T E"}, DryWeather},
		{Data{Code: "t e"}, DryWeather},
		{Data{Code: "FUNDEFINEDG"}, FogWeather},
		{Data{Code: "fundefinedg"}, FogWeather},
		{Data{Code: "RANDOM STRING RETURNS NORMAL WEATHER!!!"}, NormalWeather}}

	for _, data := range testData {
		data.data.AddType()

		assert.Equal(data.expectedWeather, data.data.Type, data.data.Code)
	}
}
