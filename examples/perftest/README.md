Perf Test
===

We test the performance of bloom filters on word dictionary for various lookups.
The word dictionary used is same as the one used in profanity-detector example.

## How to run?

Download the [words_alpha.txt](https://github.com/dwyl/english-words) file and store it at location
`examples/profanity-detector/words.txt`.

```
$ go test -bench=. ./examples/perftest
```

## Results

```
goos: linux
goarch: amd64
pkg: github.com/arpitbbhayani/abloom/examples/perftest
cpu: AMD Ryzen 7 4800U with Radeon Graphics         
BenchmarkBFCheck-16                          381           3172999 ns/op
BenchmarkBFFirstCheck-16                20828637                51.24 ns/op
BenchmarkBFOneRandomCheck-16            11330626               112.2 ns/op
BenchmarkSetCheck-16                         230           4906682 ns/op
BenchmarkSetFirstCheck-16               82293306                14.12 ns/op
BenchmarkSetOneRandomCheck-16            6412029               178.6 ns/op
```
