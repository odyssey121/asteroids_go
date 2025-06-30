package main

import (
	"fmt"
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func (v *Vector) Normalize() Vector {
	fmt.Print("Normolize")
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return Vector{v.X / magnitude, v.Y / magnitude}
}
