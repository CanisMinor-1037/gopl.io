// Package tempconv performs Kelvin, Celsius and Fehrenheit temperature computations
package tempconv

import "fmt"

type Celsius float64    // 摄氏温度
type Fehrenheit float64 // 华氏温度
type Kelvin float64     // 开尔文温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func (f Fehrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
