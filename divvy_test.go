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
