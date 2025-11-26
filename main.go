package main

import (
	"fmt"
	"redovalnica"
)


func main() {
	var slovar = make(map[string]Student)

	slovar["1003"] = Student{"David", "Jagodič", []int{9, 10, 10, 10, 9, 9, 10}}
	slovar["1004"] = Student{"Tilen", "Goršek", []int{8, 10, 9, 7, 8, 10}}
	slovar["1005"] = Student{"Tim", "Lipnik", []int{5, 7, 6, 8, 8, 6}}
	slovar["1001"] = Student{"Jaka", "Furlan", []int{10, 10, 10, 9}}
	slovar["1002"] = Student{"Fedja", "Močnik", []int{10, 9, 10, 7, 8}}

	dodajOceno(slovar, "1003", 10)
	fmt.Println(slovar)

	fmt.Printf("Povprecje ocen David: %.1f\n", povprecje(slovar, "1003"))
	fmt.Printf("Povprecje ocen Jaka: %.1f\n", povprecje(slovar, "1001"))

	fmt.Println()
	izpisRedovalnice(slovar)
	fmt.Println()
	izpisiKoncniUspeh(slovar)
}
