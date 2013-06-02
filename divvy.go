// The MIT License (MIT)

// Copyright (c) 2013 Chuck Preslar

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Common array operations for GO lang's slices.
package divvy

import (
  "fmt"
  "strings"
)

//  Divvy type to allow receiver methods to manipulate go slice.
type Divvy []interface{}

// Returns a pointer to a new Divvy type
func New() *Divvy {
  return &Divvy{}
}

// Inserts an supplied item at a given index in a Divvy type, returning the original
// Divvy for continued chaining.
func (d *Divvy) InsertAt(item interface{}, index int) *Divvy {
  temp := *d
  temp = append(temp[:index], append([]interface{}{item}, temp[index:]...)...)
  *d = temp
  return d
}

// Removes an item at a given index from a Divvy type, returning the removed item.
func (d *Divvy) RemoveAt(index int) interface{} {
  temp := *d
  item := temp[index]
  copy(temp[index:], temp[index+1:])
  temp[len(temp)-1] = nil
  temp = temp[:len(temp)-1]
  *d = temp
  return item
}

// Splice removes items starting at a specified index up to a maximum of the given count,
// returning a new Divvy containing the removed items for continued chaining.
func (d *Divvy) Splice(index, count int) *Divvy {
  temp := make([]interface{}, count)
  for i := 0; i < count; i += 1 {
    temp[i] = d.RemoveAt(index)
  }
  result := Divvy(temp)
  return &result
}

// Inserts the given item(s) at the end of the Divvy type, returning the original
// Divvy for continued chaining.
func (d *Divvy) Append(items ...interface{}) *Divvy {
  temp := append(*d, items...)
  *d = temp
  return d
}

// Inserts the given item(s) at the begining of a Divvy type, returning the original
// Divvy for continued chaining
func (d *Divvy) Prepend(items ...interface{}) *Divvy {
  temp := append(items, *d...)
  *d = temp
  return d
}

// Alias method for Append, allowing the Divvy to be thought of as a stack, returning the original
// Divvy for continued chaining
func (d *Divvy) Push(items ...interface{}) *Divvy {
  return d.Append(items...)
}

// Removes and returns an item from the back of the Divvy type so the Divvy may be thought of as
// a stack.
func (d *Divvy) Pop() interface{} {
  temp := *d
  index := len(temp) - 1
  item := temp[index]
  *d = temp[:index]
  return item
}

// Alias method for Appebd, inserting an item(s) to the end of a Divvy type so it may
// be thought of as a queue.
func (d *Divvy) Queue(items ...interface{}) *Divvy {
  return d.Append(items...)
}

// Removes and returns an item from the front of the Divvy type so the Divvy may be thought of as
// a queue.
func (d *Divvy) Dequeue() interface{} {
  temp := *d
  item := temp[0]
  *d = temp[1:]
  return item
}

// Each takes a function which takes a single interface{} type as an argument, calling the function
// with each item stored in the Divvy array and returning it for continued chaining.
func (d *Divvy) Each(fn func(interface{})) *Divvy {
  temp := *d
  for _, item := range temp {
    fn(item)
  }
  return d
}

// EachWithIndex takes a function which takes an interface{} and int types as arguments,
// calling the function with each item stored in the Divvy array and its index,
// and returning the original Divvy for continued chaining.
func (d *Divvy) EachWithIndex(fn func(interface{}, int)) *Divvy {
  temp := *d
  for index, item := range temp {
    fn(item, index)
  }
  return d
}

// Map, similar to Each, takes a function which takes and returns a single interface{} type,
// calling the function with each item stored in the Divvy array and mapping the returned
// results to a new Divvy array.  The new Divvy array is returned for continued chaining.
func (d *Divvy) Map(fn func(interface{}) interface{}) *Divvy {
  temp := *d
  for index, item := range temp {
    temp[index] = fn(item)
  }
  result := Divvy(temp)
  return &result
}

// Select takes a function as an argument that returns a boolean.  Looping through each item
// within the Divvy type, if the supplied function returns true when supplied the item at
// the current index of the initial Divvy, the item is added to a new Divvy array.  The
// new Divvy array is returned for continued chaining.
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

// Select takes a function as an argument that returns a boolean.  Looping through each item
// within the Divvy type, if the supplied function returns false when supplied the item at
// the current index of the initial Divvy, the item is added to a new Divvy array.  The
// new Divvy array is returned for continued chaining.
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

// IndexOf returns the index of the first occurrence of an item in the Divvy, or
// -1 if the item is not found.
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

// IndexOf returns the index of the last occurrence of an item in the Divvy, or
// -1 if the item is not found.
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

// Reverse stores the reverse the Divvy array and returns it as a new Divvy array
// for continued chaining.
func (d *Divvy) Reverse() *Divvy {
  temp := *d
  half, total := len(temp)/2, len(temp)-1
  for i := 0; i < half; i += 1 {
    item := temp[i]
    temp[i] = temp[total-i]
    temp[total-i] = item
  }
  result := Divvy(temp)
  return &result
}

// Unique returns a new Divvy array containing each item stored in the original, removing
// duplicate items.
func (d *Divvy) Unique() *Divvy {
  temp := *d
  unique := Divvy{}
  for _, item := range temp {
    if ^unique.IndexOf(item) == 0 {
      unique.Append(item)
    }
  }
  return &unique
}

func (d *Divvy) Length() int {
  return len(*d)
}

func (d *Divvy) Count() int {
  return d.Length()
}

func (d *Divvy) Join(delimiter string) string {
  temp := make([]string, d.Length())
  for index, item := range *d {
    temp[index] = fmt.Sprintf("%v", item)
  }
  return strings.Join(temp, delimiter)
}
