package result

import (
	"strings"
)

//Data for storing end information about the fight
type Data struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Victory bool
}

//Summarize to look if fight was won or lost
func (d *Data) Summarize() {
	d.Victory = strings.ToLower(d.Status) == "victory"
}
