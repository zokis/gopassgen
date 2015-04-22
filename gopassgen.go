package gopassgen

import (
  "math/rand"
  "math"
  "time"
)

const (
  minPasswordLength = 8
  maxPasswordLength = 30
)

// StdChars default alphabet
var StdChars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$&(){}[]<>.:-_+=")

// NPParams parameters
type NPParams struct {
  chars  []rune
  length int
}

// OptionLength adds the option of password length
func OptionLength(l int) func(p *NPParams) error {
  return func(p *NPParams) error {
    p.length = l
    return nil
  }
}

// OptionChars adds the alphabet option used
func OptionChars(c []rune) func(p *NPParams) error {
  return func(p *NPParams) error {
    p.chars = c
    return nil
  }
}

func randInt(min int, max int) int {
  rand.Seed(time.Now().UTC().UnixNano())
  return min + rand.Intn(max - min)
}


// NewPassword generates a new password
func NewPassword(options ...func(*NPParams) error) string {

    rand.Seed( time.Now().UTC().UnixNano())

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
        newPassword[i] = p.chars[randInt(0, len(p.chars))]
    }
    return string(newPassword)
}
