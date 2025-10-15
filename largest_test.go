package main

import "testing"

func TestLargest(t *testing.T){
	actual := Largest([]int{1,2,3,4})
	expected := 4
	if actual != expected{
		t.Error("expected",expected,"but got",actual)
	}
}