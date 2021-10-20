package main

import "fmt"

type store struct {
	name string
	lat  float64
	long float64
}

func (inp store) address() string {
	return fmt.Sprintf("Address for %s: %f %f\n", inp.name, inp.lat, inp.long)
}

type name string

func (n name) hello() string {
	return string(n)
}
