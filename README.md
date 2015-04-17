# Password Generator

[![GoDoc](https://godoc.org/github.com/zokis/gopassgen?status.svg)](https://godoc.org/github.com/zokis/gopassgen)

a simple application to generate passwords

```go
package main

import f "fmt"
import gpg "github.com/zokis/gopassgen"

func main() {
    np1 := gpg.NewPassword()
    np2 := gpg.NewPassword(gpg.OptionChars([]rune("ABCDEF.")))
    np3 := gpg.NewPassword(gpg.OptionLenght(15))
    np4 := gpg.NewPassword(gpg.OptionChars([]rune("abcdef...")), gpg.OptionLenght(15))
    np5 := gpg.NewPassword(gpg.OptionLenght(15), gpg.OptionChars([]rune("012345...")))
    f.Println(np1)
    f.Println(np2)
    f.Println(np3)
    f.Println(np4)
    f.Println(np5)
}
```
