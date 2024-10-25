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

func randomNum(max int) int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))

	if err != nil {
		panic(err)
	}
	return n.Int64()
}

func randomChar(chars string) string {
	i := randomNum(len(chars))
	return string(chars[i])
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

func addUpper(slice []string) []string {
	i := randomNum(len(slice))

	slice[i] = strings.ToUpper(slice[i][:1]) + slice[i][1:]
	return slice
}

func addNum(slice []string) []string {
	i := randomNum(len(slice))
	n := randomNum(9) + 1
	b := randomNum(2) == 0

	if i == 0 || i == int64(len(phrase)-1) || b {
		slice[i] += strconv.FormatInt(n, 10)
	} else {
		slice[i] = strconv.FormatInt(n, 10) + slice[i]
	}
	return slice
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

	phrase = addUpper(phrase)
	phrase = addNum(phrase)

	// fmt.Println(phrase)
	fmt.Println(strings.Join(phrase, "-"))
}
