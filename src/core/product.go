package core

import "github.com/uptrace/bun"

type Product struct {
	bun.BaseModel `bun:"table:main.products"`
	Id            int
	Name          string
	Price         string
	Done          bool
}
