package types

import (
	"go_pg_es_sync/internals/utils"
)

type Where struct {
	Column    string `json:"column"`
	Condition string `json:"condition"`
}
type Wheres []*Where

func (wheres *Wheres) Parse(config any) error {
	err := utils.ParseMap(config, wheres)
	if err != nil {
		return err
	}
	return nil
}
