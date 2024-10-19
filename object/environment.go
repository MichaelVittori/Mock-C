package object

// Declared variables are saved in a hashmap-based environment, defined below

/*
 Constructor for enclosed environments, ex. environment of a function
 */
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

/*
 Constructor for an unenclosed environment, ex. the global environment
 */
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

/*
 Fetch from environment map
 If an outer map exists, check there if name not in self
 */
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil { // If the name cannot be found and there's an outer environemnt, check there
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

/*
 Add a value to the environment
 */
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}