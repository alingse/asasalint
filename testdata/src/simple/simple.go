package main

func getArgsLength(args ...any) int {
	return len(args)
}

func getOneOrMore(arg any, others ...any) int {
	return len(append([]any{arg}, others...))
}
