Bloom Paradox
===

As False Positivity rate of a Bloom Filter increases it makes no sense to check the BF.

## How to run?

Download the [words_alpha.txt](https://github.com/dwyl/english-words) file and store it at location
`examples/profanity-detector/words.txt`.

```
$ go run go run examples/bloom-paradox/main.go
$ pip install graph-cli
$ graph paradox-bench.csv --xcol 1 -o paradox-bench.png
```

## Results

![paradox-benchmark](https://user-images.githubusercontent.com/4745789/201531618-1a736e49-3927-4f86-b60e-4a810fa69d90.png)

```
2022/11/13 21:32:51 bf_size=10240,total_cost_get=7496841,total_cost_get_no_bf=7772218
2022/11/13 21:32:52 bf_size=9216,total_cost_get=7542168,total_cost_get_no_bf=7772218
2022/11/13 21:32:52 bf_size=8192,total_cost_get=7598970,total_cost_get_no_bf=7772218
2022/11/13 21:32:52 bf_size=7168,total_cost_get=7665855,total_cost_get_no_bf=7772218
2022/11/13 21:32:52 bf_size=6144,total_cost_get=7748121,total_cost_get_no_bf=7772218
2022/11/13 21:32:52 bf_size=5120,total_cost_get=7842351,total_cost_get_no_bf=7772218
2022/11/13 21:32:52 bf_size=4096,total_cost_get=7945986,total_cost_get_no_bf=7772218
2022/11/13 21:32:53 bf_size=3072,total_cost_get=8049123,total_cost_get_no_bf=7772218
2022/11/13 21:32:53 bf_size=2048,total_cost_get=8122641,total_cost_get_no_bf=7772218
2022/11/13 21:32:53 bf_size=1024,total_cost_get=8142138,total_cost_get_no_bf=7772218
2022/11/13 21:32:53 bf_size=512,total_cost_get=8142324,total_cost_get_no_bf=7772218
2022/11/13 21:32:53 bf_size=256,total_cost_get=8142324,total_cost_get_no_bf=7772218
2022/11/13 21:32:53 bf_size=128,total_cost_get=8142324,total_cost_get_no_bf=7772218
2022/11/13 21:32:53 bf_size=64,total_cost_get=8142324,total_cost_get_no_bf=7772218
```
