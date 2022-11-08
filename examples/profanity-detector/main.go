package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/arpitbbhayani/abloom"
)

var corpusMap map[string]bool
var testMap map[string]bool
var sizeData int

func evalBF(size int, hashSeeds []int) {
	start := time.Now()
	var falsePos int = 0
	bf := abloom.NewBloom(size, hashSeeds)
	for word := range corpusMap {
		bf.Put([]byte(word))
	}
	for word := range testMap {
		if exists, _ := bf.Check([]byte(word)); exists {
			falsePos++
		}
	}

	log.Printf("bloom filter: len = %d bytes, frac size = %f%%, numHash = %d, false postivity = %f%%, time = %s\n",
		size, float64(size)/float64(sizeData)*100, len(hashSeeds), float64(falsePos)/float64(len(testMap))*100, time.Since(start))
}

func evalSet(size int) {
	var falsePos int = 0
	for word := range testMap {
		if _, ok := corpusMap[word]; ok {
			falsePos++
		}
	}
	log.Printf("regular set: len = %d bytes, false postivity = %f%%\n",
		size, float64(falsePos)/float64(len(testMap))*100)
}

func setup() {
	corpusMap = make(map[string]bool)
	testMap = make(map[string]bool)

	frac := 0.9

	fp, err := os.Open("./examples/profanity-detector/words.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}
	sizeData = len(data)

	words := bytes.Split(data, []byte("\n"))

	var i int = 0
	for i = range words {
		corpusMap[string(words[i])] = true
		if float64(i)/float64(len(words)) >= frac {
			break
		}
	}

	for i++; i < len(words); i++ {
		testMap[string(words[i])] = true
	}

	log.Println("total length of data", len(data), "bytes")
	log.Println("number of words in corpus", len(corpusMap))
	log.Println("number of words in test", len(testMap))
}

func main() {
	setup()

	var seeds []int = []int{
		938759194,
		214438831,
		242931268,
		319709952,
		766598923,
		206168488,
		486192512,
		587419578,
		573401536,
		471941957,
		338178239,
		341368536,
		610376822,
		542614060,
		677357443,
		733789834,
	}

	evalSet(sizeData)
	evalBF(1*1024, seeds[:2])
	evalBF(2*1024, seeds[:2])
	evalBF(3*1024, seeds[:2])
	evalBF(4*1024, seeds[:2])
	evalBF(5*1024, seeds[:2])
	evalBF(10*1024, seeds[:2])
	evalBF(20*1024, seeds[:2])
	evalBF(30*1024, seeds[:2])
	evalBF(40*1024, seeds[:2])
	evalBF(50*1024, seeds[:2])
	evalBF(100*1024, seeds[:2])
	evalBF(150*1024, seeds[:2])
	evalBF(200*1024, seeds[:2])
	evalBF(250*1024, seeds[:2])
	evalBF(300*1024, seeds[:2])
	evalBF(350*1024, seeds[:2])
	evalBF(400*1024, seeds[:2])
	evalBF(450*1024, seeds[:2])
	evalBF(500*1024, seeds[:2])

	for i := 1; i < len(seeds); i++ {
		evalBF(500*1024, seeds[:i])
	}
}
