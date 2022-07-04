package main

func getArgsLength(args ...any) int {
	return len(args)
}

func main() {
	var a = []any{1, 2, 3}
	getArgsLength(a)
	getArgsLength(a...)
}
