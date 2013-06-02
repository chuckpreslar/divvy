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

func TestRemoveAt(t *testing.T) {
  d := New()
  d.Append(1, 2, 3, 4, 5)
  index, expected := 4, 5
  if value := d.RemoveAt(index); value != expected {
    t.Errorf("RemoveAt(%v) = %v, want %v\n", index, value, expected)
  }
  if d.RemoveAt(10) != nil {
    t.Errorf("Expected RemoveAt to return nil when provided an unoccupied index.")
  }
}

func TestSplice(t *testing.T) {
  d := New()
  d.Append(1, 2, 3, 4, 5)
  index, number := 0, 3
  expected := []interface{}{1, 2, 3}
  got := *d.Splice(0, 3)
  for i, _ := range expected {
    if expected[i] != got[i] {
      t.Errorf("Splice(%v, %v) = %v, want %v\n", index, number, got, expected)
    }
  }
  if len(*d.Splice(10, 1)) != 0 {
    t.Errorf("Expected Splice to return an empty Divvy when provided an unoccupied index.")
  }
}

func TestAppend(t *testing.T) {
  d := New()
  expected := []interface{}{1, 2, 3}
  d.Append(expected...)
  for i, v := range *d {
    if expected[i] != v {
      t.Errorf("Append(%v) => %v, want &[%v]\n", expected, d, expected)
    }
  }
}
