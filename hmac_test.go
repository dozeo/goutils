package goutils

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"os/exec"
	"testing"
)

func MyConvey(items ...interface{}) {
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("-- " + items[0].(string))
	fmt.Println("-------------------------------------------------------------------------")
	Convey(items...)
}

func Test_HMAC(t *testing.T) {
	MyConvey("URL 1", t, func() {
		url := "http://www.eklenet.de"
		pass := "geheim"
		out, err := exec.Command("php5", "hmac_test/ref.php", url, pass).Output()
		So(nil, ShouldEqual, err)
		So(signUrl(url, pass), ShouldEqual, string(out))
		So(true, ShouldEqual, validate(string(out), pass))
	})
	MyConvey("URL 2", t, func() {
		url := "http://www.eklenet.de/"
		pass := "geheim"
		out, err := exec.Command("php5", "hmac_test/ref.php", url, pass).Output()
		So(nil, ShouldEqual, err)
		So(signUrl(url, pass), ShouldEqual, string(out))
		So(true, ShouldEqual, validate(string(out), pass))
	})
	MyConvey("URL 3", t, func() {
		url := "http://www.eklenet.de/abc"
		pass := "geheim"
		out, err := exec.Command("php5", "hmac_test/ref.php", url, pass).Output()
		So(nil, ShouldEqual, err)
		So(signUrl(url, pass), ShouldEqual, string(out))
		So(true, ShouldEqual, validate(string(out), pass))
	})
	MyConvey("URL 4", t, func() {
		url := "http://www.eklenet.de/abc/"
		pass := "geheim"
		out, err := exec.Command("php5", "hmac_test/ref.php", url, pass).Output()
		So(nil, ShouldEqual, err)
		So(signUrl(url, pass), ShouldEqual, string(out))
		So(true, ShouldEqual, validate(string(out), pass))
	})
	MyConvey("URL 5", t, func() {
		url := "http://www.eklenet.de/abc/def?xyz=123&ok=no"
		pass := "geheim"
		out, err := exec.Command("php5", "hmac_test/ref.php", url, pass).Output()
		So(nil, ShouldEqual, err)
		So(signUrl(url, pass), ShouldEqual, string(out))
		So(true, ShouldEqual, validate(string(out), pass))
	})
	MyConvey("URL 6", t, func() {
		url := "http://www.eklenet.de/abc/def?ok=no&xyz=123"
		pass := "geheim"
		out, err := exec.Command("php5", "hmac_test/ref.php", url, pass).Output()
		So(nil, ShouldEqual, err)
		So(signUrl(url, pass), ShouldEqual, string(out))
		So(true, ShouldEqual, validate(string(out), pass))
	})
	MyConvey("URL 7", t, func() {
		url := "http://www.eklenet.de/abc/def?ok=no&xyz=123#what"
		pass := "geheim"
		out, err := exec.Command("php5", "hmac_test/ref.php", url, pass).Output()
		So(nil, ShouldEqual, err)
		So(signUrl(url, pass), ShouldEqual, string(out))
		So(true, ShouldEqual, validate(string(out), pass))
	})
	MyConvey("URL 8", t, func() {
		url := "http://www.eklenet.de/abc/def?ok=no&xyz=123#what"
		pass := "geheim"
		surl := signUrl(url, pass)
		So(true, ShouldEqual, validate(surl, pass))
	})
	MyConvey("URL 9", t, func() {
		url := "http://www.eklenet.de/abc/def?ok=no&xyz=123#what"
		pass := "geheim"
		surl := signUrl(url, pass)
		So(true, ShouldNotEqual, validate("_"+surl, pass))
	})
}
