go-queue
========
[![Build Status](https://travis-ci.org/petrkotek/go-queue.svg?branch=master)](https://travis-ci.org/petrkotek/go-queue)
[![Coverage Status](https://coveralls.io/repos/petrkotek/go-queue/badge.svg?branch=master&service=github)](https://coveralls.io/github/petrkotek/go-vector?branch=master)

## About

Currently has only one implementation of a queue in golang:
- `OverflowingQueue`, which has specified capacity and starts discarding new items, when capacity is reached.

## Usage

### 1. Fetch the package 

```
go get github.com/petrkotek/go-queue
```

### 2. Import the package
Import the package into your `.go` file:

```go
package main

import (
        "fmt"

        "github.com/petrkotek/go-queue/queue"
)

func main() {    
        myQueue := queue.NewOverflowingQueue(10)
        myQueue.Enqueue(42)
        myQueue.Enqueue(123)
        fmt.Println(myQueue.Dequeue()) // 42
        fmt.Println(myQueue.Dequeue()) // 123
}
```

## License

Released under an MIT license. See `LICENSE.md` file for details.
