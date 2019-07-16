package vapper

import (
	"github.com/dave/flux"
	"github.com/pubgo/errors"
)

func Store(store flux.StoreInterface) {
	errors.T(errors.IsNone(store), "please init store")
	_vapper.stores = append(_vapper.stores, store)
}
