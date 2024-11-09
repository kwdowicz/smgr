# State Manager (smgr)

`smgr` is a Go package that provides a simple and flexible state management system. It allows developers to define states, manage transitions between states, and handle state-specific behavior through a clean interface-driven design.

## Features

- **State Definition**: Define states with `OnEnter`, `Update`, and `OnExit` functions to handle different phases of state lifecycle.
- **State Transitions**: Manage transitions between states and keep track of the previous state.
- **State Manager**: Use `StateManager` to control state transitions and handle the current active state.

## Installation

To install the package, use:

```bash
go get github.com/kwdowicz/smgr
```

## Usage

The package provides two key interfaces: `IState` and `IStateManager`.

- `IState` defines the behavior of a state.
- `IStateManager` manages the active state and transitions between states.

### Defining a State

A state can be defined using the `State` struct, and you can initialize it using the `NewState` function:

```go
import "github.com/kwdowicz/smgr"

onEnter := func() {
    fmt.Println("Entering state A")
}

update := func() {
    fmt.Println("Updating state A")
}

onExit := func() {
    fmt.Println("Exiting state A")
}

stateA := smgr.NewState(onEnter, update, onExit)
```

### Adding Next States

You can specify which states can be transitioned to from a given state:

```go
stateB := smgr.NewState(nil, nil, nil)
stateA.AddNextState(stateB)
```

### Managing States with StateManager

The `StateManager` helps manage transitions between states:

```go
stateManager := smgr.NewStateManager(stateA)

stateManager.Update() // Calls stateA's update function

if stateManager.NextState(stateB) {
    fmt.Println("Successfully transitioned to state B")
}
```

### Interfaces Overview

- **`IState`**
  - `Update()`: Invoked to perform the update logic of the state.
  - `OnEnter()`: Called when entering the state.
  - `OnExit()`: Called when exiting the state.
  - `AddNextState(IState)`: Adds a next state that this state can transition to.
  - `GetNextStates() []IState`: Returns a list of states that this state can transition to.
  - `SetPreviousState(IState)`: Sets the previous state.
  - `GetPreviousState() IState`: Gets the previous state.
  - `GetData() map[string]any`: Returns state-specific data.

- **`IStateManager`**
  - `Update()`: Updates the current state.
  - `NextState(IState) bool`: Attempts to transition to the given state.
  - `GetCurrentState() IState`: Returns the current state.

## Example

Here's a full example demonstrating how to create states and use the `StateManager` to handle transitions:

```go
package main

import (
    "fmt"
    "github.com/kwdowicz/smgr"
)

func main() {
    onEnterA := func() { fmt.Println("Entering State A") }
    updateA := func() { fmt.Println("Updating State A") }
    onExitA := func() { fmt.Println("Exiting State A") }

    stateA := smgr.NewState(onEnterA, updateA, onExitA)

    onEnterB := func() { fmt.Println("Entering State B") }
    stateB := smgr.NewState(onEnterB, nil, nil)

    stateA.AddNextState(stateB)

    stateManager := smgr.NewStateManager(stateA)

    stateManager.Update()  // Output: Updating State A
    if stateManager.NextState(stateB) {
        fmt.Println("Transitioned to State B")
    }
    stateManager.Update()  // No update function for State B, so no output
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributions

Contributions are welcome! Feel free to open issues or submit pull requests to help improve the package.

## Contact

For more information, contact [Kamil Wdowicz](https://github.com/kwdowicz).

---
Feel free to expand this further if you have additional use cases or suggestions in mind!

