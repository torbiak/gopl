package word

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
	"unicode"
)

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestRandomNonPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 500; i++ {
		np := randomNonPalindrome(rng)
		if IsPalindrome(np) {
			t.Errorf("IsPalindrome(%q) = true", np)
		}
		np = randomPunctuatedNonPalindrome(rng)
		if IsPalindrome(np) {
			t.Errorf("IsPalindrome(%q) = true", np)
		}
	}
}

// randomPalindrome returns a palindrome whose length and contents are derived
// from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

// Using a bunch of ideas from Eli Bendersky's excellent article:
// http://eli.thegreenplace.net/2010/01/28/generating-random-sentences-from-a-context-free-grammar/
//
// NON = acb | ab | aNONa | aNONb | aPALb
// PAL = eps | a | aa | aPALa
//
// - a, b, and c are any letter, but b has a different value than a if they
//   appear together in a production, and multiple a's in a production share
//   the same value.
// - eps is the empty string
//
// Unlike the grammar Bendersky was playing with, the non-palindrome grammar
// here converges very quickly unless we increase the weights of the
// non-terminals. Also, decaying the probability of productions frequently chosen in
// a recursive call tree is unnecessary. We don't, but if we wanted to we could
// probably get a consistent range of lengths by adding decay and tweaking the
// weights, so that it's very likely for nonterminals to be chosen but then
// decay at a certain rate.
var grammar = map[string][]weighted{
	"NON": []weighted{
		{"a c b", 1},
		{"a b", 1},
		{"a NON a", 30},
		{"a NON b", 30},
		{"a PAL b", 30},
	}, "PAL": []weighted{
		{"eps ", 1},
		{"a ", 1},
		{"a a", 1},
		{"a PAL a", 40},
	},
}
var letters []rune
var punctuation []rune
var punctProb = 0.1

type weighted struct {
	s      string
	weight float64
}

func randomNonPalindrome(rng *rand.Rand) string {
	return expand("NON", rng)
}

func randomPunctuatedNonPalindrome(rng *rand.Rand) string {
	b := &bytes.Buffer{}
	for _, r := range randomNonPalindrome(rng) {
		if rng.Float64() < punctProb {
			b.WriteRune(choosePunct(rng))
		}
		b.WriteRune(r)
	}
	return b.String()
}

func expand(symbol string, rng *rand.Rand) string {
	prod := choose(grammar[symbol], rng)
	buf := &bytes.Buffer{}

	var a rune
	for _, sym := range strings.Fields(prod) {
		if _, ok := grammar[sym]; ok { // recurse
			buf.WriteString(expand(sym, rng))
			continue
		}
		switch sym {
		case "a":
			if a == 0 {
				a = chooseLetter(rng)
			}
			buf.WriteRune(a)
		case "b":
			buf.WriteRune(chooseOtherLetter(a, rng))
		case "c":
			buf.WriteRune(chooseLetter(rng))
		case "eps":
			// nop: empty string
		default:
			panic(fmt.Sprintf("unexpected symbol %q", sym))
		}
	}
	return buf.String()
}

// choose returns a string from a slice, decaying the probability of those
// already-chosen.
func choose(choices []weighted, rng *rand.Rand) string {
	if len(choices) == 0 {
		panic("choose: no choices")
	}
	var sum float64
	for _, c := range choices {
		sum += c.weight
	}
	r := rng.Float64() * sum
	for _, c := range choices {
		r -= c.weight
		if r <= 0 {
			return c.s
		}
	}
	panic("choose: r was chosen incorrectly")
}

func chooseLetter(rng *rand.Rand) rune {
	return letters[rng.Intn(len(letters))]
}

// Choose a letter that isn't r or an upper/lowercase variant of r.
func chooseOtherLetter(r rune, rng *rand.Rand) rune {
	for {
		r2 := letters[rng.Intn(len(letters))]
		if unicode.ToLower(r2) == unicode.ToLower(r) {
			continue
		}
		return r2
	}
}

func choosePunct(rng *rand.Rand) rune {
	return punctuation[rng.Intn(len(punctuation))]
}

func init() {
	// Let's just stick with ASCII, since the odds of a font having some random
	// unicode codepoint seems pretty low.
	for r := rune(0x21); r < 0x7e; r++ {
		switch {
		case unicode.IsLetter(r):
			letters = append(letters, r)
		case unicode.IsPunct(r):
			punctuation = append(punctuation, r)
		}
	}
	// Visualizing the algorithm might be easier with a small character set:
	// letters = []rune{'a', 'b'}
}
