package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/alecthomas/units"
	"gopkg.in/alecthomas/kingpin.v2"
)

type summer interface {
	parse(s string) (float64, error)
	format(n float64) string
}

var (
	formatNumeric  = kingpin.Flag("numeric", "Numeric (default)").Short('n').Bool()
	formatHuman    = kingpin.Flag("human-readable", "Human readable (prefixes like M, K, etc)").Short('h').Bool()
	humanIECPrefix = kingpin.Flag("human-iec", "Use IEC prefixes for human-readable (Mi, Ki, etc)").Bool()
	humanUnit      = kingpin.Flag("human-unit", "Specify a unit to follow the prefix").PlaceHolder("UNIT").String()
)

func main() {
	kingpin.Parse()

	if *formatNumeric && *formatHuman {
		fmt.Fprintf(os.Stderr, "Exactly one format must be selected")
		os.Exit(1)
	}

	var sr summer
	switch {
	case *formatNumeric:
		sr = numericSummer{}
	case *formatHuman:
		var u map[string]float64
		if *humanIECPrefix {
			u = units.MakeUnitMap(fmt.Sprintf("i%s", *humanUnit), *humanUnit, 1024)
		} else {
			u = units.MakeUnitMap(*humanUnit, *humanUnit, 1000)
		}
		sr = humanSummer{u}
	default:
		sr = numericSummer{}
	}

	var output float64

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if err := s.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stream: %v\n", err)
			os.Exit(1)
		}

		inc, err := sr.parse(s.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not parse %q: %v\n", s.Text(), err)
			os.Exit(1)
		}

		output += inc
	}

	fmt.Fprintln(os.Stdout, sr.format(output))
}

type numericSummer struct{}

func (numericSummer) parse(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
func (numericSummer) format(n float64) string {
	return strconv.FormatFloat(n, 'g', -1, 64)
}

type humanSummer struct {
	units map[string]float64
}

func (h humanSummer) parse(s string) (float64, error) {
	n, err := units.ParseUnit(s, h.units)
	return float64(n), err
}

func (h humanSummer) format(n float64) string {
	var unit string
	smallestDiff := n
	for k, v := range h.units {
		diff := n - v
		if diff >= 0 && diff < smallestDiff {
			unit = k
			smallestDiff = diff
		}
	}

	return fmt.Sprintf("%.1f%s", n/h.units[unit], unit)
}
