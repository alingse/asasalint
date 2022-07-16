package main

import (
	"fmt"
	"log"
	"testing"
)

func TestStd(t *testing.T) {
	var a = []any{1, 2, 3}

	fmt.Println(a)      // want `pass \[\]any as any to func Println func\(a \.\.\.any\) \(n int, err error\)`
	fmt.Printf("%v", a) // want `pass \[\]any as any to func Printf func\(format string, a \.\.\.any\) \(n int, err error\)`
	fmt.Print(a)        // want `pass \[\]any as any to func Print func\(a \.\.\.any\) \(n int, err error\)`

	log.Println(a)      // want `pass \[\]any as any to func Println func\(v \.\.\.any\)`
	log.Printf("%v", a) // want `pass \[\]any as any to func Printf func\(format string, v \.\.\.any\)`
	log.Print(a)        // want `pass \[\]any as any to func Print func\(v \.\.\.any\)`

	log.Fatalln(a)      // want `pass \[\]any as any to func Fatalln func\(v \.\.\.any\)`
	log.Fatalf("%v", a) // want `pass \[\]any as any to func Fatalf func\(format string, v \.\.\.any\)`
	log.Fatal(a)        // want `pass \[\]any as any to func Fatal func\(v \.\.\.any\)`

	log.Panicln(a)      // want `pass \[\]any as any to func Panicln func\(v \.\.\.any\)`
	log.Panicf("%v", a) // want `pass \[\]any as any to func Panicf func\(format string, v \.\.\.any\)`
	log.Panic(a)        // want `pass \[\]any as any to func Panic func\(v \.\.\.any\)`
}
