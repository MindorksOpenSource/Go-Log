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
func main() {
	golog.D("A basic primitive debug log.")
}
```

This golog prints the message without any extra information like:

```bash
$ A basic primitive debug log.
```

### 2. GoLog Error 

```go
func main() {
	golog.E("This a basic primitive error log.")
}
```

This golog prints the message while tagging it as an ERROR like:

```bash
$ [ERROR] This a basic primitive error log.
```
### Configuration (Additional Information)

```go
// Adds time to the log
golog.ConfigureTimer()
// Adds the calling function path to log
golog.ConfigureCallingFunction()
```
