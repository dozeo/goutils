package goutils

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/zenazn/goji/web"
)

var TestResponseWriterContent []byte
var TestResponseWriterCode int

type TestResponseWriter struct {
}
type testin struct {
	Name string
}
type testout struct {
	Name string
}

func (trw TestResponseWriter) Header() http.Header {
	return nil
}

func (trw TestResponseWriter) Write(b []byte) (int, error) {
	TestResponseWriterContent = append(TestResponseWriterContent, b...)
	return len(b), nil
}

func (trw TestResponseWriter) WriteHeader(c int) {
	TestResponseWriterCode = c
	return
}

func Test_New(t *testing.T) {
	MyConvey("working", t, func() {
		tf := func(d testin, c web.C) (testout, error) { return testout{Name: "out"}, nil }
		wit := TestResponseWriter{}
		req, _ := http.NewRequest("GET", "", bytes.NewBufferString("{\"Name\":\"hi\"}"))
		GojiRestJsonWrapper(tf)(web.C{}, wit, req)
		So(TestResponseWriterCode, ShouldEqual, 200)
	})
	MyConvey("working error case", t, func() {
		tf := func(d testin, c web.C) (testout, error) {
			return testout{Name: "out"}, errors.New("this error should happen for testing")
		}
		wit := TestResponseWriter{}
		req, _ := http.NewRequest("GET", "", bytes.NewBufferString("{\"Name\":\"hi\"}"))
		GojiRestJsonWrapper(tf)(web.C{}, wit, req)
		So(TestResponseWriterCode, ShouldEqual, 500)
	})
	MyConvey("wrong json", t, func() {
		tf := func(d testin, c web.C) (testout, error) { return testout{Name: "out"}, nil }
		wit := TestResponseWriter{}
		req, _ := http.NewRequest("GET", "", bytes.NewBufferString("Name:\"hi\"}"))
		GojiRestJsonWrapper(tf)(web.C{}, wit, req)
		So(TestResponseWriterCode, ShouldEqual, 500)
	})
	MyConvey("empty json", t, func() {
		tf := func(d testin, c web.C) (testout, error) { return testout{Name: "out"}, nil }
		wit := TestResponseWriter{}
		req, _ := http.NewRequest("GET", "", bytes.NewBufferString(""))
		GojiRestJsonWrapper(tf)(web.C{}, wit, req)
		So(TestResponseWriterCode, ShouldEqual, 200)
	})
	MyConvey("wrong parameter count 3", t, func() {
		tf := func(d testin, c web.C, s string) (testout, error) { return testout{Name: "out"}, nil }
		wit := TestResponseWriter{}
		req, _ := http.NewRequest("GET", "", bytes.NewBufferString("{\"Name\":\"hi\"}"))
		GojiRestJsonWrapper(tf)(web.C{}, wit, req)
		So(TestResponseWriterCode, ShouldEqual, 500)
	})
	MyConvey("wrong parameter count 1", t, func() {
		tf := func(d testin) (testout, error) { return testout{Name: "out"}, nil }
		wit := TestResponseWriter{}
		req, _ := http.NewRequest("GET", "", bytes.NewBufferString("{\"Name\":\"hi\"}"))
		GojiRestJsonWrapper(tf)(web.C{}, wit, req)
		So(TestResponseWriterCode, ShouldEqual, 500)
	})
	MyConvey("wrong parameter type in1", t, func() {
		tf := func(d string, c web.C) (testout, error) { return testout{Name: "out"}, nil }
		wit := TestResponseWriter{}
		req, _ := http.NewRequest("GET", "", bytes.NewBufferString("{\"Name\":\"hi\"}"))
		GojiRestJsonWrapper(tf)(web.C{}, wit, req)
		So(TestResponseWriterCode, ShouldEqual, 500)
	})
	MyConvey("wrong parameter type in2", t, func() {
		tf := func(d testin, c string) (testout, error) { return testout{Name: "out"}, nil }
		wit := TestResponseWriter{}
		req, _ := http.NewRequest("GET", "", bytes.NewBufferString("{\"Name\":\"hi\"}"))
		GojiRestJsonWrapper(tf)(web.C{}, wit, req)
		So(TestResponseWriterCode, ShouldEqual, 500)
	})
	MyConvey("wrong parameter type out1", t, func() {
		tf := func(d testin, c web.C) (string, error) { return "out", nil }
		wit := TestResponseWriter{}
		req, _ := http.NewRequest("GET", "", bytes.NewBufferString("{\"Name\":\"hi\"}"))
		GojiRestJsonWrapper(tf)(web.C{}, wit, req)
		So(TestResponseWriterCode, ShouldEqual, 500)
	})
	MyConvey("wrong parameter type out2", t, func() {
		tf := func(d testin, c web.C) (testout, string) { return testout{Name: "out"}, "" }
		wit := TestResponseWriter{}
		req, _ := http.NewRequest("GET", "", bytes.NewBufferString("{\"Name\":\"hi\"}"))
		GojiRestJsonWrapper(tf)(web.C{}, wit, req)
		So(TestResponseWriterCode, ShouldEqual, 500)
	})
}
