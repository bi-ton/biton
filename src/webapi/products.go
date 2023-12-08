package webapi

import "biton/core"

func products(ctx HandleContext, c *core.Core) {
	Get(ctx, "", func(r Request[void], w *Responce[[]core.Product]) {
		w.Data, w.Error = c.Products()
	})
}
