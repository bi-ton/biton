package webapi

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/msw-x/moon/parse"
)

type Request[T any] struct {
	Data T

	r *http.Request
}

func (o Request[T]) RemoteAddr() string {
	return o.r.RemoteAddr
}

func (o Request[T]) Var(name string) string {
	return mux.Vars(o.r)[name]
}

func (o Request[T]) VarInt(name string) (int, error) {
	return strconv.Atoi(o.Var(name))
}

func (o Request[T]) VarIntDef(name string) int {
	i, _ := o.VarInt(name)
	return i
}

func (o Request[T]) VarInt64(name string) (int64, error) {
	return strconv.ParseInt(o.Var(name), 10, 64)
}

func (o Request[T]) VarInt64Def(name string) int64 {
	i, _ := o.VarInt64(name)
	return i
}

func (o Request[T]) VarUUID(name string) (uuid.UUID, error) {
	return uuid.Parse(o.Var(name))
}

func (o Request[T]) ParamExists(name string) bool {
	return o.Param(name) != ""
}

func (o Request[T]) Param(name string) string {
	return o.r.URL.Query().Get(name)
}

func (o Request[T]) ParamBool(name string) (bool, error) {
	return parse.Bool(o.Param(name))
}

func (o Request[T]) ParamInt(name string) (int, error) {
	return strconv.Atoi(o.Param(name))
}

func (o Request[T]) ParamInt64(name string) (int64, error) {
	return strconv.ParseInt(o.Param(name), 10, 64)
}

func (o Request[T]) ParamUUID(name string) (uuid.UUID, error) {
	return uuid.Parse(o.Param(name))
}

func (o Request[T]) ParamBoolAnyway(name string) (r bool, err error) {
	if o.ParamExists(name) {
		return parse.Bool(o.Param(name))
	}
	return
}

func (o Request[T]) Params(v any) error {
	o.r.ParseForm()
	return schema.NewDecoder().Decode(v, o.r.Form)
}

func (o *Request[T]) DataFromParams() error {
	return o.Params(&o.Data)
}
