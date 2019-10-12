package fsm

import (
	"errors"
	"fmt"
)

type StateTransition struct {
	// State in the table
	State string

	// Map of input and output per state
	StateTable map[string]string
}

type FSM struct {
	// Current State of the machine
	CurrentState string

	// Array of StateTransition
	StateTransition []StateTransition

	// Current state value, can be different from the CurrentState
	CurrentStateValue string
}

/*
Returns a new state machine
*/
func NewFSM(initialState string, stateValue string) *FSM {
	fsm := &FSM{}
	fsm.CurrentState = initialState
	fsm.CurrentStateValue = stateValue
	fsm.StateTransition = make([]StateTransition, 0)

	return fsm
}

/*
Adds StateTransistion to the state machine
*/

func (f *FSM) AddStateTransition(st StateTransition) {
	f.StateTransition = append(f.StateTransition, st)
}

/*
Takes an input string and changes the internal
state of the state machine.

Returns an error if an invalid state is encountered
*/
func (f *FSM) Transition(input string) error {
	for _, i := range input {
		in := string(i)
		currentState, err := f.nextState(in)
		if err != nil {
			return err
		}
		f.CurrentState = currentState
		f.CurrentStateValue = in
	}
	return nil
}

/*

Transition will call this internal function with
sliced input from the input from the input string
to get the next state

Returns an error on an invalid input string

*/
func (f *FSM) nextState(input string) (string, error) {
	for _, v := range f.StateTransition {
		if v.State == f.CurrentState {
			if val, ok := v.StateTable[input]; ok {
				return val, nil
			} else {
				return "-", errors.New(
					fmt.Sprintf("Encountered invalid input: %s for current state: %s", input,
						f.CurrentState))
			}
		}
	}
	return "-", errors.New("Invalid state transition")
}

// Returns the current state
func (f *FSM) GetCurrent() string {
	return f.CurrentState
}

// Returns the current state value
func (f *FSM) GetCurrentStateValue() string {
	return f.CurrentStateValue
}
