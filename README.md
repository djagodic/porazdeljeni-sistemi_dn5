# porazdeljeni-sistemi_dn5
Publicly available module derived from dn1

Package *redovalnica* provides a simulated storage for students grades.

Example usage:

    var slovar = make(map[string]redovalnica.Student)
	slovar["studentID"] = redovalnica.Student{"Ime", "Priimek", []int{1, 5, ...}}
    ...

    redovalnica.DodajOceno(slovar, "studentID", ocena, minOcena, maxOcena)

    redovalnica.izpisRedovalnice(slovar)

    redovalnica.izpisiKoncniUspeh(slovar, stOcen)
