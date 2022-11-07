package abloom_test

import (
	"testing"

	"github.com/arpitbbhayani/abloom"
)

type tcase struct {
	op     byte
	key    string
	exists bool
}

func TestBloom(t *testing.T) {
	b := abloom.NewBloom(1)
	for _, tc := range []tcase{
		{'p', "a", true},
		{'p', "b", true},
		{'p', "c", true},
		{'c', "a", true},
		{'c', "b", true},
		{'c', "c", true},
		{'c', "d", false}, // false positive case
		{'c', "h", true},  // false positive case
	} {
		switch tc.op {
		case 'p':
			b.Put([]byte(tc.key))
		case 'c':
			if v, err := b.Check([]byte(tc.key)); err != nil || v == !tc.exists {
				t.Errorf("check on key %s should be %v but it is observed to be %v", tc.key, tc.exists, v)
			}
		}
	}
}
