Profanity Detector
===

We use this example to test the efficiency of Bloom Filter on Word Dictionary. We used
the [words_alpha.txt](https://github.com/dwyl/english-words) from the repository
[dwyl/english-words](https://github.com/dwyl/english-words) that holds 370105 words.

We split the words into two segments in the ratio 9:1.

1. corpus - these would be correct english words
2. test - these would be words that are considered as profane (for benchmarking)

The total size of the dataset is 3864811 bytes ~ 3.8 MB. Corpus contains 333096 words
while the test (profane list) contains 37009 words.

For set (map) and bloom filters of various length, we compute and analyze the false positivity rate.

## How to run?

Download the [words_alpha.txt](https://github.com/dwyl/english-words) file and store it at location
`examples/profanity-detector/words.txt`.

```
$ go run examples/profanity-detector/main.go
```

## Results

The following result shows

1. storage - regular set or bloom filter
2. len - size occupied by the set or bloom filter
3. frac size - fraction of the size occupied as compared to the size of the data
4. false positivity rate on the test (profane) list.

We can see, how bloom filter's false positivity rate drops significantly with the
increase in the length of the filter. At the bloom filter length of 512 KB
which is mere 13.25% of the data size, we get a false positivity rate of 2.33%.

```
2022/11/08 01:17:34 regular set: len = 3864811 bytes, false postivity = 0.000000%
2022/11/08 01:17:34 bloom filter: len = 1024 bytes, frac size = 0.026495%, false postivity = 100.000000%
2022/11/08 01:17:34 bloom filter: len = 2048 bytes, frac size = 0.052991%, false postivity = 100.000000%
2022/11/08 01:17:34 bloom filter: len = 3072 bytes, frac size = 0.079486%, false postivity = 100.000000%
2022/11/08 01:17:34 bloom filter: len = 4096 bytes, frac size = 0.105982%, false postivity = 100.000000%
2022/11/08 01:17:34 bloom filter: len = 5120 bytes, frac size = 0.132477%, false postivity = 100.000000%
2022/11/08 01:17:34 bloom filter: len = 10240 bytes, frac size = 0.264955%, false postivity = 99.967575%
2022/11/08 01:17:34 bloom filter: len = 20480 bytes, frac size = 0.529909%, false postivity = 96.630549%
2022/11/08 01:17:34 bloom filter: len = 30720 bytes, frac size = 0.794864%, false postivity = 87.113945%
2022/11/08 01:17:34 bloom filter: len = 40960 bytes, frac size = 1.059819%, false postivity = 75.808587%
2022/11/08 01:17:34 bloom filter: len = 51200 bytes, frac size = 1.324774%, false postivity = 64.397849%
2022/11/08 01:17:34 bloom filter: len = 102400 bytes, frac size = 2.649547%, false postivity = 30.984355%
2022/11/08 01:17:35 bloom filter: len = 153600 bytes, frac size = 3.974321%, false postivity = 17.698398%
2022/11/08 01:17:35 bloom filter: len = 204800 bytes, frac size = 5.299095%, false postivity = 11.159448%
2022/11/08 01:17:35 bloom filter: len = 256000 bytes, frac size = 6.623869%, false postivity = 7.606258%
2022/11/08 01:17:35 bloom filter: len = 307200 bytes, frac size = 7.948642%, false postivity = 5.704018%
2022/11/08 01:17:35 bloom filter: len = 358400 bytes, frac size = 9.273416%, false postivity = 4.223297%
2022/11/08 01:17:35 bloom filter: len = 409600 bytes, frac size = 10.598190%, false postivity = 3.380259%
2022/11/08 01:17:35 bloom filter: len = 460800 bytes, frac size = 11.922963%, false postivity = 2.780405%
2022/11/08 01:17:35 bloom filter: len = 512000 bytes, frac size = 13.247737%, false postivity = 2.331865%
```
