// +build convey

package main

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func work(success bool) error {
	if !success {
		return errors.New("FAILURE")
	}

	return nil
}

func TestConvey(t *testing.T) {
	Convey("Some thing", t, func() {
		Convey("Should pass", func() {
			err := work(true)
			So(err, ShouldBeNil)
		})
		Convey("Should fail", func() {
			err := work(false)
			So(err, ShouldBeNil)
		})
	})
}
