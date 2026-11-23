// Package vreme omogoča upravljanje vremenskih postaj, shranjevanje
// temperaturnih meritev ter izpis meritev in statističnih podatkov.
package vreme

import "fmt"

// Postaja predstavlja vremensko postajo z imenom, krajem in zbranimi meritvami.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo v slovar doda novo vremensko postajo z dano oznako, imenom in krajem.
func DodajPostajo(slovarPostaj map[string]Postaja, oznaka, ime, kraj string) {
	_, ok := slovarPostaj[oznaka]
	
	if ok {
		fmt.Printf("Postaja z oznako '%s' že obstaja.\n", oznaka)
		return
	}
	slovarPostaj[oznaka] = Postaja{
		ime:     ime,
		kraj:    kraj,
		meritve: []float64{},
	}
}

// DodajMeritev doda novo temperaturno meritev postaji z dano oznako,
// če je temperatura znotraj dovoljenega območja [minTemp, maxTemp].
func DodajMeritev(slovarPostaj map[string]Postaja, oznaka string, temperatura float64, minTemp float64, maxTemp float64) {
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Printf("Temperatura %.1f °C ni v dovoljenem območju (med %.1f °C in %.1f °C).\n", temperatura, minTemp, maxTemp)
		return
	}

	e, ok := slovarPostaj[oznaka]
	if !ok {
		fmt.Printf("Postaja z oznako '%s' ne obstaja.\n", oznaka)
		return
	}

	e.meritve = append(e.meritve, temperatura)
	slovarPostaj[oznaka] = e
}

func statistika(slovarPostaj map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	e, ok := slovarPostaj[oznaka]
	if !ok || len(e.meritve) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja = e.meritve[0]
	najvisja = e.meritve[0]
	vsota := 0.0

	for _, meritev := range e.meritve {
		if meritev < najnizja {
			najnizja = meritev
		}
		if meritev > najvisja {
			najvisja = meritev
		}
		vsota += meritev
	}

	povprecje := vsota / float64(len(e.meritve))
	return najnizja, najvisja, povprecje, true
}

// IzpisMeritev na zaslon izpiše vse postaje in njihove zbrane meritve.
func IzpisMeritev(slovarPostaj map[string]Postaja) {
	fmt.Printf("MERITVE:\n")
	for oznaka, postaja := range slovarPostaj {
		fmt.Printf("%s - %s (%s): %v\n", oznaka, postaja.ime, postaja.kraj, postaja.meritve)
	}
}

// IzpisPorocila izpiše statistično poročilo in opisno stanje za vsako postajo,
// ki ima vsaj stMeritev meritev.
func IzpisPorocila(slovarPostaj map[string]Postaja, stMeritev int) {
	for oznaka, postaja := range slovarPostaj {
		najnizja, najvisja, povprecje, ok := statistika(slovarPostaj, oznaka, stMeritev)
		if !ok {
			fmt.Printf("%s: premalo meritev\n", postaja.ime)
			continue
		}
		var stanje string
		switch {
		case povprecje >= 25:
			stanje = "Vroče"
		case povprecje >= 15:
			stanje = "Toplo"
		case povprecje >= 5:
			stanje = "Hladno"
		default:
			stanje = "Mrzlo"
		}
		fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", postaja.ime, najnizja, najvisja, povprecje, stanje)
	}
}