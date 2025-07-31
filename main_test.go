package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a, b := 2, 3
	want := a+b
	if want != add(a, b){
		t.Errorf("wanted %d got %d", want, add(a,b))
	}
}
