package smgr_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/kwdowicz/smgr"
)

func TestState_OnEnter(t *testing.T) {
	var called bool
	state := smgr.NewState(func() { called = true }, nil, nil)

	state.OnEnter()

	assert.True(t, called, "OnEnter should call the provided onEnter function")
}

func TestState_Update(t *testing.T) {
	var called bool
	state := smgr.NewState(nil, func() { called = true }, nil)

	state.Update()

	assert.True(t, called, "Update should call the provided update function")
}

func TestState_OnExit(t *testing.T) {
	var called bool
	state := smgr.NewState(nil, nil, func() { called = true })

	state.OnExit()

	assert.True(t, called, "OnExit should call the provided onExit function")
}

func TestState_AddNextState(t *testing.T) {
	stateA := smgr.NewState(nil, nil, nil)
	stateB := smgr.NewState(nil, nil, nil)

	stateA.AddNextState(stateB)

	nextStates := stateA.GetNextStates()
	assert.Len(t, nextStates, 1, "NextStates should contain one state")
	assert.Equal(t, stateB, nextStates[0], "NextState should be equal to the added state")
}

func TestState_SetGetPreviousState(t *testing.T) {
	stateA := smgr.NewState(nil, nil, nil)
	stateB := smgr.NewState(nil, nil, nil)

	stateB.SetPreviousState(stateA)

	assert.Equal(t, stateA, stateB.GetPreviousState(), "PreviousState should be equal to the set state")
}

func TestState_GetData(t *testing.T) {
	state := smgr.NewState(nil, nil, nil)
	data := state.GetData()
	assert.NotNil(t, data, "GetData should return a non-nil map")

	data["key"] = "value"
	assert.Equal(t, "value", state.GetData()["key"], "Data map should persist values correctly")
}

func TestStateManager_Update(t *testing.T) {
	var called bool
	state := smgr.NewState(nil, func() { called = true }, nil)
	stateManager := smgr.NewStateManager(state)

	stateManager.Update()

	assert.True(t, called, "Update should call the current state's Update function")
}

func TestStateManager_NextState_ValidTransition(t *testing.T) {
	stateA := smgr.NewState(nil, nil, nil)
	stateB := smgr.NewState(nil, nil, nil)
	stateA.AddNextState(stateB)
	stateManager := smgr.NewStateManager(stateA)

	transitioned := stateManager.NextState(stateB)

	assert.True(t, transitioned, "NextState should return true for a valid transition")
	assert.Equal(t, stateB, stateManager.GetCurrentState(), "Current state should be updated to the next state")
	assert.Equal(t, stateA, stateB.GetPreviousState(), "Previous state of the new state should be set correctly")
}

func TestStateManager_NextState_InvalidTransition(t *testing.T) {
	stateA := smgr.NewState(nil, nil, nil)
	stateB := smgr.NewState(nil, nil, nil)
	stateManager := smgr.NewStateManager(stateA)

	transitioned := stateManager.NextState(stateB)

	assert.False(t, transitioned, "NextState should return false for an invalid transition")
	assert.Equal(t, stateA, stateManager.GetCurrentState(), "Current state should remain the same if the transition is invalid")
}

func TestStateManager_GetCurrentState(t *testing.T) {
	state := smgr.NewState(nil, nil, nil)
	stateManager := smgr.NewStateManager(state)

	assert.Equal(t, state, stateManager.GetCurrentState(), "GetCurrentState should return the current state")
}
