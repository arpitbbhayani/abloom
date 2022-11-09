package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/arpitbbhayani/abloom"
)

var corpusMap map[string]bool
var testMap map[string]bool
var sizeData int

type datapoint struct {
	size        int
	sizeFrac    float64
	numHash     int
	fpRate      float64
	timeTakenMs int64
}

func (d *datapoint) toStringSlice() []string {
	return []string{
		strconv.FormatInt(int64(d.size), 10),
		strconv.FormatFloat(d.sizeFrac, 'f', 2, 64),
		strconv.FormatInt(int64(d.numHash), 10),
		strconv.FormatFloat(d.fpRate, 'f', 2, 64),
		strconv.FormatInt(int64(d.timeTakenMs), 10),
	}
}

func (d *datapoint) String() string {
	return fmt.Sprintf("size=%d,size_frac=%f,num_hash=%d,fp_rate=%f,time_taken_ms=%d",
		d.size, d.sizeFrac, d.numHash, d.fpRate, d.timeTakenMs)
}

func evalBF(size int, hashSeeds []int) *datapoint {
	start := time.Now()
	var falsePos int = 0
	bf := abloom.NewSimpleBF(size, hashSeeds)
	for word := range corpusMap {
		bf.Put([]byte(word))
	}
	for word := range testMap {
		if exists, _ := bf.Check([]byte(word)); exists {
			falsePos++
		}
	}

	return &datapoint{
		size:        size,
		sizeFrac:    float64(size) / float64(sizeData) * 100,
		numHash:     len(hashSeeds),
		fpRate:      float64(falsePos) / float64(len(testMap)) * 100,
		timeTakenMs: time.Since(start).Milliseconds(),
	}
}

func evalSet(size int) {
	var falsePos int = 0
	for word := range testMap {
		if _, ok := corpusMap[word]; ok {
			falsePos++
		}
	}
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

func flushDatapoints(name string, datapoints []*datapoint) {
	fp, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(fp)
	w.Write([]string{"size", "size frac", "num hash", "fp rate", "time taken"})
	for i := range datapoints {
		w.Write(datapoints[i].toStringSlice())
	}
	w.Flush()
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

	var sizeDatapoints []*datapoint
	for _, sizeKB := range []int{1, 2, 3, 4, 5, 10, 20, 30, 40, 50, 100, 150, 200, 250, 300, 350, 400, 350, 500} {
		dp := evalBF(sizeKB*1024, seeds[:2])
		sizeDatapoints = append(sizeDatapoints, dp)
		log.Println(dp)
	}

	var hashFnDatapoints []*datapoint
	for i := 1; i < len(seeds); i++ {
		dp := evalBF(500*1024, seeds[:i])
		hashFnDatapoints = append(hashFnDatapoints, dp)
		log.Println(dp)
	}

	flushDatapoints("size-bench.csv", sizeDatapoints)
	flushDatapoints("hashfn-bench.csv", hashFnDatapoints)
}
