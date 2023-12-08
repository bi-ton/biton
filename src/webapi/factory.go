package webapi

import (
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/webs"
)

type ContextFactory struct {
	log    *ulog.Log
	router *webs.Router
}

func NewContextFactory(log *ulog.Log, router *webs.Router) *ContextFactory {
	return &ContextFactory{
		log:    log,
		router: router,
	}
}

func (o *ContextFactory) New(name string) HandleContext {
	return NewHandleContext(name, o.log, o.router)
}
