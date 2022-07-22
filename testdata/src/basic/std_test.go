package main

import (
	"fmt"
	"io"
	"log"
	"testing"
)

func TestStd(t *testing.T) {
	var a = []any{1, 2, 3}

	fmt.Println(a)
	fmt.Printf("%v", a)
	fmt.Print(a)

	fmt.Sprintln(a)
	fmt.Sprintf("%v", a)
	fmt.Sprint(a)

	var w io.Writer
	fmt.Fprintln(w, a)
	fmt.Fprintf(w, "%v", a)
	fmt.Fprint(w, a)

	log.Println(a)
	log.Printf("%v", a)
	log.Print(a)

	log.Fatalln(a)
	log.Fatalf("%v", a)
	log.Fatal(a)

	log.Panicln(a)
	log.Panicf("%v", a)
	log.Panic(a)
}

func TestTest(t *testing.T) {
	var a = []any{1, 2, 3}
	if len(a) != 3 {
		t.Log(a)
		t.Logf("%v", a)
		t.Error(a)
		t.Errorf("%v", a)
		t.Fatal(a)
		t.Fatalf("%v", a)
	}

	var tt = t
	if len(a) != 3 {
		tt.Log(a)          // want `pass \[\]any as any to func tt.Log func\(args \.\.\.any\)`
		tt.Logf("%v", a)   // want `pass \[\]any as any to func tt.Logf func\(format string, args \.\.\.any\)`
		tt.Error(a)        // want `pass \[\]any as any to func tt.Error func\(args \.\.\.any\)`
		tt.Errorf("%v", a) // want `pass \[\]any as any to func tt.Errorf func\(format string, args \.\.\.any\)`
		tt.Fatal(a)        // want `pass \[\]any as any to func tt.Fatal func\(args \.\.\.any\)`
		tt.Fatalf("%v", a) // want `pass \[\]any as any to func tt.Fatalf func\(format string, args \.\.\.any\)`
	}
}
