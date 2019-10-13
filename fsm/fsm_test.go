package fsm

import (
	"testing"
)

func TestFSM(t *testing.T) {

	tables := []struct {
		input      string
		finalState string
	}{
		{"110", "S0"},
		{"1010", "S1"},
		{"100011011", "S1"},
		{"100011", "S2"},
	}

	for _, table := range tables {
		sm := generateStateMachine()

		err := sm.Transition(table.input)

		if err != nil {
			t.Error(err)
			continue
		}

		if sm.CurrentState != table.finalState {
			t.Errorf("Final State is incorrect, got: %s, want: %s.", sm.CurrentState, table.finalState)
		}
	}
}

func TestFSMInvalidState(t *testing.T) {
	tables := []struct {
		input      string
		finalState string
	}{
		{"112", "Encountered invalid input: 2 for current state: S0"},
		{"1013", "Encountered invalid input: 3 for current state: S2"},
	}

	for _, table := range tables {
		sm := generateStateMachine()

		err := sm.Transition(table.input)

		if err != nil {
			if table.finalState != err.Error() {
				t.Errorf("Got incorrect error, got: %s, want: %s.", err.Error(), table.finalState)
			}
		} else {
			t.Errorf("The input should result in error, input: %s", table.input)
		}
	}

}

func generateStateMachine() *FSM {
	sm := NewFSM("S0", "0")

	s0State := map[string]string{
		"0": "S0",
		"1": "S1",
	}

	s1State := map[string]string{
		"0": "S2",
		"1": "S0",
	}

	s2State := map[string]string{
		"0": "S1",
		"1": "S2",
	}

	// State Transition
	s0 := StateTransition{
		State:      "S0",
		StateTable: s0State,
	}

	s1 := StateTransition{
		State:      "S1",
		StateTable: s1State,
	}

	s2 := StateTransition{
		State:      "S2",
		StateTable: s2State,
	}

	sm.AddStateTransition(s0)
	sm.AddStateTransition(s1)
	sm.AddStateTransition(s2)

	return sm

}
