package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Measurement interface {
	String() string
}

type Distance struct {
	meters float64
}

type Temperature float64

func (d Distance) String() string {
	return fmt.Sprintf("%gm", d.meters)
}

func (t Temperature) String() string {
	return fmt.Sprintf("%gC = %gF", t.Celcius(), t.Farenheit())
}

func fromCelcius(c float64) Temperature {
	return Temperature(c)
}

func fromFarenheit(f float64) Temperature {
	return Temperature((f * 5 / 9) - 32)
}

func (t Temperature) Celcius() float64 {
	return float64(t)
}

func (t Temperature) Farenheit() float64 {
	return float64((t * 9 / 5) + 32)
}

func newMeasurement(f float64, unit string) (Measurement, error) {
	unit = strings.ToLower(unit)
	switch unit {
	case "m":
		return Distance{f}, nil
	case "\"", "ft":
		return Distance{(f * 12 * 2.54) / 100}, nil
	case "c":
		return fromCelcius(f), nil
	case "f":
		return fromFarenheit(f), nil
	default:
		return Distance{}, fmt.Errorf("Unexpected unit %v", unit)
	}
}

func main() {
	re := regexp.MustCompile(`(\d+(?:\.\d+)?)([A-Za-z]+)`)
	for _, arg := range os.Args[1:] {
		match := re.FindStringSubmatch(arg)
		f, err := strconv.ParseFloat(match[1], 64)
		if err != nil {
			panic(fmt.Errorf("%v isn't a number.", match[1]))
		}
		if match[2] == "" {
			panic("No unit specific.")
		}
		unit := match[2]
		m, err := newMeasurement(f, unit)
		if err != nil {
			panic(err)
		}
		fmt.Println(m)
	}
}
