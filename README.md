# Go Dojo

This is a [dojo](https://en.wikipedia.org/wiki/Dojo) to experiment with and learn from [Go](https://go.dev/).

# Snippets

Here I might add little code snippets as I experiment with things.

# Benchmarks

Run benchmarks with `make bench`. These benchmarks will run a series of tests to answer the question "What is the fastest way to do X in Go?"

## Results

Last update: `Feb 2nd, 2025`
Go version: `1.23.5`

```
Running tests...
go test -run=Test. -race ./bench
ok      github.com/stefanovazzocell/GoDojo/bench        1.008s
Running benchmarks...
go test -run=Bench -bench=. -benchmem ./bench
goos: linux
goarch: amd64
pkg: github.com/stefanovazzocell/GoDojo/bench
cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
BenchmarkControlFlowSimple/switch-12            1000000000               0.5270 ns/op          0 B/op          0 allocs/op
BenchmarkControlFlowSimple/if-12                1000000000               0.5025 ns/op          0 B/op          0 allocs/op
BenchmarkControlFlowSimple/ifElse-12            1000000000               0.5102 ns/op          0 B/op          0 allocs/op
BenchmarkControlFlowComplex/switch-12           1000000000               1.011 ns/op           0 B/op          0 allocs/op
BenchmarkControlFlowComplex/if-12               1000000000               0.7849 ns/op          0 B/op          0 allocs/op
BenchmarkControlFlowComplex/ifElse-12           1000000000               0.7824 ns/op          0 B/op          0 allocs/op
BenchmarkSliceLoopValue/rangeKey-12                 5439            215526 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopValue/rangeKeyReuse-12            5422            223808 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopValue/rangeVal-12                 4647            218455 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopValue/rangeValReuse-12            5444            214505 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopValue/len-12                      5497            221481 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopValue/lenOptimizedA-12            5470            220183 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopValue/lenOptimizedB-12            4891            215579 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopValue/lenRange-12                 4935            220943 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopValue/lenRangeReuse-12            4644            222280 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopCount/rangeKey-12                  957           1167477 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopCount/rangeKeyReuse-12             961           1207362 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopCount/rangeVal-12                  963           1171690 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopCount/rangeValReuse-12             960           1212378 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopCount/rangeBlank-12                970           1196525 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopCount/len-12                       976           1196759 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopCount/lenOptimizedA-12             964           1207626 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopCount/lenOptimizedB-12             951           1213748 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopCount/lenRange-12                  968           1199078 ns/op               0 B/op          0 allocs/op
BenchmarkIntToString/Sprintf-12                 15172555                73.30 ns/op           16 B/op          1 allocs/op
BenchmarkIntToString/strconvItoa-12             43874653                23.95 ns/op            7 B/op          0 allocs/op
BenchmarkIntToString/strconvFormatInt-12        44477778                23.84 ns/op            7 B/op          0 allocs/op
BenchmarkLocks/Mutex/Sequential-12              88732797                12.23 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/Mutex/Parallel-12                18501244                61.84 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/RWMutex/Sequential-12            51084265                21.79 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/RWMutex/Parallel-12              20775333                56.63 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/RWMutexRead/Sequential-12        88384423                12.25 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/RWMutexRead/Parallel-12          35315664                34.20 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/Bool/Sequential-12               94611458                11.78 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/Bool/Parallel-12                  4873657               253.4 ns/op             0 B/op          0 allocs/op
BenchmarkLocks/int32/Sequential-12              97941865                10.97 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/int32/Parallel-12                 4621382               257.4 ns/op             0 B/op          0 allocs/op
BenchmarkLocks/int64/Sequential-12              96177319                11.85 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/int64/Parallel-12                 5449341               206.9 ns/op             0 B/op          0 allocs/op
BenchmarkLocks/Uint32/Sequential-12             87853814                12.05 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/Uint32/Parallel-12                5573180               204.7 ns/op             0 B/op          0 allocs/op
BenchmarkLocks/Uint64/Sequential-12             95775746                11.57 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/Uint64/Parallel-12                5597190               232.2 ns/op             0 B/op          0 allocs/op
BenchmarkLocksNoWork/Mutex/Sequential-12        126280051                9.486 ns/op           0 B/op          0 allocs/op
BenchmarkLocksNoWork/Mutex/Parallel-12          18600835                63.94 ns/op            0 B/op          0 allocs/op
BenchmarkLocksNoWork/RWMutex/Sequential-12      55858296                19.94 ns/op            0 B/op          0 allocs/op
BenchmarkLocksNoWork/RWMutex/Parallel-12        17739396                66.67 ns/op            0 B/op          0 allocs/op
BenchmarkLocksNoWork/RWMutexRead/Sequential-12          128159222                9.506 ns/op           0 B/op          0 allocs/op
BenchmarkLocksNoWork/RWMutexRead/Parallel-12            41687808                29.14 ns/op            0 B/op          0 allocs/op
BenchmarkLocksNoWork/Bool/Sequential-12                 96374092                11.09 ns/op            0 B/op          0 allocs/op
BenchmarkLocksNoWork/Bool/Parallel-12                    6397933               191.9 ns/op             0 B/op          0 allocs/op
BenchmarkLocksNoWork/int32/Sequential-12                142250307                8.541 ns/op           0 B/op          0 allocs/op
BenchmarkLocksNoWork/int32/Parallel-12                   6055423               209.3 ns/op             0 B/op          0 allocs/op
BenchmarkLocksNoWork/int64/Sequential-12                125472151                9.256 ns/op           0 B/op          0 allocs/op
BenchmarkLocksNoWork/int64/Parallel-12                   4915215               246.8 ns/op             0 B/op          0 allocs/op
BenchmarkLocksNoWork/Uint32/Sequential-12               143362797                8.702 ns/op           0 B/op          0 allocs/op
BenchmarkLocksNoWork/Uint32/Parallel-12                  4548354               252.6 ns/op             0 B/op          0 allocs/op
BenchmarkLocksNoWork/Uint64/Sequential-12               126149358                9.511 ns/op           0 B/op          0 allocs/op
BenchmarkLocksNoWork/Uint64/Parallel-12                  6141579               187.5 ns/op             0 B/op          0 allocs/op
BenchmarkStringCharLoop/len10/rangeKey-12               153000598                7.771 ns/op           0 B/op          0 allocs/op
BenchmarkStringCharLoop/len10/rangeVal-12               234523108                5.057 ns/op           0 B/op          0 allocs/op
BenchmarkStringCharLoop/len10/len-12                    354179029                3.362 ns/op           0 B/op          0 allocs/op
BenchmarkStringCharLoop/len10/lenOptimizedA-12          355382803                3.454 ns/op           0 B/op          0 allocs/op
BenchmarkStringCharLoop/len10/lenOptimizedB-12          353923364                3.329 ns/op           0 B/op          0 allocs/op
BenchmarkStringCharLoop/len10/stringsReader-12          100000000               10.23 ns/op            0 B/op          0 allocs/op
BenchmarkStringCharLoop/len1000000/rangeKey-12               604           1950987 ns/op               0 B/op          0 allocs/op
BenchmarkStringCharLoop/len1000000/rangeVal-12               718           1629614 ns/op               0 B/op          0 allocs/op
BenchmarkStringCharLoop/len1000000/len-12                   4858            228146 ns/op               0 B/op          0 allocs/op
BenchmarkStringCharLoop/len1000000/lenOptimizedA-12         4916            221950 ns/op               0 B/op          0 allocs/op
BenchmarkStringCharLoop/len1000000/lenOptimizedB-12         4852            221679 ns/op               0 B/op          0 allocs/op
BenchmarkStringCharLoop/len1000000/stringsReader-12         1172           1080518 ns/op               0 B/op          0 allocs/op
BenchmarkRandom64Bytes/random64ArrayExpanded-12           991857              1064 ns/op               0 B/op          0 allocs/op
BenchmarkRandom64Bytes/random64SliceBuffer-12            1062883              1148 ns/op              64 B/op          1 allocs/op
BenchmarkRandom64Bytes/randomSlice64Make-12              1000000              1012 ns/op              64 B/op          1 allocs/op
BenchmarkRandom64Bytes/random64SliceAppend-12             918925              1134 ns/op              64 B/op          1 allocs/op
BenchmarkRandom64Bytes/random64SliceReader-12           13395184                87.12 ns/op            0 B/op          0 allocs/op
BenchmarkRandom64Bytes/random64Array-12                  1000000              1007 ns/op               0 B/op          0 allocs/op
BenchmarkSliceAppendOne/simple-12                       23136823                48.25 ns/op           96 B/op          1 allocs/op
BenchmarkSliceAppendOne/appendEachWithHint-12           1000000000               0.2339 ns/op          0 B/op          0 allocs/op
BenchmarkSliceAppendOne/appendSliceWithHint-12          165629017                7.337 ns/op           0 B/op          0 allocs/op
BenchmarkSliceAppendOne/appendTogetherWithHint-12       1000000000               0.2338 ns/op          0 B/op          0 allocs/op
BenchmarkSliceAppendOne/preSized-12                     1000000000               0.2291 ns/op          0 B/op          0 allocs/op
BenchmarkSliceAppendOne/preSizedAll-12                  1000000000               0.2272 ns/op          0 B/op          0 allocs/op
BenchmarkStrHash-12                                     500094735                2.394 ns/op           0 B/op          0 allocs/op
BenchmarkHashFnv-12                                     225231529                5.537 ns/op           0 B/op          0 allocs/op
BenchmarkStringBufferReader/bufio.Reader-12              1197684              1002 ns/op            5768 B/op          8 allocs/op
BenchmarkStringBufferReader/native-12                   55936630                20.95 ns/op            0 B/op          0 allocs/op
BenchmarkStringStartWithShort/stringsHasPrefix-12       135734665                8.810 ns/op           0 B/op          0 allocs/op
BenchmarkStringStartWithShort/LoopStartsWith-12         144803272                8.259 ns/op           0 B/op          0 allocs/op
BenchmarkStringStartWithLong/stringsHasPrefix-12        319221398                3.948 ns/op           0 B/op          0 allocs/op
BenchmarkStringStartWithLong/LoopStartsWith-12          13972076                84.61 ns/op            0 B/op          0 allocs/op
BenchmarkTextConcatSmall/Sprint-12                      14502111                78.05 ns/op           16 B/op          1 allocs/op
BenchmarkTextConcatSmall/Plus-12                        1000000000               0.2356 ns/op          0 B/op          0 allocs/op
BenchmarkTextConcatSmall/Builder-12                     26493958                40.16 ns/op           24 B/op          2 allocs/op
BenchmarkTextConcatSmall/Join-12                        26392899                40.58 ns/op           16 B/op          1 allocs/op
BenchmarkTextConcatLarge/Plus-12                          157178              6758 ns/op           31728 B/op         99 allocs/op
BenchmarkTextConcatLarge/Builder-12                      2252326               529.4 ns/op          1912 B/op          8 allocs/op
BenchmarkTextConcatLarge/BuilderPrealloc-12              4879078               247.0 ns/op           640 B/op          1 allocs/op
BenchmarkTextConcatLarge/Join-12                          706540              1689 ns/op            5104 B/op          9 allocs/op
BenchmarkTextConcatLarge/JoinReserved-12                 1514073               793.8 ns/op           640 B/op          1 allocs/op
BenchmarkTextConcatLarge/JoinPrealloc-12                 1538802               765.0 ns/op           640 B/op          1 allocs/op
PASS
ok      github.com/stefanovazzocell/GoDojo/bench        136.308s
```