package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestEvil(t *testing.T) {
	rand.Seed(time.Now().Unix())
	if rand.Intn(20) == 13 {
		t.Fail()
	}
}
