package vapper

import "reflect"

func (t *Vapper) RegisterConfig(cfg interface{}) {
	if cfg == nil {
		panic("config error")
	}
	t.cfg = reflect.ValueOf(cfg)
}
