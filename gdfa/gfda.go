package gdfa

type BasicStateConstraints interface {
	uint | ~int | string
}

type BasicStateConstraintsSlices[T BasicStateConstraints] interface {
	[]T | []*T
}

// I really don't remember what's this for
// But I'll keep it anyway, ya never know...
type Pair[F any, S any] struct {
	First  F
	Second S
}

// Defines the function call which takes a state and must return another state
// Or an error (in case of a Rejected state? Yeah for now)
type GDFANextFunc[S BasicStateConstraints, I any] func(state S, input I) (S, error)

// Struct for the DFA or FSM, whatever
// Only possible using StateCons
type GDFA[S BasicStateConstraints, I any] struct {
	input  I
	state  S
	states map[S]bool
	Next   GDFANextFunc[S, I]
}

func NewGDFA[S BasicStateConstraints, I any](input I, state S, states map[S]bool, next GDFANextFunc[S, I]) (*GDFA[S, I], error) {

	return &GDFA[S, I]{
		input:  input,
		state:  state,
		states: states,
		Next:   next,
	}, nil
}

func (g *GDFA[StateConstraints, any]) Process() error {
	for next, err := g.Next(g.state, g.input); err != nil && !g.states[next]; next, err = g.Next(next, g.input) {

	}

	return nil
}
