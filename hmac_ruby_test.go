package goutils

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"os/exec"
	"strconv"
	"testing"
	"time"
	//"strings"
)

var test_tmp string

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
	MyConvey("URL 0a OK", t, func() {
		url := "http://dev.doz.io:4080/info.json?url=https%3A%2F%2Fdemo.dozeoapp.com.dev.doz.io%2Fmeetings%2F50340910-96f1-0131-8282-7214ed8242fc%2Fattachments%2Fshow%3Fauth%5Bdate%5D%3DWed%252C%252026%2520Mar%25202014%252009%253A08%253A43%2520GMT%26auth%5Bsignature%5D%3D5cfc250f6ddfda5ed0e37d036ec8a6b8a8744a0c%26file%3De8d32e881c8c01bf036353154d3d2517b150cbee"
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
	MyConvey("URL Double 1", t, func() {
		url := "https://demo.dozeoapp.com.dev.doz.io/meetings/50340910-96f1-0131-8282-7214ed8242fc/attachments/show?file=e8d32e881c8c01bf036353154d3d2517b150cbee"
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
		test_tmp = string(out)
	})
	MyConvey("URL Double 2", t, func() {
		url := "http://dev.doz.io:4080/info.json?url=" + url.QueryEscape(test_tmp)
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
	MyConvey("URL Double 3", t, func() {
		url := "http://dev.doz.io:4080/info.json?url=" + url.QueryEscape(test_tmp) + "&page=1&size=50"
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
