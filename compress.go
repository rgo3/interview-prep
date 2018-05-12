package main

import (
  "fmt"
)

func compress(s string) string {
  var buffer string
  previous := rune(s[0])
  count := 0
  
  for _, c := range s {
    if c == previous {
      count++
    } else {
      buffer = fmt.Sprintf(buffer + "%s%d", string(previous), count)
      previous = c
      count = 1
    }
  }
  
  buffer = fmt.Sprintf(buffer + "%s%d", string(previous), count)
  
  return buffer
}