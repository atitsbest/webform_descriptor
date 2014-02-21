package valueobjects

import "strconv"

type WorkAmount uint

// Aus String mach WorkAmount.
func (w *WorkAmount) FromString(str string) error {
  f, e := strconv.ParseUint(str, 10, 64)
  *w = WorkAmount(f)
  return e
}
