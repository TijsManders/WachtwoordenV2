package main

import (
	"math/rand"
	"strings"
	"time"
)

var (
	KleineLetters    = "abcdedfghijklmnopqrst"
	GroteLetters     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Symbolen         = "!@#$%&*"
	Nummers          = "0123456789"
	AlleTekens       = KleineLetters + GroteLetters + Symbolen + Nummers
	Lengte           = 8
	MinKleineLetters = 1
	MinGroteLetters  = 1
	MinimaleSymbolen = 1
	MinNummers       = 1
)

func generator() string {
	rand.Seed(time.Now().Unix())

	var Wachtwoord strings.Builder

	for i := 0; i < MinKleineLetters; i++ {
		random := rand.Intn(len(KleineLetters))
		Wachtwoord.WriteString(string(KleineLetters[random]))
	}

	for i := 0; i < MinGroteLetters; i++ {
		random := rand.Intn(len(GroteLetters))
		Wachtwoord.WriteString(string(GroteLetters[random]))
	}

	for i := 0; i < MinimaleSymbolen; i++ {
		random := rand.Intn(len(Symbolen))
		Wachtwoord.WriteString(string(Symbolen[random]))
	}

	for i := 0; i < MinNummers; i++ {
		random := rand.Intn(len(Nummers))
		Wachtwoord.WriteString(string(Nummers[random]))
	}

	Overgebleven := Lengte - MinKleineLetters - MinGroteLetters - MinimaleSymbolen - MinNummers
	for i := 0; i < Overgebleven; i++ {
		random := rand.Intn(len(AlleTekens))
		Wachtwoord.WriteString(string(AlleTekens[random]))
	}

	return Wachtwoord.String()
}
