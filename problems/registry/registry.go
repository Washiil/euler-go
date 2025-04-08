// problems/registry/registry.go
package registry

var Problems = make(map[string]func() int)

func Register(n string, fn func() int) {
	Problems[n] = fn
}
