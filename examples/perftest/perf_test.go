package perftest_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/arpitbbhayani/abloom"
)

var corpusMap map[string]bool
var testWordsB [][]byte
var testWordsS []string
var bf *abloom.Bloom

func init() {
	corpusMap = make(map[string]bool)
	testWordsB = make([][]byte, 0)
	testWordsS = make([]string, 0)
	bf = abloom.NewBloom(500 * 1024)
}

func checkBF(word []byte) bool {
	if exists, _ := bf.Check(word); exists {
		return true
	}
	return false
}

func checkSet(word string) bool {
	if _, ok := corpusMap[word]; ok {
		return true
	}
	return false
}

func setup() {
	frac := 0.9

	fp, err := os.Open("../profanity-detector/words.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}

	words := bytes.Split(data, []byte("\n"))

	var i int = 0
	for i = range words {
		corpusMap[string(words[i])] = true
		bf.Put([]byte(words[i]))
		if float64(i)/float64(len(words)) >= frac {
			break
		}
	}

	for i++; i < len(words); i++ {
		testWordsS = append(testWordsS, string(words[i]))
		testWordsB = append(testWordsB, words[i])
	}
}

func BenchmarkBFCheck(b *testing.B) {
	setup()
	for i := 0; i < b.N; i++ {
		for _, word := range testWordsB {
			checkBF(word)
		}
	}
}

func BenchmarkBFOneCheck(b *testing.B) {
	setup()
	for i := 0; i < b.N; i++ {
		checkBF(testWordsB[0])
	}
}

func BenchmarkBFOneRandomCheck(b *testing.B) {
	setup()
	for i := 0; i < b.N; i++ {
		idx := rand.Int() % len(testWordsB)
		checkBF(testWordsB[idx])
	}
}

func BenchmarkSetCheck(b *testing.B) {
	setup()
	for i := 0; i < b.N; i++ {
		for _, word := range testWordsS {
			checkSet(word)
		}
	}
}

func BenchmarkSetOneCheck(b *testing.B) {
	setup()
	for i := 0; i < b.N; i++ {
		checkSet(testWordsS[0])
	}
}

func BenchmarkSetOneRandomCheck(b *testing.B) {
	setup()
	for i := 0; i < b.N; i++ {
		idx := rand.Int() % len(testWordsS)
		checkSet(testWordsS[idx])
	}
}
