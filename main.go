package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	KleineLetters    string
	GroteLetters     string
	Symbolen         string
	Nummers          string
	Aantal           int
	MinKleineLetters int
	MinGroteLetters  int
	MinSymbolen      int
	MinNummers       int
)

func init() {
	rand.Seed(time.Now().Unix())
	flag.StringVar(&KleineLetters, "WKL", "abcdedfghijklmnopqrstuvwxyz", "Welke kleine letters gebruikt worden")
	flag.StringVar(&GroteLetters, "WGL", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "Welke grote letters gebruikt worden")
	flag.StringVar(&Symbolen, "WS", "!@#$%&*", "Welke symbolen gebruikt worden")
	flag.StringVar(&Nummers, "WN", "0123456789", "Welke nummers gebruikt worden")
	flag.IntVar(&Aantal, "A", 8, "De lengte van het wachtwoord (aantal tekens)")
	flag.IntVar(&MinKleineLetters, "KL", 1, "Minimale hoeveelheid kleine letters")
	flag.IntVar(&MinGroteLetters, "GL", 1, "Minimale hoeveelheid grote letters")
	flag.IntVar(&MinSymbolen, "S", 1, "Minimale hoeveelheid symbolen")
	flag.IntVar(&MinNummers, "N", 1, "Minimale hoeveelheid nummers")
	flag.Parse()
}

func main() {

	GegenereerdWachtwoord := generator()
	fmt.Println(GegenereerdWachtwoord)
}

func generator() string {

	AlleTekens := KleineLetters + GroteLetters + Symbolen + Nummers
	var Wachtwoord strings.Builder

	for i := 0; i < MinKleineLetters; i++ {
		random := rand.Intn(len(KleineLetters))
		Wachtwoord.WriteString(string(KleineLetters[random]))
	}

	for i := 0; i < MinGroteLetters; i++ {
		random := rand.Intn(len(GroteLetters))
		Wachtwoord.WriteString(string(GroteLetters[random]))
	}

	for i := 0; i < MinSymbolen; i++ {
		random := rand.Intn(len(Symbolen))
		Wachtwoord.WriteString(string(Symbolen[random]))
	}

	for i := 0; i < MinNummers; i++ {
		random := rand.Intn(len(Nummers))
		Wachtwoord.WriteString(string(Nummers[random]))
	}

	Overgebleven := Aantal - MinKleineLetters - MinGroteLetters - MinSymbolen - MinNummers
	for i := 0; i < Overgebleven; i++ {
		random := rand.Intn(len(AlleTekens))
		Wachtwoord.WriteString(string(AlleTekens[random]))
	}

	Shuffle := strings.Split(Wachtwoord.String(), "")
	rand.Shuffle(len(Shuffle), func(i, j int) {
		Shuffle[i], Shuffle[j] = Shuffle[j], Shuffle[i]
	})
	Hussel := strings.Join(Shuffle, "")
	return Hussel
}
