package result

import "strings"

//Summarize to look if fight was won or lost
func (d *Data) Summarize() {
	d.Victory = strings.ToLower(d.Status) == "victory"
}
