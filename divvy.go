package divvy

import (
  _ "fmt"
)

type Divvy []interface{}

func New() *Divvy {
  return &Divvy{}
}

func (d *Divvy) InsertAt(item interface{}, index int) *Divvy {
  temp := *d
  temp = append(temp[:index], append([]interface{}{item}, temp[index:]...)...)
  *d = temp
  return d
}

func (d *Divvy) RemoveAt(index int) *Divvy {
  temp := *d
  copy(temp[index:], temp[index+1:])
  temp[len(temp)-1] = nil
  temp = temp[:len(temp)-1]
  *d = temp
  return d
}

func (d *Divvy) Splice(index, count int) *Divvy {
  for i := 0; i < count; i += 1 {
    d.RemoveAt(index)
  }
  return d
}

func (d *Divvy) Append(items ...interface{}) *Divvy {
  temp := append(*d, items...)
  *d = temp
  return d
}

func (d *Divvy) Prepend(items ...interface{}) *Divvy {
  temp := append(items, *d...)
  *d = temp
  return d
}

func (d *Divvy) Push(items ...interface{}) *Divvy {
  return d.Append(items...)
}

func (d *Divvy) Pop() interface{} {
  temp := *d
  index := len(temp) - 1
  item := temp[index]
  *d = temp[:index]
  return item
}

func (d *Divvy) Queue(items ...interface{}) *Divvy {
  return d.Append(items...)
}

func (d *Divvy) Dequeue() interface{} {
  temp := *d
  item := temp[0]
  *d = temp[1:]
  return item
}

func (d *Divvy) Each(fn func(interface{})) *Divvy {
  temp := *d
  for _, item := range temp {
    fn(item)
  }
  return d
}

func (d *Divvy) EachWithIndex(fn func(interface{}, int)) *Divvy {
  temp := *d
  for index, item := range temp {
    fn(item, index)
  }
  return d
}

func (d *Divvy) Map(fn func(interface{}) interface{}) *Divvy {
  temp := *d
  for index, item := range temp {
    temp[index] = fn(item)
  }
  *d = temp
  return d
}

func (d *Divvy) Select(fn func(interface{}) bool) *Divvy {
  temp := *d
  selection := []interface{}{}
  for _, item := range temp {
    if fn(item) {
      selection = append(selection, item)
    }
  }
  result := Divvy(selection)
  return &result
}

func (d *Divvy) Reject(fn func(interface{}) bool) *Divvy {
  temp := *d
  rejection := []interface{}{}
  for _, item := range temp {
    if !fn(item) {
      rejection = append(rejection, item)
    }
  }
  result := Divvy(rejection)
  return &result
}

func (d *Divvy) IndexOf(item interface{}) int {
  index := -1
  temp := *d
  for i, v := range temp {
    if v == item {
      index = i
      break
    }
  }
  return index
}

func (d *Divvy) LastIndexOf(item interface{}) int {
  index := d.IndexOf(item)
  if ^index == 0 {
    return index
  } else {
    temp := *d
    for i := index; i < len(temp); i += 1 {
      if item == temp[i] {
        index = i
      }
    }
  }
  return index
}
