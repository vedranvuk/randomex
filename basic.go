// Copyright 2019 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Basic alphanumeric string randomness. ASCII set. Not fast.

package randomex

import (
	"math/rand"
	"strings"
)

const (
	nums       = "0123456789"
	alphaUpper = "ABCDEFGHIJKLMNOPQRSTUVXYZ"
	alphaLower = "abcdefghijklmnopqrstuvxyz"
	alpha      = alphaUpper + alphaLower
	alphaNums  = nums + alpha

	// PasswordSet is a common special character set used in passwords.
	// < and > may cause issues on some systems.
	PasswordSet = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

// StringNum returns a string containing a single digit random number.
func StringNum() string {
	return string(nums[rand.Intn(len(nums))])
}

// StringNums returns a string of random numbers of "length".
func StringNums(length int) string {
	if length < 1 {
		return ""
	}
	r := make([]byte, length)
	for i := 0; i < length; i++ {
		r[i] = byte(StringNum()[0])
	}
	return string(r)
}

// StringUpper returns a string containing a random uppercase letter.
func StringUpper() string {
	return string(alphaUpper[rand.Intn(len(alphaUpper))])
}

// StringUppers returns a string of random uppercase letters of "length".
func StringUppers(length int) string {
	if length < 1 {
		return ""
	}
	r := make([]byte, length)
	for i := 0; i < length; i++ {
		r[i] = byte(StringUpper()[0])
	}
	return string(r)
}

// StringLower returns a string containing a random lowercase letter.
func StringLower() string {
	return string(alphaLower[rand.Intn(len(alphaLower))])
}

// StringLowers returns a string of random lowercase letters of "length".
func StringLowers(length int) string {
	if length < 1 {
		return ""
	}
	r := make([]byte, length)
	for i := 0; i < length; i++ {
		r[i] = byte(StringLower()[0])
	}
	return string(r)
}

// StringSpecial returns a string containing a random password special character.
func StringSpecial() string {
	return string(PasswordSet[rand.Intn(len(PasswordSet))])
}

// StringSpecials returns a string of random special characters of "length".
func StringSpecials(length int) string {
	if length < 1 {
		return ""
	}
	r := make([]byte, length)
	for i := 0; i < length; i++ {
		r[i] = byte(StringSpecial()[0])
	}
	return string(r)
}

// StringSet returns a single random character from the specified set.
func StringSet(set string) string {
	return string(set[rand.Intn(len(set))])
}

// String returns a random string of "length".
// If "lo" includes lowercase letters.
// If "up" includes uppercase letters.
// If "num" includes numbers.
// If "special" includes password special characters.
func String(lo, up, nums, special bool, length int) string {
	funcs := make([]func() string, 0, 4)
	if lo {
		funcs = append(funcs, StringLower)
	}
	if up {
		funcs = append(funcs, StringUpper)
	}
	if nums {
		funcs = append(funcs, StringNum)
	}
	if special {
		funcs = append(funcs, StringSpecial)
	}
	return StringOf(length, funcs...)
}

// StringOf takes any number of functins returning a string and calls them
// randomly until resulting string is of specified length.
func StringOf(length int, funcs ...func() string) string {
	if length < 1 {
		return ""
	}
	fs := make([]func() string, 0, len(funcs))
	for _, f := range funcs {
		fs = append(fs, f)
	}
	sb := strings.Builder{}
	for sb.Len() < length {
		sb.WriteString(fs[rand.Intn(len(fs))]())
	}
	return sb.String()[:length]
}

// Rand is a shorthand for Random which includes all characters.
func Rand(length int) string {
	return String(true, true, true, true, length)
}
