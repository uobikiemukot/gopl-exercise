package main

import "fmt"

const (
	InchPerMeter = 39.37
	FeetPerInch  = 1.0 / 12.0
	InchPerFeet  = 12.0
	MeterPerInch = 1.0 / 39.37
)

type Feet float64
type Meter float64

func (f Feet) String() string  { return fmt.Sprintf("%g ft", f) }
func (m Meter) String() string { return fmt.Sprintf("%g m", m) }

func FToM(f Feet) Meter { return Meter(f) * InchPerFeet * MeterPerInch }
func MToF(m Meter) Feet { return Feet(m) * InchPerMeter * FeetPerInch }
