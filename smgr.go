package smgr

// IState interface defines the methods for managing state behavior.
type IState interface {
	Update()
	OnEnter()
	OnExit()
	AddNextState(IState)
	GetNextStates() []IState
	SetPreviousState(IState)
	GetPreviousState() IState
	GetData() map[string]any
}

// IStateManager interface defines methods to manage states and transitions.
type IStateManager interface {
	Update()
	NextState(IState) bool
	GetCurrentState() IState
}

// State implements the IState interface.
type State struct {
	update       func()
	onEnter      func()
	onExit       func()
	nextStates   []IState
	data         map[string]any
	previousState IState
}

// NewState creates and initializes a new State.
func NewState(onEnter, update, onExit func()) *State {
	return &State{
		update:     update,
		onEnter:    onEnter,
		onExit:     onExit,
		nextStates: make([]IState, 0),
		data:       make(map[string]any),
	}
}

// Update calls the update function if defined.
func (s *State) Update() {
	if s.update != nil {
		s.update()
	}
}

// OnEnter calls the onEnter function if defined.
func (s *State) OnEnter() {
	if s.onEnter != nil {
		s.onEnter()
	}
}

// OnExit calls the onExit function if defined.
func (s *State) OnExit() {
	if s.onExit != nil {
		s.onExit()
	}
}

// AddNextState adds a state to the nextStates list.
func (s *State) AddNextState(ns IState) {
	s.nextStates = append(s.nextStates, ns)
}

// GetNextStates returns the list of next states.
func (s *State) GetNextStates() []IState {
	return s.nextStates
}

// SetPreviousState sets the previous state.
func (s *State) SetPreviousState(prev IState) {
	s.previousState = prev
}

// GetPreviousState retrieves the previous state.
func (s *State) GetPreviousState() IState {
	return s.previousState
}

// GetData returns the state's data map.
func (s *State) GetData() map[string]any {
	return s.data
}

// StateManager implements the IStateManager interface.
type StateManager struct {
	currentState IState
}

// NewStateManager creates a new StateManager with an initial state.
func NewStateManager(initialState IState) *StateManager {
	return &StateManager{currentState: initialState}
}

// Update calls the Update method on the current state.
func (sm *StateManager) Update() {
	if sm.currentState != nil {
		sm.currentState.Update()
	}
}

// NextState transitions to a given next state if it's a valid transition.
func (sm *StateManager) NextState(ns IState) bool {
	for _, s := range sm.currentState.GetNextStates() {
		if ns == s {
			// Handle the transition
			holdCurrentState := sm.currentState
			sm.currentState.OnExit()
			sm.currentState = ns
			sm.currentState.SetPreviousState(holdCurrentState)
			sm.currentState.OnEnter()
			return true
		}
	}
	return false
}

// GetCurrentState returns the current state.
func (sm *StateManager) GetCurrentState() IState {
	return sm.currentState
}
