abloom
===

`abloom` is a pure Go implementation of Bloom Filter.

## Installation

```
$ go get github.com/arpitbbhayani/abloom
```

## How to use

A simple example of using bloom filter is as shown below. More examples can be found
in the [examples directory](https://github.com/arpitbbhayani/abloom/master/examples) of the package.

```go
package examples

import (
	"log"

	"github.com/arpitbbhayani/abloom"
)

func SimpleBloom() {
	b := abloom.NewBloom(512)
	b.Put([]byte("apple"))
	b.Put([]byte("banana"))
	b.Put([]byte("cat"))

	v, err := b.Check([]byte("apple"))
	if err != nil {
		log.Fatal("error while computing the hash")
	}
	log.Println("is apple present?", v)
}
```

## Running Tests

To run all the tests fire the following command

```sh
$ go test ./...
```

## Benchmark

Preliminary benchmark results on [profanity usecase](https://github.com/arpitbbhayani/abloom/tree/master/examples/perftest) are shown below. Please refer to the perf test code to see what has been
tested and benchmarked.

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

## Citations

- Bloom, Burton H. “Space/Time Trade-Offs in Hash Coding with Allowable Errors.” Communications of the ACM 13, no. 7 (July 1970): 422–26. https://doi.org/10.1145/362686.362692.

## Contributors

<a href = "https://github.com/arpitbbhayani/abloom/graphs/contributors">
  <img src = "https://contrib.rocks/image?repo=arpitbbhayani/abloom"/>
</a>

## License

DiceDB is open-sourced under [Apache License, Version 2.0](LICENSE.md).
