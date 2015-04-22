package gopassgen

import "testing"
import "strings"

type TestLen struct {
  pass string
  len_ int
}

type TestChar struct {
  pass     string
  alphabet []rune
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
    },
    TestChar{
      pass: NewPassword(OptionChars([]rune("ABCDEF."))),
      alphabet: []rune("ABCDEF."),
    },
    TestChar{
      pass: NewPassword(OptionChars([]rune("abcdef.")),OptionLength(17)),
      alphabet: []rune("abcdef."),
    },
    TestChar{
      pass: NewPassword(OptionLength(18), OptionChars([]rune("012345."))),
      alphabet: []rune("12345."),
    },
  }

  for _, test := range tests {
    for _, c := range []byte(test.pass) {
      
      if strings.Contains(string(c), string(test.alphabet)) {
        t.Errorf("%s not in %s", c, test.alphabet)
      }
    }
  }

}