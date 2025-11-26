package redovalnica

import (
	"fmt"
)

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	if ocena < 1 || ocena > 10 {
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
	student.ocene = append(student.ocene, ocena)
	studenti[vpisnaStevilka] = student

	// ne gre, go tegale ne dovoli -> šlo bi, če bi imel v štartu kazalce na študente v slovarju.....kompliciranje
	//studenti[vpisnaStevilka].ocene = append(studenti[vpisnaStevilka].ocene, ocena)
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, ok := studenti[vpisnaStevilka]

	if !ok {
		return -1.0
	}

	if len(student.ocene) < 6 {
		return 0.0
	}
	
	vsota := 0
	for _, o := range(student.ocene) {
		vsota += o
		//fmt.Println(o)
	}

	return float64(vsota/len(student.ocene))
}

func IzpisRedovalnice(studenti map[string]Student) {
	fmt.Println("REDOVALNICA")

	for studentKey := range(studenti) {
		fmt.Printf("%s - %s %s: %v\n", studentKey, studenti[studentKey].ime, studenti[studentKey].priimek, studenti[studentKey].ocene)
	}

}

func IzpisiKoncniUspeh(studenti map[string]Student) {
	for studentKey := range(studenti) {
		povpr := povprecje(studenti, studentKey)
		fmt.Printf("%s %s: povprečna ocena %.1f ->", studenti[studentKey].ime, studenti[studentKey].priimek, povpr)

		if povpr >= 9 {
			fmt.Println("Odličen študent!")
		} else if povpr >= 6 {
			fmt.Println("Povprečen študent")
		} else {
			fmt.Println("Neuspešen študent")
		}
	}
}