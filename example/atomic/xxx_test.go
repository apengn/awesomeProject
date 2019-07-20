package main

import "testing"

func get() string {
	return "www"
}
func TestDome(t *testing.T) {
	want := "wwsw"

	if got := get(); got != want {
		t.Errorf("get()=%q,want=%q", want, got)
	}
}
