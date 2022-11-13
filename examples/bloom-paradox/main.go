package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/arpitbbhayani/abloom"
)

var allWords []string
var wordsMap map[string]bool
var cacheMap map[string]bool
var bf *abloom.SimpleBF

const COST_BF int = 1
const COST_CACHE int = 3
const COST_MEMORY int = 20

type datapoint struct {
	bfSize           int
	totalCostGet     int
	totalCostGetNoBF int
}

func (d *datapoint) toStringSlice() []string {
	return []string{
		strconv.FormatInt(int64(d.bfSize), 10),
		strconv.FormatInt(int64(d.totalCostGet), 10),
		strconv.FormatInt(int64(d.totalCostGetNoBF), 10),
	}
}

func (d *datapoint) String() string {
	return fmt.Sprintf("bf_size=%d,total_cost_get=%d,total_cost_get_no_bf=%d", d.bfSize, d.totalCostGet, d.totalCostGetNoBF)
}

func init() {
	allWords = make([]string, 0)
	wordsMap = make(map[string]bool)
	fp, err := os.Open("./examples/profanity-detector/words.txt")
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
		wordsMap[string(words[i])] = true
		allWords = append(allWords, string(words[i]))
	}
}

func Get(k string, checkBF bool) int {
	cost := 0

	if checkBF {
		// if BF says key does not exist
		// then key definitely does not exist in cache
		// hence we have to hit the memory and get the key
		cost += COST_BF
		if exists, _ := bf.Check([]byte(k)); !exists {
			cost += COST_MEMORY
			return cost
		}
	}

	// bloom filter said yes, hence we check the cache
	// if cache says yes ... great serve
	// if no then hit the memory
	cost += COST_CACHE
	if _, ok := cacheMap[k]; ok {
		return cost
	} else {
		cost += COST_MEMORY
		return cost
	}
}

func evalParadox(sizeBF int) *datapoint {
	cacheMap = make(map[string]bool)
	bf = abloom.NewSimpleBF(sizeBF, []int{312312, 43242})

	rand.Shuffle(len(allWords), func(i, j int) { allWords[i], allWords[j] = allWords[j], allWords[i] })

	// populate the cache with 10% of the words
	// populate the BF with the words put in the cache
	for i := range allWords {
		cacheMap[allWords[i]] = true
		bf.Put([]byte(allWords[i]))
		if i == len(allWords)/10 {
			break
		}
	}

	rand.Shuffle(len(allWords), func(i, j int) { allWords[i], allWords[j] = allWords[j], allWords[i] })

	var totalCostGet int = 0
	for i := range allWords {
		word := allWords[i]
		totalCostGet += Get(word, true)
	}

	var totalCostGetNoBF int = 0
	for i := range allWords {
		word := allWords[i]
		totalCostGetNoBF += Get(word, false)
	}

	dpoint := &datapoint{
		bfSize:           sizeBF,
		totalCostGet:     totalCostGet,
		totalCostGetNoBF: totalCostGetNoBF,
	}
	log.Println(dpoint)
	return dpoint
}

func flushDatapoints(name string, datapoints []*datapoint) {
	fp, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(fp)
	w.Write([]string{"bf size", "get with bf", "get no bf"})
	for i := range datapoints {
		w.Write(datapoints[i].toStringSlice())
	}
	w.Flush()
}

func main() {
	var datapoints []*datapoint
	for _, size := range []int{10 * 1024, 9 * 1024, 8 * 1024, 7 * 1024, 6 * 1024, 5 * 1024, 4 * 1024, 3 * 1024, 2 * 1024, 1 * 1024, 512, 256, 128, 64} {
		datapoints = append(datapoints, evalParadox(size))
	}
	flushDatapoints("paradox-bench.csv", datapoints)
}
