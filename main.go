package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	// voor de generator
	KleineLetters    string
	GroteLetters     string
	Symbolen         string
	Nummers          string
	Aantal           int
	MinKleineLetters int
	MinGroteLetters  int
	MinSymbolen      int
	MinNummers       int
	// voor de checker
	Wachtwoord         string
	ScoreLengte        int
	ScoreGroteLetters  int
	ScoreKleineLetters int
	ScoreNummers       int
	ScoreSymbolen      int
	HGL                int
	HKL                int
	HN                 int
	HS                 int
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
	flag.StringVar(&Wachtwoord, "WW", "", "Vul wachtwoord in dat gecheckt mag worden.")
	flag.Parse()
}

func main() {

	if Wachtwoord == "" {
		GegenereerdWachtwoord := generator()
		fmt.Println(GegenereerdWachtwoord)
	} else {
		GecontroleerdWachtwoord := checker()
		fmt.Println("De score van het wachtwoord kan variëren van 1 tot 10. Waarbij 1 slecht is en 10 goed. De score van je opgegeven wachtwoord is", GecontroleerdWachtwoord)
	}

}

func checker() string {
	OpgedeeldWW := []rune(Wachtwoord)
	// Score voor de lengte van het ww
	if len(OpgedeeldWW) < 8 {
		ScoreLengte = 0
	}
	if len(OpgedeeldWW) >= 8 && len(OpgedeeldWW) < 12 {
		ScoreLengte = 1
	}
	if len(OpgedeeldWW) >= 12 {
		ScoreLengte = 2
	}
	// Score voor de hoofdletters in ww
	for _, i := range OpgedeeldWW {
		if unicode.IsUpper(i) {
			HGL++
		}
	}
	if HGL < 1 {
		ScoreGroteLetters = 0
	}
	if HGL < 3 && HGL >= 1 {
		ScoreGroteLetters = 1
	}
	if HGL >= 3 {
		ScoreGroteLetters = 2
	}
	// Score voor de kleine letters in ww
	for _, i := range OpgedeeldWW {
		if unicode.IsLower(i) {
			HKL++
		}
	}
	if HKL < 1 {
		ScoreKleineLetters = 0
	}
	if HKL < 3 && HKL >= 1 {
		ScoreKleineLetters = 1
	}
	if HKL >= 3 {
		ScoreKleineLetters = 2
	}
	// Score voor de nummers in ww
	for _, i := range OpgedeeldWW {
		if unicode.IsNumber(i) {
			HN++
		}
	}
	if HN < 1 {
		ScoreNummers = 0
	}
	if HN < 3 && HN >= 1 {
		ScoreNummers = 1
	}
	if HN >= 3 {
		ScoreNummers = 2
	}
	// Score voor de symbolen in ww
	for _, i := range OpgedeeldWW {
		if unicode.IsSymbol(i) {
			HS++
		}
	}
	if HS < 1 {
		ScoreSymbolen = 0
	}
	if HS < 3 && HS >= 1 {
		ScoreSymbolen = 1
	}
	if HS >= 3 {
		ScoreSymbolen = 2
	}
	// Berekening totaalscore
	TotaalScore := ScoreLengte + ScoreGroteLetters + ScoreKleineLetters + ScoreNummers + ScoreSymbolen
	return strconv.Itoa(TotaalScore)
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
