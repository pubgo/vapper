package bufferpool

import (
	"github.com/pubgo/errors"
	"testing"
)

func TestBufferPool(t *testing.T) {
	defer errors.Assert()

	buff := GetBuffer()
	buff.WriteString("do be do be do")
	PutBuffer(buff)

	errors.TT("do be do be do" == buff.String() || buff.Len() != 0, "error").
		M("input1", buff.String()).
		M("input2", buff.Len()).
		Done()
}
