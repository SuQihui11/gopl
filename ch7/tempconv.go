package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

type celsiusFlag Celsius

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		*f = celsiusFlag(value)
		return nil
	case "F", "°F":
		*f = celsiusFlag(FToC(Fahrenheit(value)))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (f *celsiusFlag) String() string {
	return (*Celsius)(f).String()
}

// 工厂注册函数，将这个底层的值返回
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag(value)
	flag.CommandLine.Var(&f, name, usage)
	return (*Celsius)(&f)
}

var temp = CelsiusFlag("temp", 20.0, "temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
