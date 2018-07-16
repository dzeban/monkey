package object

// Environment represents interpreter environment
// that stores bindings between object names and values
type Environment struct {
	store map[string]Object
	outer *Environment
}

// NewEnvironment construct new Environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment constructs the new Environment saving the pointer to
// the outer environment. This is used for evaluating function calls.
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}

// Get returns object value by a given name
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set creates a new binding between object name and its value
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
