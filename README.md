# Go-Log

[![Go Report Card](https://goreportcard.com/badge/github.com/yashishdua/Go-Log)](https://goreportcard.com/report/github.com/yashishdua/Go-Log)

Go-Log is a utility log package built to customize the Go's normal log package. Prominent features are
* Tag the logs into debug and error variant.
* Add/Remove Timestamp to logs
* Get the calling function details in logs

### Install

```bash
go get github.com/yashishdua/go-log
```

### Import

```go
import (
	"github.com/yashishdua/go-log"
)
```
## Examples

### 1. GoLog Debug 

```go
package main

import "go-log"

func main() {
	golog.D("This a basic debug log without time")
}
```

This golog prints the message without any extra information like:

```go
$ This a basic primitive debug log
```
