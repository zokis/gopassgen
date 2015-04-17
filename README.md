# Password Generator

a simple application to generate passwords

```go

newPass1 = gopassgen.NewPassword()
newPass2 = gopassgen.NewPassword(gopassgen.OptionChars([]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")))
newPass3 = gopassgen.NewPassword(gopassgen.OptionLenght(15))
newPass4 = gopassgen.NewPassword(gopassgen.OptionChars([]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")), gopassgen.OptionLenght(15))
newPass5 = gopassgen.NewPassword(gopassgen.OptionLenght(15), gopassgen.OptionChars([]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")))

```
