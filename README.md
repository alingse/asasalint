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

ignore some func that was by desgin

```sh
asasalint -e append,Append ./...
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

1. given a SuggestEdition
2. add `append` to default exclude ?
3. ingore pattern `fn(a, b, []any{1,2,3})` , because `[]any{1,2,3}` is most likely by design
