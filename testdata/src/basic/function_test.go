package main

import "testing"

func TestFunction(t *testing.T) {
	var a = []any{1, 2, 3}

	Debug(a)
	Debugf("%v", a)
	Debugln(a)

	Info(a)
	Infof("%v", a)
	Infoln(a)

	Print(a)
	Printf("%v", a)
	Println(a)

	Warn(a)
	Warnf("%v", a)
	Warnln(a)

	Warning(a)
	Warningf("%v", a)
	Warningln(a)

	Error(a)
	Errorf("%v", a)
	Errorln(a)

	Fatal(a)
	Fatalf("%v", a)
	Fatalln(a)

	Panic(a)
	Panicf("%v", a)
	Panicln(a)
}
