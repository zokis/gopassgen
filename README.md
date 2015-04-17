# Password Generator

[![GoDoc](https://godoc.org/github.com/zokis/gopassgen?status.svg)](https://godoc.org/github.com/zokis/gopassgen)

a simple application to generate passwords

```go

import gpg "github.com/zokis/gopassgen"

// ...

newPass1 = gpg.NewPassword()
newPass2 = gpg.NewPassword(gpg.OptionChars([]rune("ABCDFHJLNPRTVXZbdfhjlnprtvxz")))
newPass3 = gpg.NewPassword(gpg.OptionLenght(15))
newPass4 = gpg.NewPassword(gpg.OptionChars([]rune("ABCDFHJLNPRTVXZbdfhjlnprtvxz")), gpg.OptionLenght(15))
newPass5 = gpg.NewPassword(gpg.OptionLenght(15), gpg.OptionChars([]rune("ABCDFHJLNPRTVXZbdfhjlnprtvxz")))

```
