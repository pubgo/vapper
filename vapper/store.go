package vapper

import (
	"github.com/dave/flux"
	"github.com/pubgo/errors"
)

func (t *Vapper) Store(store flux.StoreInterface) {
	errors.T(errors.IsNone(store), "store is null")
	t.stores = append(t.stores, store)
}
