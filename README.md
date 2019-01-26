# TL;DR:

- Use xxhash by default.

- highwayhash is fast on large inputs, but slow on small inputs.  It claims to
  have a stronger collision resistance. If you care about such thing, it's a
  good choice.


# Benchmark results

- Go 1.11.2.
- Intel(R) Xeon(R) CPU E3-1505M v6 @ 3.00GHz

```
BenchmarkSeahashString/8-8     	100000000	        16.0 ns/op
BenchmarkSeahashString/8192-8  	  500000	      3683 ns/op
BenchmarkSeahashString/1048576-8         	    3000	    470030 ns/op
BenchmarkSeahashUInts1-8                 	100000000	        13.0 ns/op
BenchmarkFarmHashString/8-8              	100000000	        11.9 ns/op
BenchmarkFarmHashString/8192-8           	 2000000	       686 ns/op
BenchmarkFarmHashString/1048576-8        	   20000	     95575 ns/op
BenchmarkFarmHashUInts1-8                	100000000	        10.7 ns/op
BenchmarkXXHashString/8-8                	200000000	         7.70 ns/op
BenchmarkXXHashString/8192-8             	 3000000	       551 ns/op
BenchmarkXXHashString/1048576-8          	   20000	     77972 ns/op
BenchmarkXXHashUInts1-8                  	300000000	         4.76 ns/op
BenchmarkHighwayHashString/8-8           	30000000	        53.2 ns/op
BenchmarkHighwayHashString/8192-8        	 3000000	       581 ns/op
BenchmarkHighwayHashString/1048576-8     	   20000	     80398 ns/op
BenchmarkHighwayHashUInts1-8             	50000000	        23.4 ns/op
BenchmarkMurmur3HashString/8-8           	50000000	        25.6 ns/op
BenchmarkMurmur3HashString/8192-8        	 1000000	      1205 ns/op
BenchmarkMurmur3HashString/1048576-8     	   10000	    155219 ns/op
BenchmarkMurmur3HashUInts1-8             	100000000	        20.7 ns/op
```
