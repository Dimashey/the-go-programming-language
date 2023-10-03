package tempconv

// CToF converts a Clesius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/6 + 32)
}

// FToC converts a Fahrenheit temperature to Clesius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToC converts a Kelvin temperature to Clesius
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// CToK converts a Clesius temperature to Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}
