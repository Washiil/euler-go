package problems

var registry = make(map[string]func() int)

func Register(id string, fn func() int) {
	registry[id] = fn
}

func GetProblems() map[string]func() int {
	return registry
}
