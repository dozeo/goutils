package goutils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/zenazn/goji/web"
)

func WriterError(w http.ResponseWriter, e error) {
	w.WriteHeader(500)
	w.Write([]byte(e.Error()))
}

var ErrorRestFunc = errors.New("{\"error\":\"Rest function must be func(struct, web.C) (struct, error)\"}")
var ErrorRequestBody = errors.New("{\"error\":\"Could not read request body\"}")
var ErrorRequestJsonUnmarshal = errors.New("{\"error\":\"Could not generate struct vom json\"}")

func GojiRestJsonWrapper(f interface{}) func(c web.C, w http.ResponseWriter, r *http.Request) {
	return func(c web.C, w http.ResponseWriter, r *http.Request) {
		fType := reflect.TypeOf(f)
		if fType.NumIn() != 2 || fType.NumOut() != 2 {
			WriterError(w, ErrorRestFunc)
			return
		}
		if fType.In(1).Name() != "C" {
			WriterError(w, ErrorRestFunc)
			return
		}
		if fType.Out(1).Name() != "error" {
			WriterError(w, ErrorRestFunc)
			return
		}
		fIn0 := fType.In(0)
		fOut0 := fType.Out(0)
		if fIn0.Kind() != reflect.Struct || fOut0.Kind() != reflect.Struct {
			WriterError(w, ErrorRestFunc)
			return
		}
		InStruct := reflect.New(fIn0)
		content, bodyErr := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if bodyErr != nil {
			WriterError(w, ErrorRequestBody)
			return
		}
		if len(content) > 0 {
			unmarshalErr := json.Unmarshal(content, InStruct.Interface())
			if unmarshalErr != nil {
				WriterError(w, ErrorRequestJsonUnmarshal)
				return
			}
		}
		outValues := reflect.ValueOf(f).Call([]reflect.Value{InStruct.Elem(), reflect.ValueOf(c)})
		if outValues[1].Interface() != nil {
			var e error
			e = outValues[1].Interface().(error)
			WriterError(w, e)
			return
		}
		outJson, marshalErr := json.Marshal(outValues[0].Interface())
		if marshalErr == nil {
			w.WriteHeader(200)
			w.Write(outJson)
		} else {
			WriterError(w, marshalErr)
		}
	}
}

/*
import "github.com/zenazn/goji"
type data struct {
	Name string
}
type ret struct {
	Name string
}

func myrestfunc(d data, c web.C) (ret, error) {
	d.Name += " add"
	return ret{Name: d.Name}, nil
}

func main() {
	goji.Post("/hello3/:name", GojiRestJsonWrapper(myrestfunc))
	goji.Serve()
}
*/
