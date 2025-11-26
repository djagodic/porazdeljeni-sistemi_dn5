// Package redovalnica provides a simulated grades storage.
//
// Example usage:
//
// 	var slovar = make(map[string]redovalnica.Student)
// 	slovar["1234"] = redovalnica.Student{"Name", "Surname", []int{1, 5, ...}}
//
// redovalnica.DodajOceno(slovar, "1234", 10)
//
// redovalnica.izpisRedovalnice(slovar)
//
// redovalnica.izpisiKoncniUspeh(slovar)
package redovalnica

import (
	"fmt"
)

type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena, min, max int) {
	if ocena < min || ocena > max {
		fmt.Println("Ocena ni v ustreznem območju")
		return
	}

	//student je sedaj kopija, na kateri delamo
	student, ok := studenti[vpisnaStevilka]

	if !ok {
		fmt.Println("Študent s to vpisno številko ne obstaja")
		return
	}

	//moramo ga zapisati nazaj v slovar
	student.Ocene = append(student.Ocene, ocena)
	studenti[vpisnaStevilka] = student

	// ne gre, go tegale ne dovoli -> šlo bi, če bi imel v štartu kazalce na študente v slovarju.....kompliciranje
	//studenti[vpisnaStevilka].ocene = append(studenti[vpisnaStevilka].ocene, ocena)
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, ok := studenti[vpisnaStevilka]

	if !ok {
		return -1.0
	}

	if len(student.Ocene) < 6 {
		return 0.0
	}
	
	vsota := 0
	for _, o := range(student.Ocene) {
		vsota += o
		//fmt.Println(o)
	}

	return float64(vsota) / float64(len(student.Ocene))
}

func IzpisRedovalnice(studenti map[string]Student) {
	fmt.Println("REDOVALNICA")

	for studentKey := range(studenti) {
		fmt.Printf("%s - %s %s: %v\n", studentKey, studenti[studentKey].Ime, studenti[studentKey].Priimek, studenti[studentKey].Ocene)
	}

}

func IzpisiKoncniUspeh(studenti map[string]Student) {
	for studentKey := range(studenti) {
		povpr := povprecje(studenti, studentKey)
		fmt.Printf("%s %s: povprečna ocena %.1f ->", studenti[studentKey].Ime, studenti[studentKey].Priimek, povpr)

		if povpr >= 9 {
			fmt.Println("Odličen študent!")
		} else if povpr >= 6 {
			fmt.Println("Povprečen študent")
		} else {
			fmt.Println("Neuspešen študent")
		}
	}
}