package result

//Data for storing end information about the fight
type Data struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Victory bool   `json:"-"`
}
