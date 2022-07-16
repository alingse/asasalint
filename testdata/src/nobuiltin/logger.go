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

	l.Debug(a)        // want `pass \[\]any as any to func Debug func\(args \.\.\.interface{}\)`
	l.Debugf("%v", a) // want `pass \[\]any as any to func Debugf func\(format string, args \.\.\.interface{}\)`
	l.Debugln(a)      // want `pass \[\]any as any to func Debugln func\(args \.\.\.interface{}\)`

	l.Info(a)        // want `pass \[\]any as any to func Info func\(args \.\.\.interface{}\)`
	l.Infof("%v", a) // want `pass \[\]any as any to func Infof func\(format string, args \.\.\.interface{}\)`
	l.Infoln(a)      // want `pass \[\]any as any to func Infoln func\(args \.\.\.interface{}\)`

	l.Print(a)        // want `pass \[\]any as any to func Print func\(args \.\.\.interface{}\)`
	l.Printf("%v", a) // want `pass \[\]any as any to func Printf func\(format string, args \.\.\.interface{}\)`
	l.Println(a)      // want `pass \[\]any as any to func Println func\(args \.\.\.interface{}\)`

	l.Warn(a)        // want `pass \[\]any as any to func Warn func\(args \.\.\.interface{}\)`
	l.Warnf("%v", a) // want `pass \[\]any as any to func Warnf func\(format string, args \.\.\.interface{}\)`
	l.Warnln(a)      // want `pass \[\]any as any to func Warnln func\(args \.\.\.interface{}\)`

	l.Warning(a)        // want `pass \[\]any as any to func Warning func\(args \.\.\.interface{}\)`
	l.Warningf("%v", a) // want `pass \[\]any as any to func Warningf func\(format string, args \.\.\.interface{}\)`
	l.Warningln(a)      // want `pass \[\]any as any to func Warningln func\(args \.\.\.interface{}\)`

	l.Error(a)        // want `pass \[\]any as any to func Error func\(args \.\.\.interface{}\)`
	l.Errorf("%v", a) // want `pass \[\]any as any to func Errorf func\(format string, args \.\.\.interface{}\)`
	l.Errorln(a)      // want `pass \[\]any as any to func Errorln func\(args \.\.\.interface{}\)`

	l.Fatal(a)        // want `pass \[\]any as any to func Fatal func\(args \.\.\.interface{}\)`
	l.Fatalf("%v", a) // want `pass \[\]any as any to func Fatalf func\(format string, args \.\.\.interface{}\)`
	l.Fatalln(a)      // want `pass \[\]any as any to func Fatalln func\(args \.\.\.interface{}\)`

	l.Panic(a)        // want `pass \[\]any as any to func Panic func\(args \.\.\.interface{}\)`
	l.Panicf("%v", a) // want `pass \[\]any as any to func Panicf func\(format string, args \.\.\.interface{}\)`
	l.Panicln(a)      // want `pass \[\]any as any to func Panicln func\(args \.\.\.interface{}\)`
}
