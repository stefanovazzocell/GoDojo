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
