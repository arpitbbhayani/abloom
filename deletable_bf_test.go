package abloom_test

import (
	"testing"

	"github.com/arpitbbhayani/abloom"
)

func TestDeletableBF(t *testing.T) {
	b := abloom.NewDeletableBF(1, []int{3719237192, 8981437}, 3)
	for _, tc := range []tcase{
		{'p', "a", true},
		{'p', "b", true},
		{'p', "c", true},
		{'p', "d", true},
		{'p', "f", true},
		{'c', "a", true},
		{'c', "b", true},
		{'c', "c", true},
		{'d', "a", true},  // non deletable
		{'d', "b", true},  // non deletable
		{'d', "c", true},  // non deletable
		{'d', "d", true},  // non deletable
		{'d', "f", false}, // non deletable
	} {
		switch tc.op {
		case 'p':
			b.Put([]byte(tc.key))
		case 'd':
			_, err := b.Delete([]byte(tc.key))
			if err != nil {
				t.Errorf("error while deleting the key %v", tc.key)
			}
			exists, err := b.Check([]byte(tc.key))
			if err != nil {
				t.Errorf("error while checking the key %v", tc.key)
			}
			if exists != tc.exists {
				if tc.exists {
					t.Errorf("key %v should exist even after deleting but it is not", tc.key)
				} else {
					t.Errorf("key %v should not exist after deleting but it is", tc.key)
				}
			}
		case 'c':
			if v, err := b.Check([]byte(tc.key)); err != nil || v == !tc.exists {
				t.Errorf("check on key %s should be %v but it is observed to be %v", tc.key, tc.exists, v)
			}
		}
	}
}
