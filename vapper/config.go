package vapper

import "reflect"

func RegisterConfig(cfg interface{}) {
	if cfg == nil {
		panic("config error")
	}

	_vapper.cfg = reflect.ValueOf(cfg)
}
