# lvlbp

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]

Leveled byte pool. Inspired by [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics).

## Rationale

Instead of having 1 pool for byte slices, we create multiple pools that are serving slices of a particular size. This reduces overhead for the smaller slices, because they will not be resized when we request for the bigger slices. Also library provides stats to see how many allocations do we have for each class.

## Install

```
go get github.com/cristalhq/lvlbp
```

## Example

```go
size := 42
bb := lvlbp.Get(size)
defer lvlbp.Put(bb)

// do something with bb of size `closestPowerOf2(42)`
```

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/lvlbp/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/lvlbp/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/lvlbp
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/lvlbp
[reportcard-img]: https://goreportcard.com/badge/cristalhq/lvlbp
[reportcard-url]: https://goreportcard.com/report/cristalhq/lvlbp
[coverage-img]: https://codecov.io/gh/cristalhq/lvlbp/branch/master/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/lvlbp
