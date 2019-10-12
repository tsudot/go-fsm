package main

import (
	"fmt"
	"github.com/tsudot/go-fsm/fsm"
	"os"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Run with input string; For eg. go run main.go 110")
		return
	}

	inputStr := args[1]

	sm := fsm.NewFSM("S0", "0")

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
	s0 := fsm.StateTransition{
		State:      "S0",
		StateTable: s0State,
	}

	s1 := fsm.StateTransition{
		State:      "S1",
		StateTable: s1State,
	}

	s2 := fsm.StateTransition{
		State:      "S2",
		StateTable: s2State,
	}

	sm.AddStateTransition(s0)
	sm.AddStateTransition(s1)
	sm.AddStateTransition(s2)

	err := sm.Transition(inputStr)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("output for state %s = %s\n", sm.GetCurrent(), sm.GetCurrentStateValue())

}
