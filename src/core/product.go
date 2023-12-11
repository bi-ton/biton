package core

import "github.com/uptrace/bun"

type Product struct {
	bun.BaseModel `bun:"table:main.products"`
	Id            int
	Name          string
	Price         float64
	Done          bool
}
