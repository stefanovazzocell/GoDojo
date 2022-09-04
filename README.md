# Go Dojo

This is a [dojo](https://en.wikipedia.org/wiki/Dojo) to experiment with and learn from [Go](https://go.dev/).

# Snippets

Here I might add little code snippets as I experiment with things.

# Benchmarks

Run benchmarks with `make bench`. These benchmarks will run a series of tests to answer the question "What is the fastest way to do X in Go?"

## Recent Run

**Go 1.19**

| Benchmark Name | Iterations | ns/op | B/op | allocs/op |
|----------------|:-----------|:------|:-----|:----------|
| BenchmarkControlFlowSimple/switch | 1000000000 | 0.5457 | 0 | 0 |
| BenchmarkControlFlowSimple/if | 1000000000 | 0.5534 | 0 | 0 |
| BenchmarkControlFlowSimple/ifElse | 1000000000 | 0.5526 | 0 | 0 |
| BenchmarkControlFlowComplex/switch | 1000000000 | 1.001 | 0 | 0 |
| BenchmarkControlFlowComplex/if | 1000000000 | 0.5398 | 0 | 0 |
| BenchmarkControlFlowComplex/ifElse | 1000000000 | 0.5597 | 0 | 0 |
| BenchmarkSliceLoopValue/rangeKey | 2649 | 441371 | 0 | 0 |
| BenchmarkSliceLoopValue/rangeVal | 2686 | 426988 | 0 | 0 |
| BenchmarkSliceLoopValue/len | 2665 | 441655 | 0 | 0 |
| BenchmarkSliceLoopValue/lenOptimizedA | 2676 | 429415 | 0 | 0 |
| BenchmarkSliceLoopValue/lenOptimizedB | 2674 | 444470 | 0 | 0 |
| BenchmarkSliceLoopCount/rangeKey | 991 | 1157209 | 0 | 0 |
| BenchmarkSliceLoopCount/rangeVal | 1011 | 1198847 | 0 | 0 |
| BenchmarkSliceLoopCount/rangeBlank | 981 | 1172294 | 0 | 0 |
| BenchmarkSliceLoopCount/len | 994 | 1180548 | 0 | 0 |
| BenchmarkSliceLoopCount/lenOptimizedA | 1016 | 1178200 | 0 | 0 |
| BenchmarkSliceLoopCount/lenOptimizedB | 999 | 1189783 | 0 | 0 |
| BenchmarkIntToString/Sprintf | 17310380 | 67.59 | 16 | 1 |
| BenchmarkIntToString/strconvItoa | 54627283 | 22.08 | 7 | 0 |
| BenchmarkIntToString/strconvFormatInt | 54315912 | 22.10 | 7 | 0 |
| BenchmarkLocks/Mutex/Sequential-12 | 99728132 | 11.92 | 0 | 0 |
| BenchmarkLocks/Mutex/Parallel-12 | 20424799 | 60.74 | 0 | 0 |
| BenchmarkLocks/RWMutex/Sequential-12 | 54456723 | 21.22 | 0 | 0 |
| BenchmarkLocks/RWMutex/Parallel-12 | 19289060 | 63.53 | 0 | 0 |
| BenchmarkLocks/RWMutexRead/Sequential-12 | 91332409 | 12.11 | 0 | 0 |
| BenchmarkLocks/RWMutexRead/Parallel-12 | 41100186 | 28.95 | 0 | 0 |
| BenchmarkLocks/Bool/Sequential-12 | 74824760 | 15.51 | 0 | 0 |
| BenchmarkLocks/Bool/Parallel-12 | 28384950 | 42.43 | 0 | 0 |
| BenchmarkLocks/int32/Sequential-12 | 77558012 | 15.25 | 0 | 0 |
| BenchmarkLocks/int32/Parallel-12 | 28366207 | 42.29 | 0 | 0 |
| BenchmarkLocks/int64/Sequential-12 | 70657633 | 16.78 | 0 | 0 |
| BenchmarkLocks/int64/Parallel-12 | 28426708 | 41.95 | 0 | 0 |
| BenchmarkLocks/Uint32/Sequential-12 | 74940277 | 15.64 | 0 | 0 |
| BenchmarkLocks/Uint32/Parallel-12 | 28466409 | 42.08 | 0 | 0 |
| BenchmarkLocks/Uint64/Sequential-12 | 71589045 | 16.33 | 0 | 0 |
| BenchmarkLocks/Uint64/Parallel-12 | 28309184 | 42.17 | 0 | 0 |
| BenchmarkLocks/Mutex/Sequential-12 | 95810516 | 12.18 | 0 | 0 |
| BenchmarkLocks/Mutex/Parallel-12 | 19309189 | 60.73 | 0 | 0 |
| BenchmarkLocks/RWMutex/Sequential-12 | 54247112 | 21.36 | 0 | 0 |
| BenchmarkLocks/RWMutex/Parallel-12 | 19137802 | 62.44 | 0 | 0 |
| BenchmarkLocks/RWMutexRead/Sequential-12 | 91355722 | 11.91 | 0 | 0 |
| BenchmarkLocks/RWMutexRead/Parallel-12 | 33751198 | 35.47 | 0 | 0 |
| BenchmarkLocks/Bool/Sequential-12 | 100000000 | 11.40 | 0 | 0 |
| BenchmarkLocks/Bool/Parallel-12 | 3622476 | 328.3 | 0 | 0 |
| BenchmarkLocks/int32/Sequential-12 | 109838028 | 10.94 | 0 | 0 |
| BenchmarkLocks/int32/Parallel-12 | 3782245 | 353.1 | 0 | 0 |
| BenchmarkLocks/int64/Sequential-12 | 98682660 | 11.80 | 0 | 0 |
| BenchmarkLocks/int64/Parallel-12 | 3854472 | 308.7 | 0 | 0 |
| BenchmarkLocks/Uint32/Sequential-12 | 100000000 | 11.52 | 0 | 0 |
| BenchmarkLocks/Uint32/Parallel-12 | 3788566 | 331.7 | 0 | 0 |
| BenchmarkLocks/Uint64/Sequential-12 | 98974557 | 11.90 | 0 | 0 |
| BenchmarkLocks/Uint64/Parallel-12 | 3438306 | 347.1 | 0 | 0 |
| BenchmarkStringStartWithShort/stringsHasPrefix | 100000000 | 11.12 | 0 | 0 |
| BenchmarkStringStartWithShort/LoopStartsWith | 151383106 | 7.820 | 0 | 0 |
| BenchmarkStringStartWithLong/stringsHasPrefix | 215529836 | 5.787 | 0 | 0 |
| BenchmarkStringStartWithLong/LoopStartsWith | 15057891 | 78.96 | 0 | 0 |
| BenchmarkTextConcatSmall/Sprint | 15481489 | 74.62 | 16 | 1 |
| BenchmarkTextConcatSmall/Plus | 1000000000 | 0.2206 | 0 | 0 |
| BenchmarkTextConcatSmall/Builder | 28201354 | 40.07 | 24 | 2 |
| BenchmarkTextConcatSmall/Join | 32824320 | 35.28 | 16 | 1 |
| BenchmarkTextConcatLarge/Plus | 180787 | 6301 | 31728 | 99 |
| BenchmarkTextConcatLarge/Builder | 2339445 | 509.6 | 1912 | 8 |
| BenchmarkTextConcatLarge/BuilderPrealloc | 5886711 | 207.7 | 640 | 1 |
| BenchmarkTextConcatLarge/Join | 657538 | 1722 | 4720 | 9 |
| BenchmarkTextConcatLarge/JoinReserved | 1639556 | 731.4 | 640 | 1 |
| BenchmarkTextConcatLarge/JoinPrealloc | 1700940 | 697.6 | 640 | 1 |
