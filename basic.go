// Copyright 2019 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package randomex

import (
	"math/rand"
)

const (
	NumSet        = "0123456789"
	AlphaUpperSet = "ABCDEFGHIJKLMNOPQRSTUVXYZ"
	AlphaLowerSet = "abcdefghijklmnopqrstuvxyz"
	AlphaSet      = AlphaUpperSet + AlphaLowerSet
	AlphaNumSet   = NumSet + AlphaSet

	// PasswordSet is a common special character set used in passwords.
	// < and > may cause issues on some systems.
	PasswordSet = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

// Num returns a string containing a single digit random number.
func Num() string {
	return string(NumSet[rand.Intn(len(NumSet))])
}

// Nums returns a string of random numbers of "length".
func Nums(length int) string {
	if length < 1 {
		return ""
	}
	r := make([]byte, length)
	for i := 0; i < length; i++ {
		r[i] = byte(Num()[0])
	}
	return string(r)
}

// Upper returns a string containing a random uppercase letter.
func Upper() string {
	return string(AlphaUpperSet[rand.Intn(len(AlphaUpperSet))])
}

// Uppers returns a string of random uppercase letters of "length".
func Uppers(length int) string {
	if length < 1 {
		return ""
	}
	r := make([]byte, length)
	for i := 0; i < length; i++ {
		r[i] = byte(Upper()[0])
	}
	return string(r)
}

// Lower returns a string containing a random lowercase letter.
func Lower() string {
	return string(AlphaLowerSet[rand.Intn(len(AlphaLowerSet))])
}

// Lowers returns a string of random lowercase letters of "length".
func Lowers(length int) string {
	if length < 1 {
		return ""
	}
	r := make([]byte, length)
	for i := 0; i < length; i++ {
		r[i] = byte(Lower()[0])
	}
	return string(r)
}

// Special returns a string containing a random password special character.
func Special() string {
	return string(PasswordSet[rand.Intn(len(PasswordSet))])
}

// Specials returns a string of random special characters of "length".
func Specials(length int) string {
	if length < 1 {
		return ""
	}
	r := make([]byte, length)
	for i := 0; i < length; i++ {
		r[i] = byte(Special()[0])
	}
	return string(r)
}

// Random returns a random string of "length".
// If "lo" includes lowercase letters.
// If "up" includes uppercase letters.
// If "num" includes numbers.
// If "special" includes password special characters.
func Random(lo, up, nums, special bool, length int) string {
	if length < 1 {
		return ""
	}
	f := []func() string{}
	if lo {
		f = append(f, Lower)
	}
	if up {
		f = append(f, Upper)
	}
	if nums {
		f = append(f, Num)
	}
	if special {
		f = append(f, Special)
	}
	r := make([]byte, length)
	if lo || up || nums {
		for i := 0; i < length; i++ {
			r[i] = byte(f[rand.Intn(len(f))]()[0])
		}
	}
	return string(r)
}

// Rand is a shorthand for Random which includes all characters.
func Rand(length int) string {
	return Random(true, true, true, true, length)
}
