# Go Dojo

This is a [dojo](https://en.wikipedia.org/wiki/Dojo) to experiment with and learn from [Go](https://go.dev/).

# Snippets

Here I might add little code snippets as I experiment with things.

# Benchmarks

Run benchmarks with `make bench`. These benchmarks will run a series of tests to answer the question "What is the fastest way to do X in Go?"

## Results

Last update: `Nov 9th, 2023`
Go version: `1.21.4`

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
BenchmarkControlFlowSimple/switch-12            1000000000               0.5002 ns/op          0 B/op          0 allocs/op
BenchmarkControlFlowSimple/if-12                1000000000               0.5018 ns/op          0 B/op          0 allocs/op
BenchmarkControlFlowSimple/ifElse-12            1000000000               0.5135 ns/op          0 B/op          0 allocs/op
BenchmarkControlFlowComplex/switch-12           1000000000               0.9769 ns/op          0 B/op          0 allocs/op
BenchmarkControlFlowComplex/if-12               1000000000               0.7878 ns/op          0 B/op          0 allocs/op
BenchmarkControlFlowComplex/ifElse-12           1000000000               0.7925 ns/op          0 B/op          0 allocs/op
BenchmarkSliceLoopValue/rangeKey-12                 5332            222800 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopValue/rangeVal-12                 5494            219705 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopValue/len-12                      5527            218361 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopValue/lenOptimizedA-12            5511            211521 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopValue/lenOptimizedB-12            5487            218390 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopCount/rangeKey-12                  938           1189680 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopCount/rangeVal-12                  934           1202916 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopCount/rangeBlank-12                924           1194189 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopCount/len-12                       925           1197602 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopCount/lenOptimizedA-12             930           1178428 ns/op               0 B/op          0 allocs/op
BenchmarkSliceLoopCount/lenOptimizedB-12             945           1204513 ns/op               0 B/op          0 allocs/op
BenchmarkIntToString/Sprintf-12                 14194255                75.34 ns/op           16 B/op          1 allocs/op
BenchmarkIntToString/strconvItoa-12             42175081                23.79 ns/op            7 B/op          0 allocs/op
BenchmarkIntToString/strconvFormatInt-12        43847019                23.39 ns/op            7 B/op          0 allocs/op
BenchmarkLocks/Mutex/Sequential-12              99583570                11.78 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/Mutex/Parallel-12                13930578                77.25 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/RWMutex/Sequential-12            56117509                21.56 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/RWMutex/Parallel-12              13595974                76.03 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/RWMutexRead/Sequential-12        99338688                11.95 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/RWMutexRead/Parallel-12          27813744                40.17 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/Bool/Sequential-12               93801326                10.85 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/Bool/Parallel-12                  3882343               302.7 ns/op             0 B/op          0 allocs/op
BenchmarkLocks/int32/Sequential-12              85509216                11.72 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/int32/Parallel-12                 3887232               301.6 ns/op             0 B/op          0 allocs/op
BenchmarkLocks/int64/Sequential-12              104494976               11.50 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/int64/Parallel-12                 3819470               295.4 ns/op             0 B/op          0 allocs/op
BenchmarkLocks/Uint32/Sequential-12             91506321                11.22 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/Uint32/Parallel-12                3929732               299.5 ns/op             0 B/op          0 allocs/op
BenchmarkLocks/Uint64/Sequential-12             99550225                11.82 ns/op            0 B/op          0 allocs/op
BenchmarkLocks/Uint64/Parallel-12                3885295               299.1 ns/op             0 B/op          0 allocs/op
BenchmarkLocksNoWork/Mutex/Sequential-12        127493964                9.383 ns/op           0 B/op          0 allocs/op
BenchmarkLocksNoWork/Mutex/Parallel-12          13753642                81.47 ns/op            0 B/op          0 allocs/op
BenchmarkLocksNoWork/RWMutex/Sequential-12      61187467                19.30 ns/op            0 B/op          0 allocs/op
BenchmarkLocksNoWork/RWMutex/Parallel-12        16475547                65.23 ns/op            0 B/op          0 allocs/op
BenchmarkLocksNoWork/RWMutexRead/Sequential-12          127616060                9.281 ns/op           0 B/op          0 allocs/op
BenchmarkLocksNoWork/RWMutexRead/Parallel-12            26413293                43.41 ns/op            0 B/op          0 allocs/op
BenchmarkLocksNoWork/Bool/Sequential-12                 109843634               10.88 ns/op            0 B/op          0 allocs/op
BenchmarkLocksNoWork/Bool/Parallel-12                    4033866               279.3 ns/op             0 B/op          0 allocs/op
BenchmarkLocksNoWork/int32/Sequential-12                138252537                8.370 ns/op           0 B/op          0 allocs/op
BenchmarkLocksNoWork/int32/Parallel-12                   3990505               279.4 ns/op             0 B/op          0 allocs/op
BenchmarkLocksNoWork/int64/Sequential-12                127890664                9.500 ns/op           0 B/op          0 allocs/op
BenchmarkLocksNoWork/int64/Parallel-12                   4179129               280.6 ns/op             0 B/op          0 allocs/op
BenchmarkLocksNoWork/Uint32/Sequential-12               141149962                8.487 ns/op           0 B/op          0 allocs/op
BenchmarkLocksNoWork/Uint32/Parallel-12                  4106037               279.1 ns/op             0 B/op          0 allocs/op
BenchmarkLocksNoWork/Uint64/Sequential-12               125594324                9.474 ns/op           0 B/op          0 allocs/op
BenchmarkLocksNoWork/Uint64/Parallel-12                  4170722               279.0 ns/op             0 B/op          0 allocs/op
BenchmarkStringCharLoop/len10/rangeKey-12               165256262                7.460 ns/op           0 B/op          0 allocs/op
BenchmarkStringCharLoop/len10/rangeVal-12               240964270                5.101 ns/op           0 B/op          0 allocs/op
BenchmarkStringCharLoop/len10/len-12                    213225636                5.749 ns/op           0 B/op          0 allocs/op
BenchmarkStringCharLoop/len10/lenOptimizedA-12          212220340                5.578 ns/op           0 B/op          0 allocs/op
BenchmarkStringCharLoop/len10/lenOptimizedB-12          202851536                5.596 ns/op           0 B/op          0 allocs/op
BenchmarkStringCharLoop/len10/stringsReader-12          98036756                10.34 ns/op            0 B/op          0 allocs/op
BenchmarkStringCharLoop/len1000000/rangeKey-12               627           1928626 ns/op               0 B/op          0 allocs/op
BenchmarkStringCharLoop/len1000000/rangeVal-12               668           1501312 ns/op               0 B/op          0 allocs/op
BenchmarkStringCharLoop/len1000000/len-12                   5407            225982 ns/op               0 B/op          0 allocs/op
BenchmarkStringCharLoop/len1000000/lenOptimizedA-12         5408            230868 ns/op               0 B/op          0 allocs/op
BenchmarkStringCharLoop/len1000000/lenOptimizedB-12         5342            216945 ns/op               0 B/op          0 allocs/op
BenchmarkStringCharLoop/len1000000/stringsReader-12         1086           1091804 ns/op               0 B/op          0 allocs/op
BenchmarkRandom64Bytes/random64SliceBuffer-12            1041666              1145 ns/op              64 B/op          1 allocs/op
BenchmarkRandom64Bytes/randomSlice64Make-12              1195650               994.3 ns/op            64 B/op          1 allocs/op
BenchmarkRandom64Bytes/random64SliceAppend-12            1060740              1158 ns/op              64 B/op          1 allocs/op
BenchmarkRandom64Bytes/random64SliceReader-12           14907180                79.36 ns/op            0 B/op          0 allocs/op
BenchmarkRandom64Bytes/random64Array-12                  1195981               973.0 ns/op             0 B/op          0 allocs/op
BenchmarkRandom64Bytes/random64ArrayExpanded-12           902024              1115 ns/op               0 B/op          0 allocs/op
BenchmarkStrHash-12                                     371472663                2.858 ns/op           0 B/op          0 allocs/op
BenchmarkHashFnv-12                                     118002183                9.185 ns/op           0 B/op          0 allocs/op
BenchmarkStringBufferReader/bufio.Reader-12               741102              1391 ns/op            5768 B/op          8 allocs/op
BenchmarkStringBufferReader/native-12                   60258462                21.53 ns/op            0 B/op          0 allocs/op
BenchmarkStringStartWithShort/stringsHasPrefix-12       124087538                8.998 ns/op           0 B/op          0 allocs/op
BenchmarkStringStartWithShort/LoopStartsWith-12         137318502                8.408 ns/op           0 B/op          0 allocs/op
BenchmarkStringStartWithLong/stringsHasPrefix-12        212649112                5.673 ns/op           0 B/op          0 allocs/op
BenchmarkStringStartWithLong/LoopStartsWith-12          13551284                78.82 ns/op            0 B/op          0 allocs/op
BenchmarkTextConcatSmall/Sprint-12                      13660218                78.43 ns/op           16 B/op          1 allocs/op
BenchmarkTextConcatSmall/Plus-12                        1000000000               0.2327 ns/op          0 B/op          0 allocs/op
BenchmarkTextConcatSmall/Builder-12                     28635480                41.96 ns/op           24 B/op          2 allocs/op
BenchmarkTextConcatSmall/Join-12                        30907629                38.64 ns/op           16 B/op          1 allocs/op
BenchmarkTextConcatLarge/Plus-12                          123211              9137 ns/op           31728 B/op         99 allocs/op
BenchmarkTextConcatLarge/Builder-12                      1859290               649.4 ns/op          1912 B/op          8 allocs/op
BenchmarkTextConcatLarge/BuilderPrealloc-12              4104066               274.5 ns/op           640 B/op          1 allocs/op
BenchmarkTextConcatLarge/Join-12                          599594              1907 ns/op            4720 B/op          9 allocs/op
BenchmarkTextConcatLarge/JoinReserved-12                 1481776               798.2 ns/op           640 B/op          1 allocs/op
BenchmarkTextConcatLarge/JoinPrealloc-12                 1533427               775.2 ns/op           640 B/op          1 allocs/op
PASS
ok      github.com/stefanovazzocell/GoDojo/bench        139.798s
```