package main

func Debug(args ...interface{})                 {}
func Debugf(format string, args ...interface{}) {}
func Debugln(args ...interface{})               {}

func Info(args ...interface{})                 {}
func Infof(format string, args ...interface{}) {}
func Infoln(args ...interface{})

func Print(args ...interface{})                 {}
func Printf(format string, args ...interface{}) {}
func Println(args ...interface{})               {}

func Warn(args ...interface{})                 {}
func Warnf(format string, args ...interface{}) {}
func Warnln(args ...interface{})               {}

func Warning(args ...interface{})                 {}
func Warningf(format string, args ...interface{}) {}
func Warningln(args ...interface{})               {}

func Error(args ...interface{})                 {}
func Errorf(format string, args ...interface{}) {}
func Errorln(args ...interface{})               {}

func Fatal(args ...interface{})                 {}
func Fatalf(format string, args ...interface{}) {}
func Fatalln(args ...interface{})               {}

func Panic(args ...interface{})                 {}
func Panicf(format string, args ...interface{}) {}
func Panicln(args ...interface{})               {}
