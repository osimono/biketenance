package bikes

type Bike struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Brand     string `json:"brand"`
	Frame     string `json:"frame"`
	Shifting  string `json:"shifting"`
	ModelYear int    `json:"modelYear"`
}

func (b Bike) ItemId() string {
	return b.Id
}



