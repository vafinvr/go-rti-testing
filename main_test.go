package main

import "testing"

func TestCalculateNil(t *testing.T) {
	r, err := Calculate(nil, nil)
	if err != nil {
		t.Error("Error calculating", err)
	}
	if r != nil {
		t.Error(r)
	}
}

func TestCalculateNotNil(t *testing.T) {
	r, err := Calculate(&Product{}, &Condition{})
	if err != nil {
		t.Error("Error calculating", err)
	}
	if r == nil {
		t.Error(r)
	}
}
