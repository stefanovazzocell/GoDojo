# Go Dojo

This is a [dojo](https://en.wikipedia.org/wiki/Dojo) to experiment with and learn from [Go](https://go.dev/).

# Snippets

Here I might add little code snippets as I experiment with things.

# Benchmarks

Run benchmarks with `make bench`. These benchmarks will run a series of tests to answer the question "What is the fastest way to do X in Go?"

## Results

Last update: `May 10, 2025`
Go version: `1.24.2`

```
Running tests...
go test -run=Test. -race ./bench
ok  	github.com/stefanovazzocell/GoDojo/bench	1.010s
Running benchmarks...
go test -run=Bench -bench=. -benchmem ./bench
goos: linux
goarch: amd64
pkg: github.com/stefanovazzocell/GoDojo/bench
cpu: AMD Ryzen AI 9 HX 370 w/ Radeon 890M
BenchmarkControlFlowSimple/switch-24         	1000000000	         0.3274 ns/op	       0 B/op	       0 allocs/op
BenchmarkControlFlowSimple/if-24             	1000000000	         0.3388 ns/op	       0 B/op	       0 allocs/op
BenchmarkControlFlowSimple/ifElse-24         	1000000000	         0.3300 ns/op	       0 B/op	       0 allocs/op
BenchmarkControlFlowComplex/switch-24        	1000000000	         0.8283 ns/op	       0 B/op	       0 allocs/op
BenchmarkControlFlowComplex/if-24            	1000000000	         0.4789 ns/op	       0 B/op	       0 allocs/op
BenchmarkControlFlowComplex/ifElse-24        	1000000000	         0.6126 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/rangeKey-24          	    5401	    205943 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/rangeKeyReuse-24     	    5532	    202996 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/rangeVal-24          	    5589	    209579 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/rangeValReuse-24     	    5749	    201576 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/len-24               	    5301	    205321 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/lenOptimizedA-24     	    5689	    204867 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/lenOptimizedB-24     	    5377	    204227 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/lenRange-24          	    5701	    202204 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/lenRangeReuse-24     	    5695	    203562 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/rangeKey-24          	    5739	    209611 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/rangeKeyReuse-24     	    5666	    206885 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/rangeVal-24          	    5821	    207888 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/rangeValReuse-24     	    5826	    206208 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/rangeBlank-24        	    5787	    205962 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/len-24               	    5362	    205280 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/lenOptimizedA-24     	    5757	    205270 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/lenOptimizedB-24     	    5791	    206674 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/lenRange-24          	    5581	    210986 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntToString/Sprintf-24              	29432930	        42.67 ns/op	      16 B/op	       2 allocs/op
BenchmarkIntToString/strconvItoa-24          	88762628	        14.48 ns/op	       7 B/op	       0 allocs/op
BenchmarkIntToString/strconvFormatInt-24     	81055701	        14.16 ns/op	       7 B/op	       0 allocs/op
BenchmarkLocks/Mutex/Sequential-24           	147383233	         8.125 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Mutex/Parallel-24             	20973745	        69.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/RWMutex/Sequential-24         	67234214	        17.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/RWMutex/Parallel-24           	13635768	        79.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/RWMutexRead/Sequential-24     	148237738	         8.092 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/RWMutexRead/Parallel-24       	41474112	        31.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Bool/Sequential-24            	156535290	         7.585 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Bool/Parallel-24              	 2773450	       442.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/int32/Sequential-24           	152978418	         7.902 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/int32/Parallel-24             	 2774194	       455.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/int64/Sequential-24           	151763522	         7.864 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/int64/Parallel-24             	 2739468	       445.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Uint32/Sequential-24          	152674465	         7.708 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Uint32/Parallel-24            	 2713032	       451.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Uint64/Sequential-24          	153041026	         7.821 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Uint64/Parallel-24            	 2766920	       439.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Mutex/Sequential-24     	147699912	         8.042 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Mutex/Parallel-24       	16815901	        82.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/RWMutex/Sequential-24   	49417147	        22.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/RWMutex/Parallel-24     	15256942	        97.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/RWMutexRead/Sequential-24         	100000000	        11.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/RWMutexRead/Parallel-24           	35140512	        34.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Bool/Sequential-24                	152010668	         7.879 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Bool/Parallel-24                  	 3261990	       385.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/int32/Sequential-24               	150606291	         7.958 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/int32/Parallel-24                 	 3185926	       384.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/int64/Sequential-24               	151553958	         7.908 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/int64/Parallel-24                 	 3163852	       388.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Uint32/Sequential-24              	151399731	         7.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Uint32/Parallel-24                	 3340180	       371.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Uint64/Sequential-24              	151101516	         7.939 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Uint64/Parallel-24                	 3286348	       379.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/rangeKey-24              	292074915	         4.202 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/rangeVal-24              	522169681	         2.269 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/len-24                   	518135922	         2.304 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/lenOptimizedA-24         	514271253	         2.311 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/lenOptimizedB-24         	516660492	         2.296 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/stringsReader-24         	246315740	         4.889 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/rangeKey-24         	    1057	   1077093 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/rangeVal-24         	    1146	   1056490 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/len-24              	    5953	    206014 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/lenOptimizedA-24    	    5931	    201337 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/lenOptimizedB-24    	    5911	    203616 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/stringsReader-24    	    2780	    417619 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandom64Bytes/random64SliceBuffer-24          	 1573840	       712.0 ns/op	      64 B/op	       1 allocs/op
BenchmarkRandom64Bytes/randomSlice64Make-24            	 1773127	       681.4 ns/op	      64 B/op	       1 allocs/op
BenchmarkRandom64Bytes/random64SliceAppend-24          	 1619406	       719.8 ns/op	      64 B/op	       1 allocs/op
BenchmarkRandom64Bytes/random64SliceReader-24          	28839288	        41.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandom64Bytes/random64Array-24                	 1845291	       648.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandom64Bytes/random64ArrayExpanded-24        	 1884180	       638.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkSha256three/small/manyWrites-24               	14916517	        69.36 ns/op	      32 B/op	       1 allocs/op
BenchmarkSha256three/small/append-24                   	10779036	        97.63 ns/op	     160 B/op	       3 allocs/op
BenchmarkSha256three/small/buffer-24                   	13598667	        87.93 ns/op	      96 B/op	       2 allocs/op
BenchmarkSha256three/small/preprocess-24               	18045414	        65.75 ns/op	      32 B/op	       1 allocs/op
BenchmarkSha256three/large/manyWrites-24               	 1498404	       800.0 ns/op	      32 B/op	       1 allocs/op
BenchmarkSha256three/large/append-24                   	 1079098	      1053 ns/op	    3360 B/op	       3 allocs/op
BenchmarkSha256three/large/buffer-24                   	 1000000	      1038 ns/op	    1824 B/op	       2 allocs/op
BenchmarkSha256three/large/preprocess-24               	 1465922	       814.3 ns/op	      32 B/op	       1 allocs/op
BenchmarkSliceAppendOne/simple-24                      	49459909	        24.64 ns/op	      96 B/op	       1 allocs/op
BenchmarkSliceAppendOne/appendEachWithHint-24          	1000000000	         0.2056 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAppendOne/appendSliceWithHint-24         	240882889	         5.066 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAppendOne/appendTogetherWithHint-24      	1000000000	         0.2040 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAppendOne/preSized-24                    	1000000000	         0.2040 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAppendOne/preSizedAll-24                 	1000000000	         0.2041 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrHash-24                                    	994711549	         1.203 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashFnv-24                                    	317444572	         3.697 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringBufferReader/bufio.Reader-24            	 1728368	      1161 ns/op	    5768 B/op	       8 allocs/op
BenchmarkStringBufferReader/native-24                  	100000000	        10.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringStartWithShort/stringsHasPrefix-24      	279355068	         4.274 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringStartWithShort/LoopStartsWith-24        	307464217	         3.871 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringStartWithLong/stringsHasPrefix-24       	563386483	         2.124 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringStartWithLong/LoopStartsWith-24         	19057399	        60.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwap/uint8/inline-24                          	1000000000	         0.2055 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwap/uint8/tmp-24                             	1000000000	         0.2094 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwap/uint8/xor-24                             	1000000000	         0.2056 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwap/int/inline-24                            	1000000000	         0.2028 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwap/int/tmp-24                               	1000000000	         0.2039 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwap/int/xor-24                               	1000000000	         0.2242 ns/op	       0 B/op	       0 allocs/op
BenchmarkTextConcatSmall/Sprint-24                     	26166070	        46.40 ns/op	      16 B/op	       1 allocs/op
BenchmarkTextConcatSmall/Plus-24                       	1000000000	         0.2045 ns/op	       0 B/op	       0 allocs/op
BenchmarkTextConcatSmall/Builder-24                    	49797879	        24.05 ns/op	      24 B/op	       2 allocs/op
BenchmarkTextConcatSmall/Join-24                       	44334462	        22.80 ns/op	      16 B/op	       1 allocs/op
BenchmarkTextConcatLarge/Plus-24                       	  127081	      9065 ns/op	   31728 B/op	      99 allocs/op
BenchmarkTextConcatLarge/Builder-24                    	 2592487	       419.7 ns/op	    1912 B/op	       8 allocs/op
BenchmarkTextConcatLarge/BuilderPrealloc-24            	 9488110	       129.6 ns/op	     640 B/op	       1 allocs/op
BenchmarkTextConcatLarge/Join-24                       	 1226581	       974.4 ns/op	    5104 B/op	       9 allocs/op
BenchmarkTextConcatLarge/JoinReserved-24               	 2362963	       506.3 ns/op	     640 B/op	       1 allocs/op
BenchmarkTextConcatLarge/JoinPrealloc-24               	 1704064	       618.0 ns/op	     640 B/op	       1 allocs/op
PASS
ok  	github.com/stefanovazzocell/GoDojo/bench	162.402s
```
