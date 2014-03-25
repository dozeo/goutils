package goutils

import (
	. "github.com/smartystreets/goconvey/convey"
	"os/exec"
	"strconv"
	"testing"
	"time"
)

func Test_HMAC_ruby(t *testing.T) {
	MyConvey("URL 0 OK", t, func() {
		url := "http://www.eklenet.de"
		pass := "geheim"
		cur := time.Now().Unix()
		var cur32 = int32(cur)
		So(int64(cur32), ShouldEqual, cur)
		ttl := 30000
		endoflife := int(cur) + ttl
		out, err := exec.Command("ruby", "hmac_test/ref.rb", url, pass, strconv.Itoa(endoflife)).Output()
		So(nil, ShouldEqual, err)
		So(HMAC.SignUrl(url, pass, endoflife), ShouldEqual, string(out))
		So(HMAC.Validate(string(out), pass), ShouldEqual, true)
	})
	MyConvey("URL 1 OUTDATED", t, func() {
		url := "http://www.eklenet.de"
		pass := "geheim"
		cur := time.Now().Unix()
		var cur32 = int32(cur)
		So(int64(cur32), ShouldEqual, cur)
		ttl := -30000
		endoflife := int(cur) + ttl
		out, err := exec.Command("ruby", "hmac_test/ref.rb", url, pass, strconv.Itoa(endoflife)).Output()
		So(nil, ShouldEqual, err)
		So(HMAC.SignUrl(url, pass, endoflife), ShouldEqual, string(out))
		So(HMAC.Validate(string(out), pass), ShouldEqual, false)
	})
	MyConvey("URL 2 OK /", t, func() {
		url := "http://www.eklenet.de/"
		pass := "geheim"
		ttl := int(time.Now().Unix()) + 30
		out, err := exec.Command("ruby", "hmac_test/ref.rb", url, pass, strconv.Itoa(ttl)).Output()
		So(nil, ShouldEqual, err)
		So(HMAC.SignUrl(url, pass, ttl), ShouldEqual, string(out))
		So(true, ShouldEqual, HMAC.Validate(string(out), pass))
	})
	MyConvey("URL 3 OK /abc", t, func() {
		url := "http://www.eklenet.de/abc"
		pass := "geheim"
		ttl := int(time.Now().Unix()) + 30
		out, err := exec.Command("ruby", "hmac_test/ref.rb", url, pass, strconv.Itoa(ttl)).Output()
		So(nil, ShouldEqual, err)
		So(HMAC.SignUrl(url, pass, ttl), ShouldEqual, string(out))
		So(true, ShouldEqual, HMAC.Validate(string(out), pass))
	})
	MyConvey("URL 4 OK /abc/", t, func() {
		url := "http://www.eklenet.de/abc/"
		pass := "geheim"
		ttl := int(time.Now().Unix()) + 30
		out, err := exec.Command("ruby", "hmac_test/ref.rb", url, pass, strconv.Itoa(ttl)).Output()
		So(nil, ShouldEqual, err)
		So(HMAC.SignUrl(url, pass, ttl), ShouldEqual, string(out))
		So(true, ShouldEqual, HMAC.Validate(string(out), pass))
	})
	MyConvey("URL 5 OK /abc/def?xyz=123&ok=no", t, func() {
		url := "http://www.eklenet.de/abc/def?xyz=123&ok=no"
		pass := "geheim"
		ttl := int(time.Now().Unix()) + 30
		out, err := exec.Command("ruby", "hmac_test/ref.rb", url, pass, strconv.Itoa(ttl)).Output()
		So(nil, ShouldEqual, err)
		So(HMAC.SignUrl(url, pass, ttl), ShouldEqual, string(out))
		So(HMAC.Validate(string(out), pass), ShouldEqual, true)
	})
	MyConvey("URL 6 OK /abc/def?ok=no&xyz=123 (switched)", t, func() {
		url := "http://www.eklenet.de/abc/def?ok=no&xyz=123"
		pass := "geheim"
		ttl := int(time.Now().Unix()) + 30
		out, err := exec.Command("ruby", "hmac_test/ref.rb", url, pass, strconv.Itoa(ttl)).Output()
		So(nil, ShouldEqual, err)
		So(HMAC.SignUrl(url, pass, ttl), ShouldEqual, string(out))
		So(true, ShouldEqual, HMAC.Validate(string(out), pass))
	})
	MyConvey("URL 7 OK /abc/def?ok=no&xyz=123#what", t, func() {
		url := "http://www.eklenet.de/abc/def?ok=no&xyz=123#what"
		pass := "geheim"
		ttl := int(time.Now().Unix()) + 30
		out, err := exec.Command("ruby", "hmac_test/ref.rb", url, pass, strconv.Itoa(ttl)).Output()
		So(nil, ShouldEqual, err)
		So(HMAC.SignUrl(url, pass, ttl), ShouldEqual, string(out))
		So(true, ShouldEqual, HMAC.Validate(string(out), pass))
	})
	MyConvey("URL 8 OK /abc/def?ok=no&xyz=123#what", t, func() {
		url := "http://www.eklenet.de/abc/def?ok=no&xyz=123#what"
		pass := "geheim"
		surl := HMAC.SignUrl(url, pass, 30)
		So(true, ShouldEqual, HMAC.Validate(surl, pass))
	})
	MyConvey("URL 9 OK /abc/def?ok=no&xyz=123#what (switched)", t, func() {
		url := "http://www.eklenet.de/abc/def?ok=no&xyz=123#what"
		pass := "geheim"
		surl := HMAC.SignUrl(url, pass, 30)
		So(true, ShouldNotEqual, HMAC.Validate("_"+surl, pass))
	})
}
