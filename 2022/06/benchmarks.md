```
BenchmarkFind/5-10  	 9063030	       111.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkFind/6-10  	10161392	       118.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkFind/7-10  	 5984061	       200.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkFind/10-10 	 3842776	       310.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkFind/11-10 	 3254622	       371.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkFind/60-10 	 1950549	       603.7 ns/op	       0 B/op	       0 allocs/op
``

After getting rid of `clear` and instead just allocating a new map on every loop:

```
           │   old.out   │               new.out                │
           │   sec/op    │    sec/op     vs base                │
Find/5-10    111.7n ± 1%   117.5n ±  0%   +5.19% (p=0.000 n=10)
Find/6-10    117.3n ± 0%   105.3n ±  0%  -10.27% (p=0.000 n=10)
Find/7-10    197.0n ± 0%   184.8n ±  0%   -6.19% (p=0.000 n=10)
Find/10-10   304.5n ± 2%   268.0n ± 13%  -11.99% (p=0.000 n=10)
Find/11-10   372.4n ± 2%   334.8n ±  2%  -10.11% (p=0.002 n=10)
Find/60-10   615.4n ± 1%   370.0n ± 13%  -39.88% (p=0.000 n=10)
geomean      237.7n        205.8n        -13.42%
```
After using a bitwise shifting algorithm:

```
           │   old.out    │               new.out               │
           │    sec/op    │   sec/op     vs base                │
Find/5-10    112.75n ± 4%   10.30n ± 1%  -90.86% (p=0.000 n=10)
Find/6-10    100.15n ± 4%   13.39n ± 0%  -86.63% (p=0.000 n=10)
Find/7-10    177.80n ± 2%   18.61n ± 0%  -89.54% (p=0.000 n=10)
Find/10-10   246.30n ± 1%   29.49n ± 0%  -88.02% (p=0.000 n=10)
Find/11-10   327.45n ± 3%   32.94n ± 0%  -89.94% (p=0.000 n=10)
Find/60-10    341.7n ± 2%   175.2n ± 0%  -48.74% (p=0.000 n=10)
geomean       195.2n        27.54n       -85.89%
```
