package pierre

import "strings"

// Human represents a human
type Human interface {
	Name() string
}

// Pet represents a pet
type Pet interface {
	Name() string
}

type human struct {
	name string
}

// Hello ...
func Hello() (string, string) {
	h := human{name: "Pierre"}
	h.name = "Julien"
	return "Hello", h.toUpperName()
}

func (h human) toUpperName() string {
	return strings.ToUpper(h.name)
}
