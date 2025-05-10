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
ok  	github.com/stefanovazzocell/GoDojo/bench	1.011s
Running benchmarks...
go test -run=Bench -bench=. -benchmem ./bench
goos: linux
goarch: amd64
pkg: github.com/stefanovazzocell/GoDojo/bench
cpu: AMD Ryzen AI 9 HX 370 w/ Radeon 890M
BenchmarkControlFlowSimple/switch-24         	1000000000	         0.3535 ns/op	       0 B/op	       0 allocs/op
BenchmarkControlFlowSimple/if-24             	1000000000	         0.3302 ns/op	       0 B/op	       0 allocs/op
BenchmarkControlFlowSimple/ifElse-24         	1000000000	         0.3227 ns/op	       0 B/op	       0 allocs/op
BenchmarkControlFlowComplex/switch-24        	1000000000	         0.8317 ns/op	       0 B/op	       0 allocs/op
BenchmarkControlFlowComplex/if-24            	1000000000	         0.4916 ns/op	       0 B/op	       0 allocs/op
BenchmarkControlFlowComplex/ifElse-24        	1000000000	         0.6160 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/rangeKey-24          	    5245	    207123 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/rangeKeyReuse-24     	    5863	    205679 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/rangeVal-24          	    5380	    205945 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/rangeValReuse-24     	    5445	    203485 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/len-24               	    5385	    205793 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/lenOptimizedA-24     	    5085	    206980 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/lenOptimizedB-24     	    4803	    210345 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/lenRange-24          	    5820	    210423 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/lenRangeReuse-24     	    5018	    210909 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/rangeKey-24          	    5722	    207709 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/rangeKeyReuse-24     	    5214	    206103 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/rangeVal-24          	    5650	    206912 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/rangeValReuse-24     	    5779	    206523 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/rangeBlank-24        	    5827	    207629 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/len-24               	    5629	    207444 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/lenOptimizedA-24     	    5070	    211700 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/lenOptimizedB-24     	    5744	    205878 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/lenRange-24          	    5731	    214086 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntToString/Sprintf-24              	24965995	        41.78 ns/op	      16 B/op	       2 allocs/op
BenchmarkIntToString/strconvItoa-24          	72906267	        14.45 ns/op	       7 B/op	       0 allocs/op
BenchmarkIntToString/strconvFormatInt-24     	75109687	        14.19 ns/op	       7 B/op	       0 allocs/op
BenchmarkLocks/Mutex/Sequential-24           	148940631	         8.020 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Mutex/Parallel-24             	18127080	        71.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/RWMutex/Sequential-24         	63156550	        17.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/RWMutex/Parallel-24           	13779932	        93.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/RWMutexRead/Sequential-24     	139192202	         8.107 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/RWMutexRead/Parallel-24       	41072318	        31.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Bool/Sequential-24            	154083508	         7.777 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Bool/Parallel-24              	 2786180	       444.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/int32/Sequential-24           	151202598	         7.856 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/int32/Parallel-24             	 2386492	       500.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/int64/Sequential-24           	150389154	         8.028 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/int64/Parallel-24             	 2852761	       456.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Uint32/Sequential-24          	157151198	         7.671 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Uint32/Parallel-24            	 2845588	       456.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Uint64/Sequential-24          	152437029	         7.872 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Uint64/Parallel-24            	 2825475	       415.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Mutex/Sequential-24     	140504902	         8.124 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Mutex/Parallel-24       	19374847	        83.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/RWMutex/Sequential-24   	65043633	        16.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/RWMutex/Parallel-24     	13123956	        84.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/RWMutexRead/Sequential-24         	100000000	        11.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/RWMutexRead/Parallel-24           	33867933	        34.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Bool/Sequential-24                	150470889	         7.909 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Bool/Parallel-24                  	 3185300	       393.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/int32/Sequential-24               	149638389	         7.992 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/int32/Parallel-24                 	 3151860	       381.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/int64/Sequential-24               	142486820	         7.930 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/int64/Parallel-24                 	 3165752	       382.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Uint32/Sequential-24              	152290150	         9.154 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Uint32/Parallel-24                	 3119479	       383.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Uint64/Sequential-24              	151967490	         7.865 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Uint64/Parallel-24                	 3187993	       376.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/rangeKey-24              	304931546	         3.679 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/rangeVal-24              	519424158	         2.393 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/len-24                   	494705661	         2.330 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/lenOptimizedA-24         	509664469	         2.315 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/lenOptimizedB-24         	503446186	         2.454 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/stringsReader-24         	246770932	         4.915 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/rangeKey-24         	    1072	   1107840 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/rangeVal-24         	    1128	   1053047 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/len-24              	    5288	    204290 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/lenOptimizedA-24    	    4966	    211040 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/lenOptimizedB-24    	    5299	    205153 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/stringsReader-24    	    2758	    416177 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandom64Bytes/randomSlice64Make-24            	 1799664	       697.9 ns/op	      64 B/op	       1 allocs/op
BenchmarkRandom64Bytes/random64SliceAppend-24          	 1637896	       765.7 ns/op	      64 B/op	       1 allocs/op
BenchmarkRandom64Bytes/random64SliceReader-24          	28044866	        41.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandom64Bytes/random64Array-24                	 1857546	       646.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandom64Bytes/random64ArrayExpanded-24        	 1878199	       637.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandom64Bytes/random64SliceBuffer-24          	 1659165	       724.6 ns/op	      64 B/op	       1 allocs/op
BenchmarkSha256three/small/manyWrites-24               	16751763	        70.93 ns/op	      32 B/op	       1 allocs/op
BenchmarkSha256three/small/append-24                   	12299665	        97.69 ns/op	     160 B/op	       3 allocs/op
BenchmarkSha256three/small/buffer-24                   	13688967	        87.31 ns/op	      96 B/op	       2 allocs/op
BenchmarkSha256three/small/preprocess-24               	18867403	        65.71 ns/op	      32 B/op	       1 allocs/op
BenchmarkSha256three/large/manyWrites-24               	 1474746	       841.1 ns/op	      32 B/op	       1 allocs/op
BenchmarkSha256three/large/append-24                   	 1146661	      1115 ns/op	    3360 B/op	       3 allocs/op
BenchmarkSha256three/large/buffer-24                   	 1209270	       990.7 ns/op	    1824 B/op	       2 allocs/op
BenchmarkSha256three/large/preprocess-24               	 1482775	       795.9 ns/op	      32 B/op	       1 allocs/op
BenchmarkSliceAppendOne/simple-24                      	45893528	        28.47 ns/op	      96 B/op	       1 allocs/op
BenchmarkSliceAppendOne/appendEachWithHint-24          	1000000000	         0.2049 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAppendOne/appendSliceWithHint-24         	223201453	         5.112 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAppendOne/appendTogetherWithHint-24      	1000000000	         0.2051 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAppendOne/preSized-24                    	1000000000	         0.2046 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAppendOne/preSizedAll-24                 	1000000000	         0.2060 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrHash-24                                    	983582386	         1.208 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashFnv-24                                    	338362573	         3.530 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringBufferReader/bufio.Reader-24            	 1643770	       687.6 ns/op	    5768 B/op	       8 allocs/op
BenchmarkStringBufferReader/native-24                  	97091431	        10.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringStartWithShort/stringsHasPrefix-24      	283146157	         4.255 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringStartWithShort/LoopStartsWith-24        	302953368	         4.022 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringStartWithLong/stringsHasPrefix-24       	542093390	         2.138 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringStartWithLong/LoopStartsWith-24         	16477773	        66.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwap/uint8/inline-24                          	1000000000	         0.2031 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwap/uint8/tmp-24                             	1000000000	         0.2032 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwap/uint8/xor-24                             	1000000000	         0.2032 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwap/int/inline-24                            	1000000000	         0.2033 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwap/int/tmp-24                               	1000000000	         0.2037 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwap/int/xor-24                               	1000000000	         0.2074 ns/op	       0 B/op	       0 allocs/op
BenchmarkTextConcatSmall/Sprint-24                     	23570220	        46.12 ns/op	      16 B/op	       1 allocs/op
BenchmarkTextConcatSmall/Plus-24                       	1000000000	         0.2054 ns/op	       0 B/op	       0 allocs/op
BenchmarkTextConcatSmall/Builder-24                    	46300608	        23.89 ns/op	      24 B/op	       2 allocs/op
BenchmarkTextConcatSmall/Join-24                       	48222054	        22.88 ns/op	      16 B/op	       1 allocs/op
BenchmarkTextConcatLarge/Plus-24                       	  226124	      4667 ns/op	   31728 B/op	      99 allocs/op
BenchmarkTextConcatLarge/Builder-24                    	 3666709	       327.2 ns/op	    1912 B/op	       8 allocs/op
BenchmarkTextConcatLarge/BuilderPrealloc-24            	 8893496	       126.3 ns/op	     640 B/op	       1 allocs/op
BenchmarkTextConcatLarge/Join-24                       	 1012070	      1200 ns/op	    5104 B/op	       9 allocs/op
BenchmarkTextConcatLarge/JoinReserved-24               	 2334705	       499.8 ns/op	     640 B/op	       1 allocs/op
BenchmarkTextConcatLarge/JoinPrealloc-24               	 2651406	       473.3 ns/op	     640 B/op	       1 allocs/op
BenchmarkValidateZeroBits/binaryBits-24                	1000000000	         1.152 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateZeroBits/mask-24                      	995368393	         1.209 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateZeroBits/loop-24                      	652298214	         1.830 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/stefanovazzocell/GoDojo/bench	159.999s
```
