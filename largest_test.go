package main

import "testing"

func TestLargest(t *testing.T){
	actual := Largest([]int{1,2,3,4})
	expected := 4
	if actual != expected{
		t.Error("expected",expected,"but got",actual)
	}
}
func TestLargest1(t *testing.T){
	actual := Largest([]int{5,6,7,8})
	expected := 8
	if actual != expected{
		t.Error("expected",expected,"but got",actual)
	}
}
func Testsmallest(t *testing.T){
	actual := Smallest([]int{1,2,3,4})
	expected:=1
	if actual!=expected{
		t.Error("expected",expected,"but got",actual)
	}
}