abloom
===

`abloom` is a pure Go implementation of Bloom Filter.

## Installation

```
$ go get github.com/arpitbbhayani/abloom
```

## How to use

A simple example of using bloom filter is as shown below. More examples can be found
in the [examples directory](https://github.com/arpitbbhayani/abloom/tree/master/examples/) of the package.

```go
package main

import (
	"log"

	"github.com/arpitbbhayani/abloom"
)

func main() {
	b := abloom.NewSimpleBF(512, nil)
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

## Behaviour

The behaviour of bloom filter on various parameters is evalated
on the usecase of profanity detector as documented in the
[examples directory](https://github.com/arpitbbhayani/abloom/tree/master/examples/profanity-detector). Please look at the source code to understand
what and how we have benchmarked.

### Bloom Filter Size vs False Positive Rate

We can see how the size of bloom filter is related to the observed false
positivity rate. Larger the size, smaller the false positivity rate.

![size-fp](https://user-images.githubusercontent.com/4745789/200518788-d545bc41-425b-47bf-a609-33b3a9ade34a.png)

### Number of Hash Functions vs False Positive Rate and Time

Here we see how the number of hash functions used in a bloom filter
affects the false positivity rate and also time. The FP rate decreases
initially but after a certain stage it saturates and then increases.

The time takes for the operation increases almost linearly with the
number of hash functions given how expensive hash computation can get.

![hash-fp](https://user-images.githubusercontent.com/4745789/200518773-76631419-a909-408e-9063-08a366218da2.png)
![hash-time](https://user-images.githubusercontent.com/4745789/200518783-835411e1-838e-4587-8b54-1de0acb54ca1.png)

### Deletable Bloom Filter Number of Regions vs False Positive Rate

![dbf-region](https://user-images.githubusercontent.com/4745789/200841018-dffcbeb3-d6be-43f8-a8da-e29ef28698a2.png)

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
- Rothenberg, Christian Esteve, Carlos A. B. Macapuna, Fabio L. Verdi, and Mauricio F. Magalhaes. “The Deletable Bloom Filter: A New Member of the Bloom Family.” arXiv, May 3, 2010. http://arxiv.org/abs/1005.0352.

## Contributors

<a href = "https://github.com/arpitbbhayani/abloom/graphs/contributors">
  <img src = "https://contrib.rocks/image?repo=arpitbbhayani/abloom"/>
</a>

## License

Abloom is open-sourced under [Apache License, Version 2.0](LICENSE.md).
