package webapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/msw-x/moon/uerr"
	"github.com/msw-x/moon/ujson"
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/webs"
)

type HandleContext struct {
	log    *ulog.Log
	router *webs.Router
}

func NewHandleContext(branch string, log *ulog.Log, router *webs.Router) HandleContext {
	if branch != "" {
		log = log.Branch(branch)
	}
	return HandleContext{
		log:    log,
		router: router.Branch(branch),
	}
}

func processHeader[ResponceData any](h http.Header, w *Responce[ResponceData], fn func(http.Header, *Responce[ResponceData])) {
	defer uerr.Recover(func(s string) {
		w.Error = fmt.Errorf("header: %s", s)
	})
	if fn != nil {
		fn(h, w)
	}
}

func processHandle[RequestData any, ResponceData any](r Request[RequestData], w *Responce[ResponceData], fn func(Request[RequestData], *Responce[ResponceData])) {
	defer uerr.Recover(func(s string) {
		w.Error = fmt.Errorf("handle: %s", s)
	})
	if fn != nil {
		fn(r, w)
	}
}

func Handle[RequestData any, ResponceData any](
	ctx HandleContext,
	method string,
	url string,
	header func(http.Header, *Responce[ResponceData]),
	handle func(Request[RequestData], *Responce[ResponceData]),
) {
	ctx.router.Options(url, func(w http.ResponseWriter, r *http.Request) {
		corsResponse(w)
	})
	ctx.router.Handle(method, url, func(w http.ResponseWriter, r *http.Request) {
		corsResponse(w)
		var request Request[RequestData]
		var responce Responce[ResponceData]
		processHeader(r.Header, &responce, header)
		if responce.Ok() {
			request.r = r
			if reflect.TypeOf(request.Data) != reflect.TypeOf(void{}) {
				if method == http.MethodGet || method == http.MethodDelete {
					request.DataFromParams()
				} else {
					body, _ := ioutil.ReadAll(r.Body)
					if len(body) > 0 {
						responce.Error = json.Unmarshal(body, &request.Data)
						if !responce.Ok() {
							responce.RefineBadRequest("unmarshal json")
						}
					} else {
						responce.BadRequest("request is empty")
					}
				}
			}
			if responce.Ok() {
				processHandle(request, &responce, handle)
			}
		}
		if !responce.Ok() {
			ctx.log.Error(webs.RequestName(r), responce.Error)
			if responce.Status == 0 {
				responce.Status = http.StatusInternalServerError
			}
		}
		if responce.Status > 0 {
			w.WriteHeader(responce.Status)
		}
		var body []byte
		w.Header().Set("Content-Type", "application/json")
		if responce.Ok() {
			if reflect.TypeOf(responce.Data) != reflect.TypeOf(void{}) {
				body, _ = ujson.MarshalLowerCase(responce.Data)
			}
		} else {
			body, _ = ujson.MarshalLowerCase(struct {
				Error string
			}{
				Error: fmt.Sprint(responce.Error),
			})
		}
		w.Write(body)
	})
}

func Get[RequestData any, ResponceData any](ctx HandleContext, url string, handle func(Request[RequestData], *Responce[ResponceData])) {
	Handle(ctx, http.MethodGet, url, nil, handle)
}

func Post[RequestData any, ResponceData any](ctx HandleContext, url string, handle func(Request[RequestData], *Responce[ResponceData])) {
	Handle(ctx, http.MethodPost, url, nil, handle)
}

func Put[RequestData any, ResponceData any](ctx HandleContext, url string, handle func(Request[RequestData], *Responce[ResponceData])) {
	Handle(ctx, http.MethodPut, url, nil, handle)
}

func Delete[RequestData any, ResponceData any](ctx HandleContext, url string, handle func(Request[RequestData], *Responce[ResponceData])) {
	Handle(ctx, http.MethodDelete, url, nil, handle)
}

func corsResponse(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}
