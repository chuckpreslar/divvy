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
