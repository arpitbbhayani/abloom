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
BenchmarkBFCheck-16                           67          17266452 ns/op
BenchmarkBFOneCheck-16                  14426850                72.92 ns/op
BenchmarkBFOneRandomCheck-16             3340220               301.2 ns/op
BenchmarkSetCheck-16                          13          80918323 ns/op
BenchmarkSetOneCheck-16                 102268398               11.38 ns/op
BenchmarkSetOneRandomCheck-16            4730517               246.0 ns/op
```
