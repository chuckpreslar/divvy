package divvy

import (
  "testing"
)

func TestInsertAt(t *testing.T) {
  d := New()
  index, value := 2, 1
  d.InsertAt(index, value)
  temp := *d
  if temp[index] != value {
    t.Errorf("InsertAt(%v,%v) => &[%v], want %v\n", index, value, temp[index], value)
  }
}
