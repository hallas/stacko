# stacko [![GoDoc](https://godoc.org/github.com/hallas/stacko?status.svg)](https://godoc.org/github.com/hallas/stacko)

Stacko lets you generate a structured stacktrace in Go.

```go
type Frame struct {
  FileName     string
  FunctionName string
  PackageName  string
  Path         string
  LineNumber   int
  InDomain     bool
  PreContext   []string
  PostContext  []string
  Context      string
}
```

Most of the fields are self explanatory. `InDomain` is a boolean that tells you
if the frame is within the same package as the first call. The `Context` fields
are the actual lines of code that precede and proceed the context of the frame.
Please note that these context fields are only provided if the source code is
present in the local file system.

## API

### type Stacktrace

```go
type Stacktrace []Frame
```

### func NewStacktrace

```go
func NewStacktrace (skip int) (Stacktrace, error)
```

Returns a new `Stacktrace` skipping the first `skip` frames.
