package main

import "testing"

func TestNewSet(t *testing.T) {
	set := NewSet[string](4)
	if len(set) != 0 {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	set := NewSet[int](8)
	set.Add(1, 2, 3)

	if len(set) != 2 {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	set := NewSet[int](8)
	set.Add(1, 2)
	set.Add(2)

	if len(set) != 1 {
		t.Fail()
	}
}
