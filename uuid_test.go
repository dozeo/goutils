package goutils

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_UUID(t *testing.T) {
	Convey("generate uuid", t, func() {
		uuid := UUID()
		So(uuid, ShouldNotEqual, nil)
		So(strings.Count(uuid, "-"), ShouldEqual, 4)
		parts := strings.Split(uuid, "-")
		So(len(parts[0]), ShouldEqual, 8)
		So(len(parts[1]), ShouldEqual, 4)
		So(len(parts[2]), ShouldEqual, 4)
		So(len(parts[3]), ShouldEqual, 4)
		So(len(parts[4]), ShouldEqual, 12)
	})
}
