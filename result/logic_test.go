package result

import (
	"testing"
)

func TestSummarize(t *testing.T) {
	testData := []struct {
		data   Data
		result bool
	}{
		{Data{Status: "victory"}, true},
		{Data{Status: "Victory"}, true},
		{Data{Status: "VICTORY"}, true},
	}
	for _, data := range testData {
		data.data.Summarize()
		if data.data.Victory != data.result {
			t.Errorf("Received [%s] and expected result %t but received %t", data.data.Status, data.result, false)
		}
	}

}
