package webapi

import (
	"biton/core"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/msw-x/moon/webs"
)

func Routes(c *core.Core, opts Options, version string) *mux.Router {
	router := webs.NewRouter().WithLogRequest(opts.LogHttp)
	log := router.Log()
	api := router.Branch("api")

	api.Get("version", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(struct {
			Version string
		}{
			Version: version,
		})
	})

	f := NewContextFactory(log, api.Branch("v1"))
	products(f.New("products"), c)

	return router.Router()
}
