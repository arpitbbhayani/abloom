package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/arpitbbhayani/abloom"
)

func evalBF(size int, sizeData int, corpus map[string]bool, test map[string]bool) {
	var falsePos int = 0
	bf := abloom.NewBloom(size, nil)
	for word := range corpus {
		bf.Put([]byte(word))
	}
	for word := range test {
		if exists, _ := bf.Check([]byte(word)); exists {
			falsePos++
		}
	}

	log.Printf("bloom filter: len = %d bytes, frac size = %f%%, false postivity = %f%%\n", size, float64(size)/float64(sizeData)*100, float64(falsePos)/float64(len(test))*100)
}

func evalSet(size int, sizeData int, corpus map[string]bool, test map[string]bool) {
	var falsePos int = 0
	for word := range test {
		if _, ok := corpus[word]; ok {
			falsePos++
		}
	}
	log.Printf("regular set: len = %d bytes, false postivity = %f%%\n", size, float64(falsePos)/float64(len(test))*100)
}

func main() {
	frac := 0.9

	fp, err := os.Open("./examples/profanity-detector/words.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}

	words := bytes.Split(data, []byte("\n"))

	var corpusMap map[string]bool = map[string]bool{}
	var testMap map[string]bool = map[string]bool{}
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

	evalSet(len(data), len(data), corpusMap, testMap)
	evalBF(1*1024, len(data), corpusMap, testMap)
	evalBF(2*1024, len(data), corpusMap, testMap)
	evalBF(3*1024, len(data), corpusMap, testMap)
	evalBF(4*1024, len(data), corpusMap, testMap)
	evalBF(5*1024, len(data), corpusMap, testMap)
	evalBF(10*1024, len(data), corpusMap, testMap)
	evalBF(20*1024, len(data), corpusMap, testMap)
	evalBF(30*1024, len(data), corpusMap, testMap)
	evalBF(40*1024, len(data), corpusMap, testMap)
	evalBF(50*1024, len(data), corpusMap, testMap)
	evalBF(100*1024, len(data), corpusMap, testMap)
	evalBF(150*1024, len(data), corpusMap, testMap)
	evalBF(200*1024, len(data), corpusMap, testMap)
	evalBF(250*1024, len(data), corpusMap, testMap)
	evalBF(300*1024, len(data), corpusMap, testMap)
	evalBF(350*1024, len(data), corpusMap, testMap)
	evalBF(400*1024, len(data), corpusMap, testMap)
	evalBF(450*1024, len(data), corpusMap, testMap)
	evalBF(500*1024, len(data), corpusMap, testMap)
}
