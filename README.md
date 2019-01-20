# A simple hashmap

* Support for hashing and storing strings, integers, arrays, slices, and structs
* Prioritize speed over memory

## Use Cases

* Traversing or streaming a data set with repetitive elements, such as text parsing
* Validating the integrity of files or objects with checksums

## Benchmarks

```
BenchmarkPut-4           1000000              1072 ns/op
BenchmarkGet10-4        100000000               19.8 ns/op
BenchmarkGet100-4       100000000               21.4 ns/op
BenchmarkGet1000-4      50000000                33.7 ns/op
```