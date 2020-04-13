package main

import "fmt"

func main() {

	var v_0, s_0, a float64
	fmt.Println("Enter values of initial velocity, initial displacement and acceleration one value per line")
	fmt.Scan(&v_0)
	fmt.Scan(&s_0)
	fmt.Scan(&a)
	fmt.Println("Enter value of time")
	var time float64
	fmt.Scan(&time)
	fn := GenDisplaceFn(a, v_0, s_0)
	fmt.Println("Displacement: ", fn(time))
}

func GenDisplaceFn(a, v_0, s_0 float64) func(float64) float64 {
	return func(t float64) float64 {
		s := (0.5 * a * t * t) + (v_0 * t) + s_0
		return s
	}
}
