package jsvapper

import (
	"github.com/pubgo/errors"
	"reflect"
)

func (t *Vapper) Config(cfg interface{}) {
	errors.T(errors.IsNone(cfg), "please init config")
	t.cfg = reflect.ValueOf(cfg)
}
