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
$ pip install graph-cli
$ graph size-bench.csv --xcol 2 --ycol 4 -o size-fp.png
$ graph hashfn-bench.csv --xcol 3 --ycol 4 -o hash-fp.png
$ graph hashfn-bench.csv --xcol 3 --ycol 5 -o hash-time.png
$ graph dbf-region-bench.csv --xcol 6 --ycol 4 -o dbf-region.png
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

### Bloom Filter Size vs False Positive Rate

We can see how the size of bloom filter is related to the observed false
positivity rate. Larger the size, smaller the false positivity rate.

![size-fp](https://user-images.githubusercontent.com/4745789/200518788-d545bc41-425b-47bf-a609-33b3a9ade34a.png)

```
2022/11/08 14:11:15 total length of data 3864812 bytes
2022/11/08 14:11:15 number of words in corpus 333097
2022/11/08 14:11:15 number of words in test 37009
2022/11/08 14:11:15 size=1024,size_frac=0.026495,num_hash=2,fp_rate=100.000000,time_taken_ms=98
2022/11/08 14:11:15 size=2048,size_frac=0.052991,num_hash=2,fp_rate=100.000000,time_taken_ms=89
2022/11/08 14:11:15 size=3072,size_frac=0.079486,num_hash=2,fp_rate=100.000000,time_taken_ms=79
2022/11/08 14:11:15 size=4096,size_frac=0.105982,num_hash=2,fp_rate=100.000000,time_taken_ms=92
2022/11/08 14:11:15 size=5120,size_frac=0.132477,num_hash=2,fp_rate=100.000000,time_taken_ms=101
2022/11/08 14:11:15 size=10240,size_frac=0.264955,num_hash=2,fp_rate=99.959469,time_taken_ms=80
2022/11/08 14:11:15 size=20480,size_frac=0.529909,num_hash=2,fp_rate=96.619741,time_taken_ms=116
2022/11/08 14:11:15 size=30720,size_frac=0.794864,num_hash=2,fp_rate=87.203113,time_taken_ms=97
2022/11/08 14:11:16 size=40960,size_frac=1.059819,num_hash=2,fp_rate=75.349239,time_taken_ms=89
2022/11/08 14:11:16 size=51200,size_frac=1.324773,num_hash=2,fp_rate=64.595098,time_taken_ms=103
2022/11/08 14:11:16 size=102400,size_frac=2.649547,num_hash=2,fp_rate=30.957335,time_taken_ms=100
2022/11/08 14:11:16 size=153600,size_frac=3.974320,num_hash=2,fp_rate=17.720014,time_taken_ms=89
2022/11/08 14:11:16 size=204800,size_frac=5.299093,num_hash=2,fp_rate=11.489097,time_taken_ms=106
2022/11/08 14:11:16 size=256000,size_frac=6.623867,num_hash=2,fp_rate=7.844038,time_taken_ms=104
2022/11/08 14:11:16 size=307200,size_frac=7.948640,num_hash=2,fp_rate=5.652679,time_taken_ms=92
2022/11/08 14:11:16 size=358400,size_frac=9.273414,num_hash=2,fp_rate=4.436759,time_taken_ms=99
2022/11/08 14:11:16 size=409600,size_frac=10.598187,num_hash=2,fp_rate=3.528871,time_taken_ms=110
2022/11/08 14:11:16 size=358400,size_frac=9.273414,num_hash=2,fp_rate=4.436759,time_taken_ms=88
2022/11/08 14:11:17 size=512000,size_frac=13.247734,num_hash=2,fp_rate=2.437245,time_taken_ms=87
```

### Number of Hash Functions vs False Positive Rate and Time

Here we see how the number of hash functions used in a bloom filter
affects the false positivity rate and also time. The FP rate decreases
initially but after a certain stage it saturates and then increases.

The time takes for the operation increases almost linearly with the
number of hash functions given how expensive hash computation can get.

![hash-fp](https://user-images.githubusercontent.com/4745789/200518773-76631419-a909-408e-9063-08a366218da2.png)
![hash-time](https://user-images.githubusercontent.com/4745789/200518783-835411e1-838e-4587-8b54-1de0acb54ca1.png)

```
2022/11/08 14:11:15 total length of data 3864812 bytes
2022/11/08 14:11:15 number of words in corpus 333097
2022/11/08 14:11:15 number of words in test 37009
2022/11/08 14:11:17 size=512000,size_frac=13.247734,num_hash=1,fp_rate=8.119647,time_taken_ms=97
2022/11/08 14:11:17 size=512000,size_frac=13.247734,num_hash=2,fp_rate=2.437245,time_taken_ms=120
2022/11/08 14:11:17 size=512000,size_frac=13.247734,num_hash=3,fp_rate=1.075414,time_taken_ms=106
2022/11/08 14:11:17 size=512000,size_frac=13.247734,num_hash=4,fp_rate=0.605258,time_taken_ms=114
2022/11/08 14:11:17 size=512000,size_frac=13.247734,num_hash=5,fp_rate=0.413413,time_taken_ms=142
2022/11/08 14:11:17 size=512000,size_frac=13.247734,num_hash=6,fp_rate=0.337756,time_taken_ms=158
2022/11/08 14:11:17 size=512000,size_frac=13.247734,num_hash=7,fp_rate=0.289119,time_taken_ms=214
2022/11/08 14:11:18 size=512000,size_frac=13.247734,num_hash=8,fp_rate=0.259396,time_taken_ms=235
2022/11/08 14:11:18 size=512000,size_frac=13.247734,num_hash=9,fp_rate=0.262098,time_taken_ms=192
2022/11/08 14:11:18 size=512000,size_frac=13.247734,num_hash=10,fp_rate=0.272907,time_taken_ms=242
2022/11/08 14:11:18 size=512000,size_frac=13.247734,num_hash=11,fp_rate=0.286417,time_taken_ms=237
2022/11/08 14:11:19 size=512000,size_frac=13.247734,num_hash=12,fp_rate=0.329650,time_taken_ms=236
2022/11/08 14:11:19 size=512000,size_frac=13.247734,num_hash=13,fp_rate=0.389095,time_taken_ms=286
2022/11/08 14:11:19 size=512000,size_frac=13.247734,num_hash=14,fp_rate=0.478262,time_taken_ms=318
2022/11/08 14:11:19 size=512000,size_frac=13.247734,num_hash=15,fp_rate=0.491772,time_taken_ms=288
```
