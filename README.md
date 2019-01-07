# TL;DR:

- Use xxhash by default.

- highwayhash is fast on large inputs, but slow on small inputs.  It claims to
  have a stronger collision resistance. If you care about such thing, it's a
  good choice.


# Benchmark results

- Go 1.11.2.
- Intel(R) Core(TM) i5-5300U CPU @ 2.30GHz, on Mint 19.

```
goos: linux
goarch: amd64
pkg: github.com/yasushi-saito/go-hash-shootout
BenchmarkSeahashString/8-4     	50000000	        25.7 ns/op
BenchmarkSeahashString/8192-4  	  200000	      5840 ns/op
BenchmarkSeahashString/1048576-4         	    2000	    782236 ns/op
BenchmarkSeahashUInts1-4                 	50000000	        20.5 ns/op
BenchmarkFarmHashString/8-4              	100000000	        17.0 ns/op
BenchmarkFarmHashString/8192-4           	 2000000	       918 ns/op
BenchmarkFarmHashString/1048576-4        	   10000	    135421 ns/op
BenchmarkFarmHashUInts1-4                	100000000	        16.3 ns/op
BenchmarkXXHashString/8-4                	100000000	        10.8 ns/op
BenchmarkXXHashString/8192-4             	 2000000	       755 ns/op
BenchmarkXXHashString/1048576-4          	   10000	    106732 ns/op
BenchmarkXXHashUInts1-4                  	200000000	         6.75 ns/op
BenchmarkHighwayHashString/8-4           	20000000	        77.9 ns/op
BenchmarkHighwayHashString/8192-4        	 2000000	       744 ns/op
BenchmarkHighwayHashString/1048576-4     	   10000	    103292 ns/op
BenchmarkHighwayHashUInts1-4             	50000000	        35.5 ns/op
```
