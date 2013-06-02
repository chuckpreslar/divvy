package divvy

import (
  "testing"
)

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

func TestPrepend(t *testing.T) {
  d := New()
  prepend := []interface{}{1, 2, 3}
  d.Append(4).Prepend(prepend...)
  expected := []interface{}{1, 2, 3, 4}
  for i, v := range *d {
    if expected[i] != v {
      t.Errorf("Append(%v) => %v, want &[%v]\n", expected, d, expected)
    }
  }
}

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
  supplied := []interface{}{1, 2, 3, 4, 5}
  index := 2
  expected := supplied[index]
  got := New().Append(supplied...).RemoveAt(index)
  if expected != got {
    t.Errorf("RemoveAt(%v) = %v, want %v\n", index, got, expected)
  }
  if New().RemoveAt(10) != nil {
    t.Errorf("Expected RemoveAt to return nil when provided an unoccupied index.")
  }
}

func TestSplice(t *testing.T) {
  d := New().Append(1, 2, 3, 4, 5)
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

func TestPop(t *testing.T) {
  expected := 5
  got := New().Append(1, 2, 3, 4, 5).Pop()
  if expected != got {
    t.Errorf("Pop() = %v, want %v\n", got, expected)
  }
}

func TestDequeue(t *testing.T) {
  expected := 1
  got := New().Append(1, 2, 3, 4, 5).Dequeue()
  if expected != got {
    t.Errorf("Dequeue() = %v, want %v\n", got, expected)
  }
}

func TestEach(t *testing.T) {
  supplied := []interface{}{1, 2, 3, 4, 5}
  total := len(supplied)
  called := 0
  New().Append(supplied...).Each(func(item interface{}) {
    called += 1
  })
  if called != total {
    t.Errorf("Expected each to be called %v times, was called %v times instead.", total, called)
  }
}
