// ex2.2 prints measurements given on the command line or stdin in various
// units.
package main

import (
	"bufio"
	"fmt"
	"log"
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

func FromFeet(f float64) Distance {
	return Distance{f * 0.3048}
}

func FromMeters(m float64) Distance {
	return Distance{m}
}

func (d Distance) String() string {
	return fmt.Sprintf("%.3gm = %.3gft", d.meters, d.Feet())
}

func (d Distance) Meters() float64 {
	return d.meters
}

func (d Distance) Feet() float64 {
	return d.meters / 0.3048
}

type Temperature float64

func FromCelcius(c float64) Temperature {
	return Temperature(c)
}

func FromFarenheit(f float64) Temperature {
	return Temperature((f * 5 / 9) - 32)
}

func (t Temperature) String() string {
	return fmt.Sprintf("%.3gC = %.3gF", t.Celcius(), t.Farenheit())
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
		return FromMeters(f), nil
	case "\"", "ft":
		return FromFeet(f), nil
	case "c":
		return FromCelcius(f), nil
	case "f":
		return FromFarenheit(f), nil
	default:
		return Distance{}, fmt.Errorf("Unexpected unit %v", unit)
	}
}

func printMeasurement(s string) {
	re := regexp.MustCompile(`(-?\d+(?:\.\d+)?)([A-Za-z]+)`)
	match := re.FindStringSubmatch(s)
	if match == nil {
		log.Fatalf("Expecting <number><unit>, got %q", s)
	}
	f, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		log.Fatalf("%v isn't a number.", match[1])
	}
	if match[2] == "" {
		log.Fatalf("No unit specified.")
	}
	unit := match[2]
	m, err := newMeasurement(f, unit)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(m)
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			printMeasurement(arg)
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			printMeasurement(scan.Text())
		}
	}
}
