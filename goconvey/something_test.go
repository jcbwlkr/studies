package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSomething(t *testing.T) {
	Convey("Something", t, func() {
		t.Log("a")

		defer func() {
			t.Log("in defer")
		}()

		t.Log("Will pass")
		So(1, ShouldEqual, 1)

		t.Log("Will fail")
		So(2, ShouldEqual, 1)
		So(3, ShouldEqual, 1)

		t.Log("Won't see")
	})
}
