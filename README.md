# asasalint
Golang linter, lint that pass any slice as any in variadic function


## Install

```sh
go install github.com/alingse/asasalint/cmd/asasalint@latest
```

## Usage

```sh
asasalint ./...
```

## Why

two kind of unexpected usage, and `go build` success

```Go
package main

import "fmt"

func A(args ...any) int {
    return len(args)
}

func B(args ...any) int {
    return A(args)
}

func main() {

    // 1
    fmt.Println(B(1, 2, 3, 4))
}
```



```Go
package main

import "fmt"

func errMsg(msg string, args ...any) string {
    return fmt.Sprintf(msg, args...)
}

func Err(msg string, args ...any) string {
    return errMsg(msg, args)
}

func main() {

    // p1 [hello world] p2 %!s(MISSING)
    fmt.Println(Err("p1 %s p2 %s", "hello", "world"))
}
```



## TODO

1. add a setting struct (exclude/include/...)
2. add to golangci-lint
3. ignore in test.go file feature
4. given a SuggestEdition
5. add `append` to default exclude
