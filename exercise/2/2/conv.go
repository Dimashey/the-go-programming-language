package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type (
	Celsius    float64
	Fahrenheit float64
	Feet       float64
	Meters     float64
	Pounds     float64
	Kilograms  float64
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (f Feet) String() string { return fmt.Sprintf("%g feet", f) }

func (m Meters) String() string { return fmt.Sprintf("%g meters", m) }

func (p Pounds) String() string { return fmt.Sprintf("%g pounds", p) }

func (k Kilograms) String() string { return fmt.Sprintf("%g kilograms", k) }

func toCelsius(n float64) Celsius {
	return Celsius(n)
}

func toFahrenheit(n float64) Fahrenheit {
	return Fahrenheit((n * 1.8) + 32)
}

func toFeet(n float64) Feet {
	return Feet(n / 0.3048)
}

func toMeters(n float64) Meters {
	return Meters(n)
}

func toPounds(n float64) Pounds {
	return Pounds(n * 2.2046)
}

func toKilograms(n float64) Kilograms {
	return Kilograms(n)
}

func convert(n float64) []any {
	return []any{toCelsius(n), toFahrenheit(n), toFeet(n), toMeters(n), toPounds(n), toKilograms(n)}
}

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		for _, arg := range args {
			value, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				log.Fatal("no valid numeric argument")
				return
			}

			values := convert(value)

			fmt.Printf("%v, %v, %v, %v, %v, %v\n", values[0], values[1], values[2], values[3], values[4], values[5])
		}
	}
}
