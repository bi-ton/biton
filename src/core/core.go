package core

import "github.com/msw-x/moon/db"

type Core struct {
	db *db.Db
}

func New(db *db.Db) *Core {
	o := new(Core)
	o.db = db
	return o
}

func (o *Core) Close() {

}

func (o *Core) Run() {

}

func (o *Core) Products() (l []Product, err error) {
	err = o.db.SelectAll(&l)
	return
}
