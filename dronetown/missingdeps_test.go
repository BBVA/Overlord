package dronetown

import (
	"testing"
)

func testMissingDeps(t *testing.T, b, req []string, expect bool) {
	ok := hasMissingDeps(b, req)

	if ok != expect {
		t.Logf("Expect %t but %t found", expect, ok)
		t.FailNow()
	}
}

func TestMissingSimpleDep(t *testing.T) {
	b := []string{}
	req := []string{"tomato"}

	expect := true

	testMissingDeps(t, b, req, expect)
}

func TestMissingAllDeps(t *testing.T) {
	b := []string{}
	req := []string{"&", "2", "tomato", "oregano"}

	expect := true

	testMissingDeps(t, b, req, expect)
}

func TestMissingOneDep(t *testing.T) {
	b := []string{"tomato"}
	req := []string{"&", "2", "tomato", "oregano"}

	expect := true

	testMissingDeps(t, b, req, expect)
}

func TestMissingNoDeps(t *testing.T) {
	b := []string{"tomato", "oregano"}
	req := []string{"&", "2", "tomato", "oregano"}

	expect := false

	testMissingDeps(t, b, req, expect)
}

func TestMissingNoDepsTripleRainbow(t *testing.T) {
	b := []string{"tomato", "oregano", "salchipapa"}
	req := []string{"&", "3", "tomato", "oregano", "salchipapa"}

	expect := false

	testMissingDeps(t, b, req, expect)
}

func TestMissingDepsNested(t *testing.T) {
	b := []string{}
	req := []string{"&", "2", "pasta", "&", "2", "tomato", "oregano"}

	expect := true

	testMissingDeps(t, b, req, expect)
}

func TestMissingDepsAllTheSame(t *testing.T) {
	// AKA "Todo por la pasta"
	b := []string{}
	req := []string{"&", "5", "pasta", "pasta", "pasta", "pasta", "pasta"}

	expect := true

	testMissingDeps(t, b, req, expect)
}

func TestMissingDepsSimpleOr(t *testing.T) {
	b := []string{}
	req := []string{"|", "tomate", "nata"}

	expect := true

	testMissingDeps(t, b, req, expect)
}

func TestMissingDepsOrHasOneItem(t *testing.T) {
	b := []string{"tomate"}
	req := []string{"|", "tomate", "nata"}

	expect := false

	testMissingDeps(t, b, req, expect)
}

// TODO borrar & | y sus numeros
// deps [][]string
/*
	es posible que sea otro metodo, no missingdeps

	fetch & 2 tomato oregano
	fetch tomato fetch oregano

	fetch | 2 tomato oregano
	fetch tomato
	fetch oregano

	fetch & 2 pasta | 2 tomate nata
	fetch pasta fetch tomate
	fetch pasta fetch nata

	fetch | 2 pizza & 2 pasta & 2 tomato oregano
	fetch pizza
	fetch pasta fetch tomato fetch oregano

	combo breaker!!
	fetch | 2 & 2 pasta & 2 tomato oregano pizza
	fetch pasta fetch tomato fetch oregano
	fetch pizza

	extra extra
	fetch | 2 pizza & 2 pasta & 2 tomato oregano fetch pañales
	fetch pizza fetch pañales
	fetch pasta fetch tomato fetch oregano fetch pañales

	--- next
	bak: []
	go home fetch & 2 tomato oregano fetch pañales
	go home fetch tomato fetch oregano fetch pañales
	TODO: 1) change missing deps to return deps
	TODO: 2) create a method to return steps

*/
