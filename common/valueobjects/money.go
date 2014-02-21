package valueobjects

import (
  "strconv"
  "strings"
)

type Money float64

func (m *Money) FromString(str string) error {
  s := strings.Replace(str, ",", ".", 1)
  f, e := strconv.ParseFloat(s, 64)
  *m = Money(f)
  return e
}
