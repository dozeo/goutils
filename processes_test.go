package goutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_execute0(t *testing.T) {
	Convey("execute 0", t, func() {
		out, _, err := Process.Execute(10, nil, "/bin/echo", "hi")
		So(string(out), ShouldEqual, "hi\n")
		So(err, ShouldEqual, nil)
	})
	Convey("execute 1", t, func() {
		_, _, err := Process.Execute(1, nil, "/bin/sleep", "2")
		So(err, ShouldNotEqual, nil)
	})
	Convey("execute 2", t, func() {
		out, _, err := Process.Execute(2, nil, "/bin/sleep", "1")
		So(string(out), ShouldEqual, "")
		So(err, ShouldEqual, nil)
	})
	Convey("execute 3", t, func() {
		_, _, err := Process.Execute(2, nil, "/bin/not_existent", "1")
		So(err, ShouldNotEqual, nil)
	})
	Convey("execute 4", t, func() {
		text := []byte("text")
		ret, stderr, err := Process.Execute(2, text, "/bin/cat", "-")
		So(err, ShouldEqual, nil)
		So(ret, ShouldNotEqual, nil)
		So(string(stderr), ShouldEqual, string([]byte{}))
		So(string(text), ShouldEqual, string(ret))
	})
}
