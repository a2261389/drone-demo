package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if add(3, 5) != 8 {
		t.Error("add not equal")
	}
}

func TestSub(t *testing.T) {
	if sub(10, 5) != 5 {
		t.Error("1. sub result no Expected")
	}
	if sub(5, 6) != -1 {
		t.Error("2. sub result no Expected")
	}
}

func TestMultiply(t *testing.T) {
	if multiply(10, 5) != 50 {
		t.Error("1. multiply result no Expected")
	}
	if multiply(0, 3) != 0 {
		t.Error("2. multiply result no Expected")
	}
	if multiply(-1, 7) != -7 {
		t.Error("3. multiply result no Expected")
	}
}

func TestDivided(t *testing.T) {
	if divided(30, 6) != 5 {
		t.Error("1. divided result no Expected")
	}
	if divided(7, 4) != 1.75 {
		t.Error("2. divided result no Expected")
	}
	if divided(-10, 2) != -5 {
		t.Error("3. divided result no Expected")
	}
}
