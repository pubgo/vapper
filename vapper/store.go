package vapper

import (
	"github.com/dave/flux"
)

func (t *Vapper) InitStore(store flux.StoreInterface) {
	if store == nil {
		panic("store error")
	}
	t.stores = append(t.stores, store)
}
