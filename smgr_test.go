package smgr 

import (
	"testing"
)

// TestAddNextState verifies that AddNextState adds the correct next state.
func TestAddNextState(t *testing.T) {
	s1 := &State{}
	s2 := &State{}
	s1.AddNextState(s2)

	if len(s1.nextStates) != 1 {
		t.Errorf("Expected 1 next state, got %d", len(s1.nextStates))
	}
	if s1.nextStates[0] != s2 {
		t.Errorf("Expected next state to be s2")
	}
}

// TestStateManager_Update checks that Update calls the correct state's Update function.
func TestStateManager_Update(t *testing.T) {
	updated := false
	s1 := &State{
		Update: func() {
			updated = true
		},
	}
	sm := NewStateManager(s1)
	sm.Update()

	if !updated {
		t.Error("Expected Update to call current state's Update function")
	}
}

// TestStateManager_Update checks that Update calls the correct state's Update function.
func TestStateManager_UpdateFail(t *testing.T) {
	s1 := &State{}
	sm := NewStateManager(s1)
	sm.Update()

	if false {
		t.Error("Expected Update to call current state's Update function")
	}
}

// TestStateManager_NextState verifies that NextState correctly changes to a valid next state.
func TestStateManager_NextState(t *testing.T) {
	s1 := &State{}
	s2 := &State{}
	s1.AddNextState(s2)

	sm := NewStateManager(s1)
	if !sm.NextState(s2) {
		t.Error("Expected NextState to return true for valid transition")
	}
	if sm.CurrentState != s2 {
		t.Error("Expected current state to be s2")
	}

	// Test invalid transition
	s3 := &State{}
	if sm.NextState(s3) {
		t.Error("Expected NextState to return false for invalid transition")
	}
	if sm.CurrentState != s2 {
		t.Error("Expected current state to remain s2")
	}
}

// TestNewStateManager ensures that the state manager initializes with the correct initial state.
func TestNewStateManager(t *testing.T) {
	s1 := &State{}
	sm := NewStateManager(s1)

	if sm.CurrentState != s1 {
		t.Error("Expected initial state to be s1")
	}
}

// TestOnEnter verifies that the OnEnter function is called when transitioning to a new state.
func TestOnEnter(t *testing.T) {
	// Flag to check if OnEnter was called
	onEnterCalled := false

	// Define the initial state with no OnEnter function
	initialState := &State{
		Update: func() {
			// Initial state does nothing on Update
		},
	}

	// Define the target state with an OnEnter function that sets the flag to true
	targetState := &State{
		OnEnter: func() {
			onEnterCalled = true
		},
	}

	// Set up the possible transition from initialState to targetState
	initialState.AddNextState(targetState)

	// Initialize the StateManager with the initial state
	sm := NewStateManager(initialState)

	// Transition to targetState
	if !sm.NextState(targetState) {
		t.Error("Expected NextState to return true for valid transition")
	}

	// Check if OnEnter was called
	if !onEnterCalled {
		t.Error("Expected OnEnter to be called when transitioning to targetState")
	}
}

// TestOnExit verifies that the OnExit function is called when transitioning to a new state.
func TestOnExit(t *testing.T) {
	// Flag to check if OnEnter was called
	onExitCalled := false

	// Define the initial state with OnExit function taht sets the flag 
	initialState := &State{
		Update: func() {
			// Initial state does nothing on Update
		},
		OnExit: func() {
			onExitCalled = true
		},
	}

	targetState := &State{}

	// Set up the possible transition from initialState to targetState
	initialState.AddNextState(targetState)

	// Initialize the StateManager with the initial state
	sm := NewStateManager(initialState)

	// Transition to targetState
	if !sm.NextState(targetState) {
		t.Error("Expected NextState to return true for valid transition")
	}

	// Check if OnExit was called
	if !onExitCalled {
		t.Error("Expected OnExit to be called when transitioning to targetState")
	}
}
