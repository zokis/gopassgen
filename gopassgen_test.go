package gopassgen

import "testing"

type Test struct {
  pass string
  len_ int
}

func TestLen(t *testing.T) {
  t1 := Test{pass: NewPassword(), len_: 15}
  t2 := Test{pass: NewPassword(OptionChars([]rune("ABCDEF."))),len_: 15}
  t3 := Test{pass: NewPassword(OptionLength(16)),len_: 16}
  t4 := Test{pass: NewPassword(OptionChars([]rune("abcdef.")), OptionLength(17)),len_: 17}
  t5 := Test{pass: NewPassword(OptionLength(18), OptionChars([]rune("012345."))),len_: 18}
  tests := []Test{t1, t2, t3, t4, t5}

  for _, test := range tests {
    if test.len_ != len(test.pass){
      t.Errorf("Len Error %s != %s", test.len_, len(test.pass))
    }
  }
}
