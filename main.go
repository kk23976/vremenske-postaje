package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kk23976/vremenske-postaje/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "Aplikacija za delo z vremenskimi postajami in meritvami",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Usage: "Najnižja dovoljena temperatura meritve (°C)",
				Value: -50,
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Usage: "Najvišja dovoljena temperatura meritve (°C)",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "Najmanjše število meritev, potrebnih za izračun statistike",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float64("minTemp")
			maxTemp := cmd.Float64("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			slovarPostaj := make(map[string]vreme.Postaja)

			vreme.DodajPostajo(slovarPostaj, "KP01", "Koper Luka", "Koper")
			vreme.DodajPostajo(slovarPostaj, "KR01", "Kredarica", "Triglav")
			vreme.DodajPostajo(slovarPostaj, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
			vreme.DodajPostajo(slovarPostaj, "MB01", "Maribor Tabor", "Maribor")

			vreme.DodajMeritev(slovarPostaj, "KP01", 24.5, minTemp, maxTemp)
			vreme.DodajMeritev(slovarPostaj, "KP01", 27.0, minTemp, maxTemp)
			vreme.DodajMeritev(slovarPostaj, "KP01", 26.1, minTemp, maxTemp)
			vreme.DodajMeritev(slovarPostaj, "KP01", 67.0, minTemp, maxTemp)
			vreme.DodajMeritev(slovarPostaj, "KP02", 28.4, minTemp, maxTemp)
			vreme.DodajMeritev(slovarPostaj, "KP01", 28.4, minTemp, maxTemp)

			vreme.DodajMeritev(slovarPostaj, "KR01", 2.0, minTemp, maxTemp)
			vreme.DodajMeritev(slovarPostaj, "KR01", -1.5, minTemp, maxTemp)
			vreme.DodajMeritev(slovarPostaj, "KR01", 3.5, minTemp, maxTemp)

			vreme.DodajMeritev(slovarPostaj, "LJ01", 18.2, minTemp, maxTemp)
			vreme.DodajMeritev(slovarPostaj, "LJ01", 21.5, minTemp, maxTemp)
			vreme.DodajMeritev(slovarPostaj, "LJ01", 19.8, minTemp, maxTemp)
			vreme.DodajMeritev(slovarPostaj, "LJ01", 16.5, minTemp, maxTemp)

			vreme.DodajMeritev(slovarPostaj, "MB01", 12.4, minTemp, maxTemp)

			fmt.Println()
			vreme.IzpisMeritev(slovarPostaj)
			fmt.Println()
			vreme.IzpisPorocila(slovarPostaj, stMeritev)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}