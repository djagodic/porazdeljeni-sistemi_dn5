package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/djagodic/porazdeljeni-sistemi_dn5/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "Urejanje in izpis ocen študentov",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "Najmanjše število ocen potrebnih za pozitivno oceno",
				Value: 6,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "Najmanjša dovoljena ocena",
				Value: 1,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "Največja dovoljena ocena",
				Value: 10,
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {

			// Preberemo nastavitve iz CLI stikal
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")
			stOcen := cmd.Int("stOcen")

			// Inicializacija podatkov
			studenti := make(map[string]redovalnica.Student)
			studenti["1003"] = redovalnica.Student{Ime: "David", Priimek: "Jagodič", Ocene: []int{9, 10, 10, 10, 9, 9, 10}}
			studenti["1004"] = redovalnica.Student{Ime: "Tilen", Priimek: "Goršek", Ocene: []int{8, 10, 9, 7, 8, 10}}
			studenti["1005"] = redovalnica.Student{Ime: "Tim", Priimek: "Lipnik", Ocene: []int{5, 7, 6, 8, 8, 6}}
			studenti["1001"] = redovalnica.Student{Ime: "Jaka", Priimek: "Furlan", Ocene: []int{10, 10, 10, 9}}
			studenti["1002"] = redovalnica.Student{Ime: "Fedja", Priimek: "Močnik", Ocene: []int{10, 9, 10, 7, 8}}

			// Prikažemo parametre
			fmt.Println("Uporabljene nastavitve:")
			fmt.Println(" - minOcena:", minOcena)
			fmt.Println(" - maxOcena:", maxOcena)
			fmt.Println(" - stOcen:", stOcen)
			fmt.Println()

			// Primer uporabe DodajOceno
			redovalnica.DodajOceno(studenti, "1003", minOcena + (maxOcena - minOcena)/2, minOcena, maxOcena)

			// Izpis vseh ocen
			redovalnica.IzpisRedovalnice(studenti)
			fmt.Println()

			// Izpis končnega uspeha
			redovalnica.IzpisiKoncniUspeh(studenti, stOcen)
			fmt.Println()

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}