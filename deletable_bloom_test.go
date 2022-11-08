package abloom_test

import (
	"testing"

	"github.com/arpitbbhayani/abloom"
)

func TestDeletableBloom(t *testing.T) {
	b := abloom.NewDeletableBloom(3, []int{3719237192, 8981437}, 4)
	b.Put([]byte("a"))
	b.Put([]byte("b"))
	t.Fail()
}
