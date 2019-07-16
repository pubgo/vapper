package vapper

import (
	"github.com/dave/flux"
)

func RegisterStore(store flux.StoreInterface) {
	if store == nil {
		panic("store error")
	}
	_vapper.stores = append(_vapper.stores, store)
}
