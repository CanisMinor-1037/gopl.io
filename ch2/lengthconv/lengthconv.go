// lengthconv package performs feet and meter length computation
package lengthconv

import "fmt"

type Foot float64
type Meter float64

func (f Foot) String() string {
	return fmt.Sprintf("%g foot", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g meter", m)
}
