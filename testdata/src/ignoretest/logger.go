package main

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})

	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})

	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Warningln(args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})
}

func fromLogger(l Logger) {
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
