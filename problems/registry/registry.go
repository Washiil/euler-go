package registry

// Problems holds a map of registered problem functions
var Problems = make(map[string]func() any)

// Register wraps a strongly typed func() T into func() any
func Register[T any](name string, fn func() T) {
	Problems[name] = func() any {
		println("Registering " + name)
		return fn()
	}
}
