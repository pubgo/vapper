package vapper

import "reflect"

func (t *Vapper) InitConfig(cfg interface{}) {
	if cfg == nil {
		panic("config error")
	}
	t.cfg = reflect.ValueOf(cfg)
}
