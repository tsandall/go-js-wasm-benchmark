# go-js-wasm-benchmark

Running `make` will build the test program and execute the benchmark.

## Example output

```
JS Latency (ms)
 min: 0.19006499648094177
 max: 3.666119009256363
 mean: 1.6668863606452942
 median: 1.7124930024147034
 90th: 1.7124930024147034
 95th: 1.914260149002074
 99th: 3.2065688967704795
Go Latency (ms)
 min: 0.004864
 max: 0.020992
 mean: 0.008998400000000004
 median: 0.007936
 90th: 0.007936
 95th: 0.017152
 99th: 0.020738560000000003
```

> Ran on 2016 MacBook Pro (Intel Core i7 2.9 GHz) w/ Go 1.11 and node 10.9.0.

## Notes

* The `syscall/js#Callback` interface does not support return values. Some posts
  have recommended setting a value in the globals to communicate return values,
  however, calls are not synchronous so that method is not reliable. This
  benchmark passes a callback into the Go code so that JS callers can wrap the
  code in a Promise and provide a pseudo-synchronous API.
