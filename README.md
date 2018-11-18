# Go-Log

[![Go Report Card](https://goreportcard.com/badge/github.com/MindorksOpenSource/Go-Log)](https://goreportcard.com/report/github.com/MindorksOpenSource/Go-Log)

Go-Log is a utility log package built to customize the Go's normal log package. Prominent features are
* Tag the logs into debug and error variant.
* Add/Remove Timestamp to logs
* Get the calling function details in logs
* Synchronized with mutex locks to handle multiple logging

### Install

```bash
go get github.com/MindorksOpenSource/Go-Log
```

### Import

```go
import (
	"github.com/MindorksOpenSource/Go-Log"
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

### 3. GoLog Debug/Error with Time

```go
func main() {
	golog.ConfigureTimer()
	golog.D("A debug log with time")
}
```

This golog prints the message with its timestamp like:

```bash
$ 2018/10/29 15:52:24  A debug log with time
```

### 4. GoLog Debug/Error with CallingFunction

```go
func main() {
	golog.ConfigureCallingFunction()
	golog.D("A debug log with calling function")
}
```

This golog prints the message with its timestamp like:

```bash
$ [main.main] A debug log with calling function
```

You can use both the configurations together also. Go gophers!


### Contributor
[Yashish Dua](https://github.com/yashishdua)


### Show some :heart:

[![GitHub stars](https://img.shields.io/github/stars/MindorksOpenSource/Go-Log.svg?style=social&label=Star)](https://github.com/MindorksOpenSource/Go-Log) [![GitHub forks](https://img.shields.io/github/forks/MindorksOpenSource/Go-Log.svg?style=social&label=Fork)](https://github.com/MindorksOpenSource/Go-Log/fork) [![GitHub watchers](https://img.shields.io/github/watchers/MindorksOpenSource/Go-Log.svg?style=social&label=Watch)](https://github.com/MindorksOpenSource/Go-Log) [![GitHub followers](https://img.shields.io/github/followers/yashishdua.svg?style=social&label=Follow)](https://github.com/MindorksOpenSource/Go-Log)  
[![Twitter Follow](https://img.shields.io/twitter/follow/duayashish.svg?style=social)](https://twitter.com/duayashish)
