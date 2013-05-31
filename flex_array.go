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
