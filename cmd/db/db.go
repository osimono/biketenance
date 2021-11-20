package db

import (
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/buntdb"
	"sync"
)

var (
	dbFile = "data.db"
	once   sync.Once
	db *buntdb.DB
)

func UseDb(dbFileName string) {
	dbFile = dbFileName
}

func connect(dbFileName string) (*buntdb.DB, error) {
	db, err := buntdb.Open(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func DB() *buntdb.DB {
	once.Do(func() {
		var err error
		db, err = connect(dbFile)
		if err != nil {
			log.Fatalf("failed to connect to db file: '%s' due to: %s", dbFile, err.Error())
		}
	})
	return db
}
