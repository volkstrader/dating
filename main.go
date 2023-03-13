package main

import (
	"fmt"
	"github.com/volkstrader/actress"
	"github.com/volkstrader/model"
)

func main() {
	const name = "barbie"
	const age = 24

	a := actress.Profile{
		Name: name,
		Age:  age,
	}

	m := model.Profile{
		Name: name,
		Age:  age,
	}

	dest, ok := bookVacation(a)
	if !ok {
		fmt.Printf("not dating with actress %s\n", a.Name)
	} else {
		fmt.Printf("going to %s with actress %s\n", dest, a.Name)
	}

	dest, ok = bookVacation(m)
	if !ok {
		fmt.Printf("not dating with model %s\n", a.Name)
	} else {
		fmt.Printf("going to %s with model %s\n", dest, a.Name)
	}

}

type partner interface {
	Dating() bool
}

func bookVacation(p partner) (string, bool) {
	if !p.Dating() {
		return "", false
	}

	return "Hawaii", true
}
