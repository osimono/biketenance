package bikes

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/osimono/biketenance/cmd/db"
)

const (
	indexName    = "bikes"
	indexPattern = "bike:*"
)

var indexContext = db.IndexContext{
	Name:    indexName,
	Pattern: indexPattern,
}

func Save(bike *Bike) (err error) {
	if len(bike.Id) == 0 {
		uuid, _ := uuid.NewUUID()
		bike.Id = uuid.String()
	}

	return db.Save(indexContext, bike)
}

func ListAll() (allBikes []Bike, err error) {
	allItems, err := db.ListAll(indexContext)

	allBikes = make([]Bike, len(allItems))
	for i, item := range allItems {
		var bike Bike
		err := json.Unmarshal([]byte(item), &bike)
		if err != nil {

		}
		allBikes[i] = bike
	}
	return
}
