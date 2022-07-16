package main

import "testing"

func TestLogger(t *testing.T) {
	var l Logger

	var a = []any{1, 2, 3}

	l.Debug(a)
	l.Debugf("%v", a)
	l.Debugln(a)

	l.Info(a)
	l.Infof("%v", a)
	l.Infoln(a)

	l.Print(a)
	l.Printf("%v", a)
	l.Println(a)

	l.Warn(a)
	l.Warnf("%v", a)
	l.Warnln(a)

	l.Warning(a)
	l.Warningf("%v", a)
	l.Warningln(a)

	l.Error(a)
	l.Errorf("%v", a)
	l.Errorln(a)

	l.Fatal(a)
	l.Fatalf("%v", a)
	l.Fatalln(a)

	l.Panic(a)
	l.Panicf("%v", a)
	l.Panicln(a)
}
