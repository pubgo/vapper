package vapper

import (
	"github.com/pubgo/errors"
	"reflect"
)

func Config(cfg interface{}) {
	errors.T(errors.IsNone(cfg), "please init config")
	_vapper.cfg = reflect.ValueOf(cfg)
}
