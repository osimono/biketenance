package bikes

import (
	"github.com/osimono/biketenance/cmd/db"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestDB(t *testing.T) {
	db.UseDb("./test-db.db")
	bike := &Bike{
		Id:        "d0c8292c-4a23-11ec-906e-704d7b669cd9",
		Name:      "Bikomat",
		Brand:     "Scott",
		Frame:     "Contessa",
		Shifting:  "Shimano",
		ModelYear: 0,
	}
	Save(bike)


	all, err := ListAll()
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("%#v", all)
}
