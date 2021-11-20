package bikes

import (
	"encoding/json"
	"net/http"
)

func AllBikes(w http.ResponseWriter, r *http.Request) {
	allBikes := []Bike{
		{
			Name:      "Endurance",
			Brand:     "Wilier",
			Frame:     "Wilier Cento1 NDR",
			Shifting:  "Campagnolo Chorus 12s",
			ModelYear: 2020,
		},{
			Name:      "Race",
			Brand:     "Wilier",
			Frame:     "Wilier Cento1 NDR",
			Shifting:  "Campagnolo Chorus 12s",
			ModelYear: 2020,
		},
	}
	json.NewEncoder(w).Encode(allBikes)
}
