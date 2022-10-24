package lengthconv

func FToM(f Foot) Meter {
	return Meter(f / 3.2808)
}

func MToF(m Meter) Foot {
	return Foot(m * 3.2808)
}
