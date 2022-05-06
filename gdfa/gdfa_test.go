package gdfa

import (
	"runtime/debug"
	"testing"
)

type testStates int

const (
	__ testStates = iota
	GET_FIST_CHAR
	GET_NEXT_CHAR
	GET_LAST_CHAR
	___
)

type localError string

func (e localError) Error() string {
	return string(e)
}

const errRejectedState = localError("machine got into rejected state")

func TestMachine(t *testing.T) {

	states := make(map[testStates]bool)
	for i := testStates(0); i < ___; i++ {
		states[i] = false
	}
	states[GET_LAST_CHAR] = true

	cursor := -2

	v := ""
	isNumber := false

	var next func(state testStates, input string) (testStates, error)

	next = func(state testStates, input string) (testStates, error) {
		cursor++
		switch state {
		case __:
			return next(GET_FIST_CHAR, input)
		case GET_FIST_CHAR:
			if input[cursor] == ':' {
				isNumber = true
				return next(GET_NEXT_CHAR, input)
			} else if input[cursor] != '$' {
				debug.PrintStack()
				return state, errRejectedState
			}
		case GET_NEXT_CHAR:
			if isNumber {
				if input[cursor] >= 48 && input[cursor] <= 57 {
					v += string(input[cursor])
				} else if input[cursor] == '\r' {
					return next(GET_LAST_CHAR, input)
				} else {
					debug.PrintStack()
					return state, errRejectedState
				}
			} else {
				v += string(input[cursor])
			}
			return next(GET_NEXT_CHAR, input)
		case GET_LAST_CHAR:
			if input[cursor] == '\n' {
				return state, nil
			}
		}
		debug.PrintStack()
		return state, errRejectedState
	}

	g, err := NewGDFA(":10000aa00\r\n", __, states, next)
	if err != nil {
		t.Error(err)
		return
	}

	if err = g.Process(); err != nil {
		t.Error("This should be rejected >:(")
	}
}
