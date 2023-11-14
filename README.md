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
ok  	github.com/stefanovazzocell/GoDojo/bench	1.008s
Running benchmarks...
go test -run=Bench -bench=. -benchmem ./bench
goos: linux
goarch: amd64
pkg: github.com/stefanovazzocell/GoDojo/bench
cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
BenchmarkControlFlowSimple/switch-12         	1000000000	         0.4971 ns/op	       0 B/op	       0 allocs/op
BenchmarkControlFlowSimple/if-12             	1000000000	         0.5059 ns/op	       0 B/op	       0 allocs/op
BenchmarkControlFlowSimple/ifElse-12         	1000000000	         0.5087 ns/op	       0 B/op	       0 allocs/op
BenchmarkControlFlowComplex/switch-12        	1000000000	         0.9779 ns/op	       0 B/op	       0 allocs/op
BenchmarkControlFlowComplex/if-12            	1000000000	         0.7684 ns/op	       0 B/op	       0 allocs/op
BenchmarkControlFlowComplex/ifElse-12        	1000000000	         0.7819 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/rangeKey-12          	    2679	    433583 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/rangeVal-12          	    2736	    419979 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/len-12               	    2697	    439555 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/lenOptimizedA-12     	    2641	    439027 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopValue/lenOptimizedB-12     	    2632	    433556 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/rangeKey-12          	    1012	   1159858 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/rangeVal-12          	    1017	   1130863 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/rangeBlank-12        	    1017	   1173158 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/len-12               	    1027	   1191256 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/lenOptimizedA-12     	    1010	   1152723 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLoopCount/lenOptimizedB-12     	    1011	   1145234 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntToString/Sprintf-12              	16380276	        72.35 ns/op	      16 B/op	       1 allocs/op
BenchmarkIntToString/strconvItoa-12          	53153370	        23.26 ns/op	       7 B/op	       0 allocs/op
BenchmarkIntToString/strconvFormatInt-12     	52245962	        23.38 ns/op	       7 B/op	       0 allocs/op
BenchmarkLocks/Mutex/Sequential-12           	99097232	        11.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Mutex/Parallel-12             	15280332	        70.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/RWMutex/Sequential-12         	51812480	        20.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/RWMutex/Parallel-12           	17000750	        64.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/RWMutexRead/Sequential-12     	100018202	        12.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/RWMutexRead/Parallel-12       	27283442	        42.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Bool/Sequential-12            	100000000	        10.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Bool/Parallel-12              	 4004400	       301.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/int32/Sequential-12           	100000000	        11.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/int32/Parallel-12             	 4263741	       277.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/int64/Sequential-12           	97288614	        11.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/int64/Parallel-12             	 4169286	       283.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Uint32/Sequential-12          	100000000	        11.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Uint32/Parallel-12            	 4255006	       290.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Uint64/Sequential-12          	99919347	        11.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocks/Uint64/Parallel-12            	 4284783	       278.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Mutex/Sequential-12     	126421218	         9.333 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Mutex/Parallel-12       	14762418	        75.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/RWMutex/Sequential-12   	52618564	        19.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/RWMutex/Parallel-12     	17022159	        64.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/RWMutexRead/Sequential-12         	126744830	         9.475 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/RWMutexRead/Parallel-12           	26613703	        42.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Bool/Sequential-12                	100000000	        10.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Bool/Parallel-12                  	 4691275	       252.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/int32/Sequential-12               	139910595	         8.462 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/int32/Parallel-12                 	 4082986	       280.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/int64/Sequential-12               	129382972	         9.244 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/int64/Parallel-12                 	 4405255	       267.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Uint32/Sequential-12              	143594055	         8.466 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Uint32/Parallel-12                	 4424598	       266.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Uint64/Sequential-12              	127184974	         9.443 ns/op	       0 B/op	       0 allocs/op
BenchmarkLocksNoWork/Uint64/Parallel-12                	 4443876	       267.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/rangeKey-12              	166610340	         7.721 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/rangeVal-12              	239890425	         4.863 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/len-12                   	216733774	         5.681 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/lenOptimizedA-12         	207720427	         5.519 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/lenOptimizedB-12         	216538434	         5.524 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len10/stringsReader-12         	125230806	         9.600 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/rangeKey-12         	     656	   1923256 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/rangeVal-12         	     716	   1514781 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/len-12              	    5262	    220144 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/lenOptimizedA-12    	    5218	    220848 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/lenOptimizedB-12    	    5227	    220434 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringCharLoop/len1000000/stringsReader-12    	    1202	   1048560 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandom64Bytes/random64ArrayExpanded-12        	 1000000	      1094 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandom64Bytes/random64SliceBuffer-12          	  991952	      1125 ns/op	      64 B/op	       1 allocs/op
BenchmarkRandom64Bytes/randomSlice64Make-12            	 1200784	       995.0 ns/op	      64 B/op	       1 allocs/op
BenchmarkRandom64Bytes/random64SliceAppend-12          	 1000000	      1129 ns/op	      64 B/op	       1 allocs/op
BenchmarkRandom64Bytes/random64SliceReader-12          	15218361	        75.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandom64Bytes/random64Array-12                	 1242116	      1029 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAppendOne/simple-12                      	18634490	        58.34 ns/op	      96 B/op	       1 allocs/op
BenchmarkSliceAppendOne/appendEachWithHint-12          	413496826	         2.980 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAppendOne/appendSliceWithHint-12         	159381884	         7.497 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAppendOne/appendTogetherWithHint-12      	1000000000	         0.2230 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAppendOne/preSized-12                    	1000000000	         0.2260 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAppendOne/preSizedAll-12                 	1000000000	         0.2256 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrHash-12                                    	474674284	         2.335 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashFnv-12                                    	128569242	         8.909 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringBufferReader/bufio.Reader-12            	  813409	      1469 ns/op	    5768 B/op	       8 allocs/op
BenchmarkStringBufferReader/native-12                  	53755509	        20.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringStartWithShort/stringsHasPrefix-12      	138812216	         8.650 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringStartWithShort/LoopStartsWith-12        	138421202	         8.522 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringStartWithLong/stringsHasPrefix-12       	189162297	         5.795 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringStartWithLong/LoopStartsWith-12         	14816086	        82.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkTextConcatSmall/Sprint-12                     	14657605	        76.89 ns/op	      16 B/op	       1 allocs/op
BenchmarkTextConcatSmall/Plus-12                       	1000000000	         0.2221 ns/op	       0 B/op	       0 allocs/op
BenchmarkTextConcatSmall/Builder-12                    	27704907	        40.50 ns/op	      24 B/op	       2 allocs/op
BenchmarkTextConcatSmall/Join-12                       	29348859	        38.18 ns/op	      16 B/op	       1 allocs/op
BenchmarkTextConcatLarge/Plus-12                       	  122786	      9679 ns/op	   31728 B/op	      99 allocs/op
BenchmarkTextConcatLarge/Builder-12                    	 1797486	       645.4 ns/op	    1912 B/op	       8 allocs/op
BenchmarkTextConcatLarge/BuilderPrealloc-12            	 4351908	       270.0 ns/op	     640 B/op	       1 allocs/op
BenchmarkTextConcatLarge/Join-12                       	  582578	      1955 ns/op	    4720 B/op	       9 allocs/op
BenchmarkTextConcatLarge/JoinReserved-12               	 1494246	       795.5 ns/op	     640 B/op	       1 allocs/op
BenchmarkTextConcatLarge/JoinPrealloc-12               	 1561329	       760.4 ns/op	     640 B/op	       1 allocs/op
PASS
ok  	github.com/stefanovazzocell/GoDojo/bench	128.696s
```