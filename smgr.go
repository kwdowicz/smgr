package smgr 

type State struct{
	Update func()
	NextStates []*State
}

func (s *State) AddNextState(ns *State) {
	s.NextStates = append(s.NextStates, ns)
}

type StateManager struct {
	CurrentState *State
}

func (sm *StateManager) Update() {
	sm.CurrentState.Update()
}

func (sm *StateManager) NextState(ns *State) bool {
	for _, s := range sm.CurrentState.NextStates {
		if ns == s {
			sm.CurrentState = ns	
			return true
		}
	}
	return false
}

func NewStateManager(initialState *State) StateManager {
	sm := StateManager{
		CurrentState: initialState,
	}
	return sm
}