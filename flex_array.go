package utilitybelt

import (
  _ "fmt"
)

type FlexArray []interface{}

func NewFlexArray() *FlexArray {
  return &FlexArray{}
}

func (f *FlexArray) Insert(item interface{}, index int) *FlexArray {
  temp := *f
  temp = append(temp[:index], append([]interface{}{item}, temp[index:]...)...)
  *f = temp
  return f
}

func (f *FlexArray) Remove(index int) *FlexArray {
  temp := *f
  copy(temp[index:], temp[index+1:])
  temp[len(temp)-1] = nil
  temp = temp[:len(temp)-1]
  *f = temp
  return f
}

func (f *FlexArray) Append(items ...interface{}) *FlexArray {
  temp := append(*f, items...)
  *f = temp
  return f
}

func (f *FlexArray) Prepend(items ...interface{}) *FlexArray {
  temp := append(items, *f...)
  *f = temp
  return f
}

func (f *FlexArray) Push(items ...interface{}) *FlexArray {
  return f.Prepend(items...)
}

func (f *FlexArray) Queue(items ...interface{}) *FlexArray {
  return f.Append(items...)
}

func (f *FlexArray) Dequeue() interface{} {
  temp := *f
  item := temp[0]
  *f = temp[1:]
  return item
}

func (f *FlexArray) Each(fn func(interface{})) *FlexArray {
  temp := *f
  for _, item := range temp {
    fn(item)
  }
  return f
}

func (f *FlexArray) EachWithIndex(fn func(interface{}, int)) *FlexArray {
  temp := *f
  for index, item := range temp {
    fn(item, index)
  }
  return f
}

func (f *FlexArray) Map(fn func(interface{}) interface{}) *FlexArray {
  temp := *f
  for index, item := range temp {
    temp[index] = fn(item)
  }
  *f = temp
  return f
}
