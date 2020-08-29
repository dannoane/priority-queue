# priority-queue
> priority queue implementation in golang

[![Build Status](https://travis-ci.org/dannoane/priority-queue.svg?branch=master)](https://travis-ci.org/dannoane/priority-queue)

## Description
A generic priority queue implementation over a max heap.

## Usage

```go
package main

import (
    "fmt"
    "github.com/dannoane/priority-queue"
)

func main() {
    pq := priorityqueue.NewPriorityQueue()
    pq.Push(&priorityqueue.Element{Value: 1, Priority: 200})
    pq.Push(&priorityqueue.Element{Value: 2, Priority: 100})
    pq.Push(&priorityqueue.Element{Value: 3, Priority: 300})

    for i := 0; i < pq.Len(); i++ {
        fmt.Println(pq.Pop().Value)
    }
}
```
