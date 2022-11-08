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

We also see how the false positivity rate drops by inncreasing the number of hash functions
but after a certain limit it adversely affects the false positivity. Time taken to check
the existence increases with increasing the number of hash functions.

```
2022/11/08 12:31:21 total length of data 3864812 bytes
2022/11/08 12:31:21 number of words in corpus 333097
2022/11/08 12:31:21 number of words in test 37009
2022/11/08 12:31:21 regular set: len = 3864812 bytes, false postivity = 0.000000%
2022/11/08 12:31:21 bloom filter: len = 1024 bytes, frac size = 0.026495%, numHash = 2, false postivity = 100.000000%, time = 125.2767ms
2022/11/08 12:31:22 bloom filter: len = 2048 bytes, frac size = 0.052991%, numHash = 2, false postivity = 100.000000%, time = 89.0184ms
2022/11/08 12:31:22 bloom filter: len = 3072 bytes, frac size = 0.079486%, numHash = 2, false postivity = 100.000000%, time = 86.2646ms
2022/11/08 12:31:22 bloom filter: len = 4096 bytes, frac size = 0.105982%, numHash = 2, false postivity = 100.000000%, time = 95.9525ms
2022/11/08 12:31:22 bloom filter: len = 5120 bytes, frac size = 0.132477%, numHash = 2, false postivity = 100.000000%, time = 111.8181ms
2022/11/08 12:31:22 bloom filter: len = 10240 bytes, frac size = 0.264955%, numHash = 2, false postivity = 99.959469%, time = 84.5069ms
2022/11/08 12:31:22 bloom filter: len = 20480 bytes, frac size = 0.529909%, numHash = 2, false postivity = 96.619741%, time = 85.2401ms
2022/11/08 12:31:22 bloom filter: len = 30720 bytes, frac size = 0.794864%, numHash = 2, false postivity = 87.203113%, time = 150.0835ms
2022/11/08 12:31:22 bloom filter: len = 40960 bytes, frac size = 1.059819%, numHash = 2, false postivity = 75.349239%, time = 102.3893ms
2022/11/08 12:31:22 bloom filter: len = 51200 bytes, frac size = 1.324773%, numHash = 2, false postivity = 64.595098%, time = 92.3724ms
2022/11/08 12:31:22 bloom filter: len = 102400 bytes, frac size = 2.649547%, numHash = 2, false postivity = 30.957335%, time = 120.4588ms
2022/11/08 12:31:23 bloom filter: len = 153600 bytes, frac size = 3.974320%, numHash = 2, false postivity = 17.720014%, time = 150.2355ms
2022/11/08 12:31:23 bloom filter: len = 204800 bytes, frac size = 5.299093%, numHash = 2, false postivity = 11.489097%, time = 100.4356ms
2022/11/08 12:31:23 bloom filter: len = 256000 bytes, frac size = 6.623867%, numHash = 2, false postivity = 7.844038%, time = 99.6152ms
2022/11/08 12:31:23 bloom filter: len = 307200 bytes, frac size = 7.948640%, numHash = 2, false postivity = 5.652679%, time = 96.0896ms
2022/11/08 12:31:23 bloom filter: len = 358400 bytes, frac size = 9.273414%, numHash = 2, false postivity = 4.436759%, time = 105.9378ms
2022/11/08 12:31:23 bloom filter: len = 409600 bytes, frac size = 10.598187%, numHash = 2, false postivity = 3.528871%, time = 98.0646ms
2022/11/08 12:31:23 bloom filter: len = 460800 bytes, frac size = 11.922960%, numHash = 2, false postivity = 2.785809%, time = 183.4335ms
2022/11/08 12:31:23 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 2, false postivity = 2.437245%, time = 109.9409ms
2022/11/08 12:31:24 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 1, false postivity = 8.119647%, time = 103.919ms
2022/11/08 12:31:24 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 2, false postivity = 2.437245%, time = 128.8963ms
2022/11/08 12:31:24 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 3, false postivity = 1.075414%, time = 119.3191ms
2022/11/08 12:31:24 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 4, false postivity = 0.605258%, time = 127.6193ms
2022/11/08 12:31:24 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 5, false postivity = 0.413413%, time = 200.3413ms
2022/11/08 12:31:24 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 6, false postivity = 0.337756%, time = 214.3821ms
2022/11/08 12:31:25 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 7, false postivity = 0.289119%, time = 206.972ms
2022/11/08 12:31:25 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 8, false postivity = 0.259396%, time = 221.0815ms
2022/11/08 12:31:25 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 9, false postivity = 0.262098%, time = 221.0932ms
2022/11/08 12:31:25 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 10, false postivity = 0.272907%, time = 333.8154ms
2022/11/08 12:31:26 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 11, false postivity = 0.286417%, time = 318.8112ms
2022/11/08 12:31:26 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 12, false postivity = 0.329650%, time = 265.3814ms
2022/11/08 12:31:26 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 13, false postivity = 0.389095%, time = 278.5702ms
2022/11/08 12:31:26 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 14, false postivity = 0.478262%, time = 343.0103ms
2022/11/08 12:31:27 bloom filter: len = 512000 bytes, frac size = 13.247734%, numHash = 15, false postivity = 0.491772%, time = 300.6598ms
```
