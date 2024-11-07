# State Manager (smgr)

`smgr` is a simple state management package in Go that allows you to manage states and transitions between them. The package is ideal for applications where you need to define various states and control the flow between them based on certain conditions.

## Features

- Define states with customizable behaviors
- Manage state transitions
- Track current state and trigger state-specific actions

## Installation

To install `smgr`, use `go get`:

```sh
go get github.com/kwdowicz/smgr
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/kwdowicz/smgr"
    "time"
)

func main() {
    // Define initial state
    s1 := smgr.State{}
    s1.Update = func() {
        fmt.Println("I'm State 1")
    }

    // Define another state
    s2 := smgr.State{}
    s2.Update = func() {
        fmt.Println("I'm State 2")
    }

    // Set up possible state transitions
    s1.AddNextState(&s2)
    s2.AddNextState(&s1)

    // Create a state manager and set the initial state
    sm := smgr.NewStateManager(&s1)

    // Loop through states with transitions
    count := 0
    for {
        fmt.Printf("Step %d: ", count)
        sm.Update()
        time.Sleep(1 * time.Second)
        
        // Transition between states based on custom logic
        if count%3 == 0 {
            sm.NextState(&s2)
        }
        if count%4 == 0 {
            sm.NextState(&s1)
        }
        count++
    }
}
```