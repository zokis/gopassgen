package gopassgen

import "testing"

type TestLen struct {
  pass string
  len_ int
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
