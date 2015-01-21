package sorting

import "testing"
import su "sorting/sortutil"

type testpair struct {
  values []int
}

var selectionSortData = []testpair {
  { []int {1,4,3,5,6,7,8}},
  { []int {1,2,3,4,5}},
  { []int {23,5,4,1,100, -100, 8,7,2,6,3,-1}},
  { []int {5,4,1,8,7,2,6,3}},
  { []int {1}}}

var insertionSortData = []testpair {
    { []int {1,4,3,5,6,7,8}},
    { []int {1,2,3,4,5}},
    { []int {23,5,4,1,100, -100, 8,7,2,6,3,-1}},
    { []int {5,4,1,8,7,2,6,3}},
    { []int {1}}}

func TestSelectionSort(t *testing.T) {
  for _, pair := range selectionSortData {
    SelectionSort(pair.values)
    if(su.IsSorted(pair.values) != true) {
      t.Error("For ", pair.values, " expected sorted.")
    }
  }
}

func TestInsertionSort(t *testing.T) {
  for _, pair := range insertionSortData {
    InsertionSort(pair.values)
    if(su.IsSorted(pair.values) != true) {
      t.Error("For ", pair.values, " expected sorted.")
    }
  }
}
