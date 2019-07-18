package tests

import (
	"fmt"
	"github.com/pubgo/errors"
	"github.com/pubgo/vapper/pdd/cnst"
	"reflect"
	"testing"
)

type sa interface {
	a(...interface{})
}

type ss struct {
}

func (t *ss) a(i ...interface{}) {
	fmt.Println(i...)
}

func b(c sa) {
	defer errors.Assert()

	_c := reflect.ValueOf(c)
	_, ok := _c.Interface().(sa)
	fmt.Println(ok)
}

func TestName(t *testing.T) {
	b(&ss{})
}

func TestJsPkg(t *testing.T) {
	t.Log(cnst.Default.JsPkg)
}
