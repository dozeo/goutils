package goutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

func Test_goji_rest(t *testing.T) {
	MyConvey("Bench rest wrapper 1 ", t, func() {
		tf := func(d testin, c web.C) (testout, error) { return testout{Name: "out"}, nil }
		f := GojiRestJsonWrapper(tf)
		b1 := Bench{}
		for i := 0; i < 10000; i++ {
			wit := TestResponseWriter{}
			req, _ := http.NewRequest("GET", "", bytes.NewBufferString("{\"Name\":\"hi\"}"))
			b1.Start()
			f(web.C{}, wit, req)
			b1.Stop()
		}
		fmt.Println(b1.Sum() / 1000 / 1000)
		So(b1.Sum(), ShouldBeGreaterThan, 0)
	})
	MyConvey("Bench rest wrapper 2", t, func() {
		tf := func(d testin, c web.C) (testout, error) { return testout{Name: "out"}, nil }
		f := GojiRestJsonWrapper2(tf)
		b1 := Bench{}
		for i := 0; i < 10000; i++ {
			wit := TestResponseWriter{}
			req, _ := http.NewRequest("GET", "", bytes.NewBufferString("{\"Name\":\"hi\"}"))
			b1.Start()
			f(web.C{}, wit, req)
			b1.Stop()
		}
		fmt.Println(b1.Sum() / 1000 / 1000)
		So(b1.Sum(), ShouldBeGreaterThan, 0)
	})
	MyConvey("Bench rest native", t, func() {
		tf := func(d testin, c web.C) (testout, error) { return testout{Name: "out"}, nil }
		f := func(c web.C, w http.ResponseWriter, r *http.Request) {
			body, _ := ioutil.ReadAll(r.Body)
			j := testin{}
			json.Unmarshal(body, &j)
			o, _ := tf(j, c)
			od, _ := json.Marshal(o)
			w.Write(od)
		}
		b := Bench{}
		for i := 0; i < 10000; i++ {
			wit := TestResponseWriter{}
			req, _ := http.NewRequest("GET", "", bytes.NewBufferString("{\"Name\":\"hi\"}"))
			b.Start()
			f(web.C{}, wit, req)
			b.Stop()
		}
		fmt.Println(b.Sum() / 1000 / 1000)
		So(b.Sum(), ShouldBeGreaterThan, 0)
	})
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
