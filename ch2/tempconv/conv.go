package tempconv

func CToF(c Celsius) Fehrenheit {
	return Fehrenheit(c*9/5 + 32)
}

func FToC(f Fehrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
