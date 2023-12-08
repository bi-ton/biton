package webapi

import (
	"errors"
	"fmt"
	"net/http"
)

type Responce[T any] struct {
	Status int
	Error  error
	Data   T
}

func (o *Responce[T]) Ok() bool {
	return o.Error == nil && o.Status == 0
}

func (o *Responce[T]) RefineError(prefix string) bool {
	if o.Error == nil {
		return false
	}
	o.Error = fmt.Errorf("%s: %v", prefix, o.Error)
	return true
}

func (o *Responce[T]) RefineBadRequest(prefix string) bool {
	if o.RefineError(prefix) {
		o.Status = http.StatusBadRequest
		return true
	}
	return false
}

func (o *Responce[T]) SetForbidden() {
	o.Status = http.StatusForbidden
}

func (o *Responce[T]) SetUnauthorized() {
	o.Status = http.StatusUnauthorized
}

func (o *Responce[T]) BadRequest(s string) {
	o.Error = errors.New(s)
	o.Status = http.StatusBadRequest
}

func (o *Responce[T]) Forbidden(s string) {
	o.Error = errors.New(s)
	o.Status = http.StatusForbidden
}

func (o *Responce[T]) Unauthorized(s string) {
	o.Error = errors.New(s)
	o.SetUnauthorized()
}

func (o *Responce[T]) RefineErrorf(s string, args ...any) bool {
	return o.RefineError(fmt.Sprintf(s, args...))
}

func (o *Responce[T]) RefineBadRequestf(s string, args ...any) bool {
	return o.RefineBadRequest(fmt.Sprintf(s, args...))
}

func (o *Responce[T]) BadRequestf(s string, args ...any) {
	o.BadRequest(fmt.Sprintf(s, args...))
}

func (o *Responce[T]) Forbiddenf(s string, args ...any) {
	o.Forbidden(fmt.Sprintf(s, args...))
}

func (o *Responce[T]) Unauthorizedf(s string, args ...any) {
	o.Unauthorized(fmt.Sprintf(s, args...))
}
