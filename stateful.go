package stateful

import (
	"errors"
	"reflect"
)

var (
	errNotPtr            = errors.New("expected kind to be of type reflect ptr")
	errNotFunc           = errors.New("expected kind to be of type reflect func")
	errUnsetable         = errors.New("cannot set embeded Stateful field")
	errInvalidTransition = errors.New("cannot transition from current state to the state provided")
)

type State uint

type StateMachine struct {
	initial  State
	current  State
	previous State
	object   reflect.Value
	before   map[State]map[State][]reflect.Value
	after    map[State]map[State][]reflect.Value
}

func (s *StateMachine) BeforeTransition(from, to State, fn interface{}) *StateMachine {
	value := reflect.ValueOf(fn)

	if kind := value.Kind(); reflect.Func != kind {
		panic(errNotFunc)
	} else if nil == s.before[from] {
		s.before[from] = make(map[State][]reflect.Value)
	}

	s.before[from][to] = append(s.before[from][to], value)
	return s
}

func (s *StateMachine) AfterTransition(from, to State, fn interface{}) *StateMachine {
	value := reflect.ValueOf(fn)

	if kind := value.Kind(); reflect.Func != kind {
		panic(errNotFunc)
	} else if nil == s.after[from] {
		s.after[from] = make(map[State][]reflect.Value)
	}

	s.after[from][to] = append(s.after[from][to], value)
	return nil
}

func (s *StateMachine) Transition(to State) (*StateMachine, error) {
	if ok := s.CanTransitionTo(to); !ok {
		return nil, errInvalidTransition
	}

	for _, fn := range s.before[s.current][to] {
		fn.Call([]reflect.Value{s.object})
	}

	s.previous = s.current
	s.current = to

	for _, fn := range s.after[s.previous][to] {
		fn.Call([]reflect.Value{s.object})
	}

	return s, nil
}

func (s *StateMachine) CanTransitionTo(to State) bool {
	if _, ok := s.before[s.current][to]; ok {
		return true
	} else if _, ok := s.after[s.current][to]; ok {
		return true
	}

	return false
}

func NewStateMachine(initial State) (s *StateMachine) {
	s = new(StateMachine)
	s.initial = initial
	s.current = initial
	s.before = make(map[State]map[State][]reflect.Value)
	s.after = make(map[State]map[State][]reflect.Value)
	return
}

func Process(typ interface{}, embed *StateMachine) (interface{}, error) {
	value := reflect.ValueOf(typ)

	if kind := value.Kind(); reflect.Ptr != kind {
		return nil, errNotPtr
	}

	element := value.Elem()
	field := element.FieldByName("StateMachine")

	if kind := field.Kind(); reflect.Ptr != kind {
		return nil, errNotPtr
	}

	if field.CanSet() {
		field.Set(reflect.ValueOf(embed))
	} else {
		return nil, errUnsetable
	}

	embed.object = reflect.ValueOf(typ)

	return value.Interface(), nil
}
