package gopassgen

import "testing"
import "regexp"


type TestLen struct {
  pass string
  len_ int
}

type TestChar struct {
  pass     string
  alphabet []rune
  re       string
}


func TestLength(t *testing.T) {
  tests := []TestLen{
    TestLen{
      pass: NewPassword(),
      len_: 15,
    },
    TestLen{
      pass: NewPassword(OptionChars([]rune("ABCDEF."))),
      len_: 15,
    },
    TestLen{
      pass: NewPassword(OptionLength(16)),
      len_: 16,
    },
    TestLen{
      pass: NewPassword(OptionChars([]rune("abcdef.")),OptionLength(17)),
      len_: 17,
    },
    TestLen{
      pass: NewPassword(OptionLength(18), OptionChars([]rune("012345."))),
      len_: 18,
    },
  }

  for _, test := range tests {
    if test.len_ != len(test.pass){
      t.Errorf("Len Error %s != %s", test.len_, len(test.pass))
    }
  }
}

func TestChars(t *testing.T) {
  tests := []TestChar{
    TestChar{
      pass: NewPassword(),
      alphabet: StdChars,
      re: "^[ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$&(){}[\\]<>.:-_+=]+$",
    },
    TestChar{
      pass: NewPassword(OptionChars([]rune("ABCDEF."))),
      alphabet: []rune("ABCDEF."),
      re: "^[ABCDEF.]+$",
    },
    TestChar{
      pass: NewPassword(OptionChars([]rune("abcdef.")),OptionLength(17)),
      alphabet: []rune("abcdef."),
      re: "^[abcdef.]+$",
    },
    TestChar{
      pass: NewPassword(OptionLength(18), OptionChars([]rune("012345."))),
      alphabet: []rune("012345."),
      re: "^[012345.]+$",
    },
  }

  for _, test := range tests {
    alphabet_str := string(test.alphabet)
    r, _ := regexp.Compile(test.re)
    if !r.MatchString(test.pass) {
      t.Errorf("%s not match alphabet %s", test.pass, alphabet_str)
    }
  }
}