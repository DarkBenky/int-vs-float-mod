# Modulus Benchmark in Go

This project provides benchmarks to compare the performance of different methods for calculating modulus in Go. Specifically, it benchmarks integer modulus, floating-point modulus using `math.Mod`, and floating-point modulus using a truncation-based method.

## Files

- **`modulus_benchmark_test.go`**: Contains the benchmark functions.

## Benchmarks

### 1. Integer Modulus (`%` operator)

This benchmark tests the standard integer modulus operation using the `%` operator.

### 2. Floating-Point Modulus (`math.Mod`)

This benchmark tests the floating-point modulus operation using the `math.Mod` function.

### 3. Floating-Point Modulus (Truncation)

This benchmark tests a custom floating-point modulus operation implemented as:

```go
result = x - math.Trunc(x/y)*y
```

## Results

Below is an example of benchmark results:

```bash
goos: linux
goarch: amd64
pkg: main_test
cpu: Intel(R) Core(TM) i5-10300H CPU @ 2.50GHz
BenchmarkIntMod-8               1000000000               0.2475 ns/op
BenchmarkFloatMod-8              8556974               140.0 ns/op
BenchmarkFloatModTrunc-8        1000000000               0.9942 ns/op
PASS
ok      main_test       2.716s
```

### Interpretation

1. **`BenchmarkIntMod`**: Integer modulus is the fastest operation due to its simplicity and hardware-level implementation.
2. **`BenchmarkFloatMod`**: Floating-point modulus using `math.Mod` is significantly slower due to the overhead of the library function.
3. **`BenchmarkFloatModTrunc`**: Floating-point modulus using truncation is faster than `math.Mod` but slower than integer modulus.

## Running the Benchmarks

1. Save the benchmark code in a file named `modulus_benchmark_test.go`.
2. Run the benchmarks using:

   ```bash
   go test -bench=.
   ```

## Dependencies

- Go 1.19 or later.

## Notes

- These benchmarks are designed for testing modulus performance under repeated operations. Adjust the operations or constants to fit specific use cases.
- Floating-point operations are generally slower than integer operations due to the complexities of floating-point arithmetic.
