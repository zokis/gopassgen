package gopassgen

import (
  "math/rand"
  "math"
  "fmt"
)

const (
  minPasswordLength = 8
  maxPasswordLength = 30
)

var StdChars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$&(){}[]<>.:-_+=")

type NPParams struct {
  chars  []rune
  length int
}

func OptionLenght(l int) func(p *NPParams) error {
  return func(p *NPParams) error {
    p.length = l
    return nil
  }
}

func OptionChars(c []rune) func(p *NPParams) error {
  return func(p *NPParams) error {
    p.chars = c
    return nil
  }
}

func NewPassword(options ...func(*NPParams) error) string {
    p := &NPParams{}
    p.chars = StdChars
    p.length = 15

    for _, op := range options {
      err := op(p)
      if err != nil {
        panic("invalid options")
      }
    }

    newPassword := make([]rune, int(math.Min(math.Max(float64(p.length), minPasswordLength), maxPasswordLength)))
    for i := range newPassword {
        newPassword[i] = p.chars[rand.Intn(len(p.chars))]
    }
    return string(newPassword), nil
}
