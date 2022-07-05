package main

import "testing"

func TestGetArgsLength(t *testing.T) {
	var a = []any{1, 2, 3}
	if getArgsLength(a) != 3 {
		t.Errorf("getArgsLength(%v) != 3", a)
	}
	if getArgsLength(a...) != 3 {
		t.Errorf("getArgsLength(%v) != 3", a)
	}
	if getArgsLength(1, 2, 3) != 3 {
		t.Errorf("getArgsLength(%v) != 3", a)
	}
	if getArgsLength([]any{1, 2, 3}) != 3 {
		t.Errorf("getArgsLength(%v) != 3", a)
	}
	if getArgsLength(append([]any{1, 2, 3}, 4, 5, 6)) != 6 {
		t.Errorf("getArgsLength(%v) != 6", a)
	}
}
