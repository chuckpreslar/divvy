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
      t.Errorf("Prepend(%v) => %v, want &[%v]\n", expected, d, expected)
    }
  }
}

func TestAtIndex(t *testing.T) {
  item := 1
  d := New()
  if got := d.Append(item).AtIndex(0); item != got {
    t.Errorf("Expected %v, got %v.\n", item, got)
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
  expected, got := 5, New().Append(1, 2, 3, 4, 5).Pop()
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
  total, called := len(supplied), 0
  New().Append(supplied...).Each(func(item interface{}) {
    called += 1
  })
  if called != total {
    t.Errorf("Expected each to be called %v times, was called %v times instead.\n", total, called)
  }
}

func TestEachWithIndex(t *testing.T) {
  supplied := []interface{}{1, 2, 3, 4, 5}
  total, called, sum, expected := len(supplied), 0, 0, 10
  New().Append(supplied...).EachWithIndex(func(item interface{}, index int) {
    called += 1
    sum += index
  })
  if called != total {
    t.Errorf("Expected each to be called %v times, was called %v times instead.\n", total, called)
  } else if sum != expected {
    t.Errorf("Expected sum of indices to be %v, was %v instead.\n", expected, sum)
  }
}

func TestSelect(t *testing.T) {
  supplied := []interface{}{1, 2, 3, 4, 5}
  expected := []interface{}{2, 4}
  got := New().Append(supplied...).Select(func(item interface{}) bool {
    return (item.(int))%2 == 0
  })
  if len(*got) != len(expected) {
    t.Errorf("Expected length of %v to equal %v\n", *got, expected)
  } else {
    for i, v := range *got {
      if v != expected[i] {
        t.Errorf("Expected %v to equal %v\n", *got, expected)
      }
    }
  }
}

func TestReject(t *testing.T) {
  supplied := []interface{}{1, 2, 3, 4, 5}
  expected := []interface{}{1, 3, 5}
  got := New().Append(supplied...).Select(func(item interface{}) bool {
    return (item.(int))%2 != 0
  })
  if len(*got) != len(expected) {
    t.Errorf("Expected length of %v to equal %v\n", *got, len(expected))
  } else {
    for i, v := range *got {
      if v != expected[i] {
        t.Errorf("Expected %v to equal %v\n", *got, expected)
      }
    }
  }
}

func TestIndexOf(t *testing.T) {
  supplied := []interface{}{1, 2, 3, 4, 5}
  index := 4
  got := New().Append(supplied...).IndexOf(supplied[index])
  if got != index {
    t.Errorf("IndexOf(%v) = %v, want %v", supplied[index], got, index)
  }
}

func TestLastIndexOf(t *testing.T) {
  supplied := []interface{}{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
  index := 6
  got := New().Append(supplied...).LastIndexOf(supplied[index])
  if got != index {
    t.Errorf("LastIndexOf(%v) = %v, want %v", supplied[index], got, index)
  }
}

func TestReverse(t *testing.T) {
  supplied, expected := []interface{}{1, 2, 3, 4, 5}, []interface{}{5, 4, 3, 2, 1}
  got := New().Append(supplied...).Reverse()
  if len(*got) != len(expected) {
    t.Errorf("Expected length of %v to equal %v\n", *got, len(expected))
  } else {
    for i, v := range *got {
      if v != expected[i] {
        t.Errorf("Expected %v to equal %v\n", v, expected[i])
      }
    }
  }
}

func TestUnique(t *testing.T) {
  supplied, expected := []interface{}{1, 2, 3, 4, 5, 1, 3, 5}, []interface{}{1, 2, 3, 4, 5}
  got := New().Append(supplied...).Unique()
  for i, v := range *got {
    if v != expected[i] {
      t.Errorf("Expected %v to equal %v\n", v, expected[i])
    }
  }
}

func TestSort(t *testing.T) {
  supplied, expected := []interface{}{3, 2, 4, 1, 5}, []interface{}{1, 2, 3, 4, 5}
  got := New().Append(supplied...).Sort(func(left, right interface{}) bool {
    return left.(int) < right.(int)
  })
  for i, v := range *got {
    if v != expected[i] {
      t.Errorf("Expected %v to equal %v\n", v, expected[i])
    }
  }
}

func TestContains(t *testing.T) {
  supplied := []interface{}{1, 2, 3, 4, 5}
  index := 4
  if !New().Append(supplied...).Contains(supplied[index]) {
    t.Errorf("Expected %v to contain element %v\n", supplied, supplied[index])
  }
}
