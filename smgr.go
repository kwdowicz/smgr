package smgr 

type State struct{
	Update func()
	OnEnter func()
	OnExit func()
	nextStates []*State
	Data map[string]any
	PreviousState *State
}

func (s *State) AddNextState(ns *State) {
	s.nextStates = append(s.nextStates, ns)
}

func NewState() *State {
	s := &State{}
	s.Data = make(map[string]any)
	return s
}

type StateManager struct {
	CurrentState *State
}

func (sm *StateManager) Update() {
	if sm.CurrentState != nil && sm.CurrentState.Update != nil {
		sm.CurrentState.Update()
	}
}

func (sm *StateManager) NextState(ns *State) bool {
	for _, s := range sm.CurrentState.nextStates {
		if ns == s {
			holdCurrentState := sm.CurrentState
			if sm.CurrentState.OnExit != nil {
				sm.CurrentState.OnExit()
			}
			sm.CurrentState = ns	
			sm.CurrentState.PreviousState = holdCurrentState 
			if sm.CurrentState.OnEnter != nil {
				sm.CurrentState.OnEnter()
			}
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

