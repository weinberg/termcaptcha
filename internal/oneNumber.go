package internal

import (
  "strconv"
)

type xy [2]int

var four = []xy{{6, 1}, {6, 9}, {1, 5}, {9, 5}}

var curr xy

func OneNumber() (word string, c string) {
  word = "4"
  c = "\n\n\n\n\n\n\n\n\n\n"
  curr = xy{0,9}
  to(four[0], &c)
  for _, point := range four {
    c += "X"
    c += CUB(1)
    to(point, &c)
  }
  c += "X"
  return word, c
}

func to(to xy, c *string) {
  if curr[0] > to[0] {
    *c += CUB(curr[0] - to[0])
  }
  if curr[0] < to[0] {
    *c += CUF(to[0] - curr[0])
  }
  if curr[1] > to[1] {
    *c += CUD(curr[1] - to[1])
  }
  if curr[1] < to[1] {
    *c += CUU(to[1] - curr[1])
  }
  curr = to
}

func CUB(n int) string {
  return CU(n, "D")
}

func CUD(n int) string {
  return CU(n, "B")
}

func CUF(n int) string {
  return CU(n, "C")
}

func CUU(n int) string {
  return CU(n, "A")
}

func CU(n int, c string) string {
  return e + strconv.Itoa(n) + c
}
