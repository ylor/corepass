package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"

	"github.com/ylor/corepass/cli"
)

const (
	c = "bcdfghjklmnpqrstvwxz"
	v = "aeiouy"
)

var (
	syllables = 2
	words     = 2
	phrase    []string
)

func randomNum(max int64) int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	// fmt.Println(n)
	return n.Int64()
}

func randomIndex(slice []string) int64 {
	max := big.NewInt(int64(len(slice)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return n.Int64()
}

func randomChar(chars string) string {
	max := big.NewInt(int64(len(chars)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return string(chars[n.Int64()])
}

func generateWord(syllables int) string {
	var word string
	for range syllables {
		word += randomChar(c)
		word += randomChar(v)
		word += randomChar(c)
	}
	return word
}

func capitalizePhrase(p []string) []string {
	w := randomIndex(p)
	p[w] = strings.ToUpper(p[w][:1]) + p[w][1:]
	return p
}

func addNum(p []string) []string {
	numberedWordIndex := randomIndex(p)
	randomNumber := randomNum(9)

	if numberedWordIndex == 0 || numberedWordIndex == int64(len(phrase)-1) || randomNumber == 0 {
		p[numberedWordIndex] += strconv.FormatInt(randomNumber, 10)
	} else {
		p[numberedWordIndex] = strconv.FormatInt(randomNumber, 10) + p[numberedWordIndex]
	}

	return p
}

func main() {
	cli.HandleFlags()

	if cli.Preference_StrongPassword {
		words = 3
	}

	for range words {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			phrase = append(phrase, generateWord(syllables))
		}()
		wg.Wait()
	}

	phrase = capitalizePhrase(phrase)
	phrase = addNum(phrase)

	// fmt.Println(phrase)
	fmt.Println(strings.Join(phrase, "-"))
}
