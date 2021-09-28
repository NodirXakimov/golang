package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}
func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}
func circleArea(r float64) float64 {
	return math.Pi * r * r
}
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}
func main() {
	names := []string{"Nodir", "Dilshod", "Sherzod", "Hasanboy", "Bahodir", "Xasanboy"}
	cycleNames(names, sayGreeting)
	cycleNames(names, sayBye)
	area := circleArea(12.12)
	fmt.Printf("Area of the circle is %0.2f\n", area)
	fn1, sn1 := getInitials("Nodir Xakimov")
	fn2, sn2 := getInitials("Dilshod")
	fmt.Println(fn1, sn1)
	fmt.Println(fn2, sn2)
}
