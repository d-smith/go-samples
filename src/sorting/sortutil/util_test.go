package sortutil

import "testing"

type testpair struct {
  values []int
  sorted bool
}

var testdata = []testpair {
  { []int {1,4,3,5,6,7,8}, false},
  { []int {1,2,3,4,5}, true},
  { []int {1}, true}}

func TestIsSorted(t *testing.T) {
  for _, pair := range testdata {
    v := IsSorted(pair.values)
    if(v != pair.sorted) {
      t.Error("For ", pair.values, " expected ", pair.sorted, " got ", v)
    }
  }
}
