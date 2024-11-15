// +build !convey

package main

import (
	"errors"
	"testing"
)

func work(success bool) error {
	if !success {
		return errors.New("FAILURE")
	}

	return nil
}

func TestNormalSuccess(t *testing.T) {
	err := work(true)
	if err != nil {
		t.Fatal("err should be nil, was", err)
	}
}

func TestNormalFail(t *testing.T) {
	err := work(false)
	if err != nil {
		t.Fatal("err should be nil, was", err)
	}
}
