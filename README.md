# slicenilcmp

`slicenilcmp` detects comparisons between slice and nil.

```go
func f() {
	var basicSlice []int
	if basicSlice == nil { // want "suggestion: use len func for empty check"
		print(basicSlice)
	}

	if len(basicSlice) == 0 { // ok
		print(basicSlice)
	}
}
```

## Install

```bash
$ go install github.com/replu/slicenilcmp/cmd/slicenilcmp@latest
```

## How to use

```bash
$ go vet -vettool=$(which slicenilcmp) ./...
```
