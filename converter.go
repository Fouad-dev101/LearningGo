package main

import (
	"fmt"
	"strings"
)

func toCelsius(value float64, unit string) (float64, error) {
	switch unit {
	case "C":
		return value, nil
	case "K":
		return value - 273.15, nil
	case "F":
		return (value - 32) * 5 / 9, nil
	default:
		return 0, fmt.Errorf("unknown unit: %s", unit)
	}
}

func fromCelsius(c float64, unit string) (float64, error) {
	switch unit {
	case "C":
		return c, nil
	case "K":
		return c + 273.15, nil
	case "F":
		return (c * 9 / 5) + 32, nil
	default:
		return 0, fmt.Errorf("unknown unit: %s", unit)
	}
}

func main() {
	var value float64
	var from, to string

	fmt.Print("Value: ")
	if _, err := fmt.Scanln(&value); err != nil {
		fmt.Println("invalid number:", err)
		return
	}

	fmt.Print("From (K, C, F): ")
	if _, err := fmt.Scanln(&from); err != nil {
		fmt.Println("invalid unit:", err)
		return
	}

	fmt.Print("To (K, C, F): ")
	if _, err := fmt.Scanln(&to); err != nil {
		fmt.Println("invalid unit:", err)
		return
	}

	from = strings.ToUpper(from)
	to = strings.ToUpper(to)

	c, err := toCelsius(value, from)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := fromCelsius(c, to)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%.2f %s = %.2f %s\n", value, from, result, to)
}