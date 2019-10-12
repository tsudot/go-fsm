# Finite State Machine for Go


## How to run

```bash
git get github.com/tsudot/go-fsm
cd $GOPATH/src/github.com/tsudot/go-fsm/

go build main.go

./main <input_str>
```

## Test cases

```bash
cd $GOPATH/src/github.com/tsudot/go-fsm/
go test ./... -v -cover
```


## Example

```golang

    package main

    import (
        "fmt"
        "github.com/tsudot/go-fsm/fsm"
    )

	sm := fsm.NewFSM("S0", "0")

	s0State := map[string]string{
		"0": "S0",
		"1": "S1",
	}

	s1State := map[string]string{
		"0": "S2",
		"1": "S0",
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

	sm.AddStateTransition(s0)
	sm.AddStateTransition(s1)

    input :=  "1100"
	err := sm.Transition(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("output for state %s = %s\n", sm.GetCurrent(), sm.GetCurrentStateValue())

```


