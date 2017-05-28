package result

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummarize(t *testing.T) {
	assert := assert.New(t)
	testData := []struct {
		data   Data
		result bool
	}{
		{Data{Status: "victory"}, true},
		{Data{Status: "Victory"}, true},
		{Data{Status: "VICTORY"}, true},
		{Data{Status: "Lost"}, false},
		{Data{Status: "Defeat"}, false},
		{Data{Status: "GIBERISH"}, false},
	}
	for _, data := range testData {
		data.data.Summarize()
		assert.Equal(data.result, data.data.Victory)
	}

}
