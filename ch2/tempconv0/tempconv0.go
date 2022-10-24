// Package tempconv performs Celsius and Fehrenheit temperature computations
package tempconv

type Celsius float64    // 摄氏温度
type Fehrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fehrenheit {
	return Fehrenheit(c*9/5 + 32)
}

func FToC(f Fehrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
