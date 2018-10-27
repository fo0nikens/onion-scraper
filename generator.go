package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	addressLengthV2 = 16
	addressLengthV3 = 56
)

var charList = []rune("abcdefghijklmnopqrstuvwxyz234567")

// Address for onion
type Address string

// Addr returns http url
func (a Address) Addr() string {
	return fmt.Sprintf("http://%s.onion", a)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandOnionV2 generates onion service version 2
func RandOnionV2() Address {
	b := make([]rune, addressLengthV2)
	for i := range b {
		b[i] = charList[rand.Intn(len(charList))]
	}
	return Address(string(b))
}

// RandOnionV3 generates onion service version 3
func RandOnionV3() Address {
	b := make([]rune, addressLengthV3)
	for i := range b {
		b[i] = charList[rand.Intn(len(charList))]
	}
	return Address(string(b))
}
