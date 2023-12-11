package core

import (
	"github.com/msw-x/moon/db"
	"github.com/msw-x/moon/ulog"
)

type Core struct {
	log *ulog.Log
	db  *db.Db
}

func New(db *db.Db) *Core {
	o := new(Core)
	o.log = ulog.New("core").WithLifetime()
	o.db = db
	return o
}

func (o *Core) Close() {
	o.log.Close()
}

func (o *Core) Run() {

}

func (o *Core) Products() (l []Product, err error) {
	err = o.db.SelectAll(&l)
	return
}
