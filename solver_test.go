package main

import (
	"slices"
	"testing"
)

func TestSolve(t *testing.T) {
	results, err := Solve(200)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Contains(results, "98+76-5+43-2-10") {
		t.Fatal("example solution missing")
	}

}

func TestSolveConcatenation(t *testing.T) {
	results, err := Solve(9876543210)
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 || results[0] != "9876543210" {
		t.Fatalf("unexpected solutions: %v", results)
	}
}

func TestSolveImpossible(t *testing.T) {
	results, err := Solve(9876543211)
	if err == nil || results != nil {
		t.Fatalf("expected nil results and an error, got %v, %v", results, err)
	}
}
