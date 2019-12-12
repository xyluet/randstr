# Randstr

A simple random string generator using `crypto/rand` package.

## Install

```bash
go get -u github.com/xyluet/randstr
```

## Example

```go
s, err := randstr.String(32)
if err != nil {
  // check err
}
fmt.Println(s) // example output: cpjszbw8jrSFXp42
```

Example above will generate a random string default charset `[A-z0-9]`.

```go
s, err := randstr.StringWithCharset(16, "ABC")
if err != nil { panic(err) }
fmt.Println(s) // example output: BBAAACCCBBCBBCAA
```

## References

- https://github.com/golang/go/wiki/CodeReviewComments#crypto-rand
- https://www.calhoun.io/creating-random-strings-in-go/
