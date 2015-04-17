# Password Generator

[![GoDoc](https://godoc.org/github.com/zokis/gopassgen?status.svg)](https://godoc.org/github.com/zokis/gopassgen)

a simple application to generate passwords

```go

import gpg "github.com/zokis/gopassgen"

// ...

np1 = gpg.NewPassword()
np2 = gpg.NewPassword(gpg.OptionChars([]rune("ABCDEF...")))
np3 = gpg.NewPassword(gpg.OptionLenght(15))
np4 = gpg.NewPassword(gpg.OptionChars([]rune("ABCDEF...")), gpg.OptionLenght(15))
np5 = gpg.NewPassword(gpg.OptionLenght(15), gpg.OptionChars([]rune("ABCDEF...")))

```
