package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"

	"github.com/ylor/corepassword/cli"
)

var print = fmt.Println
var builder strings.Builder

var consonants = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "z"}
var vowels = []string{"a", "e", "i", "o", "u", "y"}

func randomSliceIndex(slice []string) int {
	return rand.IntN(len(slice))
}

func randomSliceElement(slice []string) string {
	slindex := rand.IntN(len(slice))
	return slice[slindex]
}

func main() {
	cli.HandleFlags()

	var phrases []string
	var rounds int = 2

	if cli.Preference_StrongPassword {
		rounds = 3
	}

	for i := 0; i < rounds; i++ {
		for j := 0; j < 2; j++ {
			builder.WriteString(randomSliceElement(consonants))
			builder.WriteString(randomSliceElement(vowels))
			builder.WriteString(randomSliceElement(consonants))
		}
		phrases = append(phrases, builder.String())
		builder.Reset()
	}

	capitalizedSextetIndex := randomSliceIndex(phrases)
	phrases[capitalizedSextetIndex] = strings.ToUpper(phrases[capitalizedSextetIndex][:1]) + phrases[capitalizedSextetIndex][1:]

	numberedSextetIndex := (randomSliceIndex(phrases))
	randomNumber := strconv.Itoa(rand.IntN(9) + 1)

	if numberedSextetIndex == 0 || numberedSextetIndex == len(phrases)-1 || rand.IntN(2) == 0 {
		phrases[numberedSextetIndex] += randomNumber
	} else {
		phrases[numberedSextetIndex] = randomNumber + phrases[numberedSextetIndex]
	}

	// print(phrases)
	print(strings.Join(phrases, "-"))
}
