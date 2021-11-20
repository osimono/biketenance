package db

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/buntdb"
)

type IndexContext struct {
	Name    string
	Pattern string
}

type IdItem interface {
	ItemId() string
}

func Save(ctx IndexContext, item IdItem) (err error) {
	err = DB().CreateIndex(ctx.Name, ctx.Pattern, buntdb.IndexString)
	err = DB().Update(func(tx *buntdb.Tx) error {

		marshal, _ := json.Marshal(item)
		oldValue, replaced, err := tx.Set(ctx.Pattern+item.ItemId(), string(marshal), nil)
		if err != nil {
			return err
		}

		fmt.Printf("old value:%q replaced:%t\n", oldValue, replaced)
		return nil
	})
	return
}

func ListAll(ctx IndexContext) (allItems []string, err error) {
	err = DB().CreateIndex(ctx.Name, ctx.Pattern, buntdb.IndexString)
	err = DB().View(func(tx *buntdb.Tx) error {
		err := tx.Ascend(ctx.Name, func(key, value string) bool {
			log.Debugf("found: key '%s' => value: '%s'", key, value)
			allItems = append(allItems, value)
			return true
		})
		return err
	})
	return
}
