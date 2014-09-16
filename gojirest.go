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
	fType := reflect.TypeOf(f)
	fValue := reflect.ValueOf(f)
	if fType.NumIn() != 2 || fType.NumOut() != 2 {
		return func(c web.C, w http.ResponseWriter, r *http.Request) { WriterError(w, ErrorRestFunc) }
	}
	if fType.In(1).Name() != "C" {
		return func(c web.C, w http.ResponseWriter, r *http.Request) { WriterError(w, ErrorRestFunc) }
	}
	if fType.Out(1).Name() != "error" {
		return func(c web.C, w http.ResponseWriter, r *http.Request) { WriterError(w, ErrorRestFunc) }
	}
	fIn0 := fType.In(0)
	fOut0 := fType.Out(0)
	if fIn0.Kind() != reflect.Struct || fOut0.Kind() != reflect.Struct {
		return func(c web.C, w http.ResponseWriter, r *http.Request) { WriterError(w, ErrorRestFunc) }
	}
	return func(c web.C, w http.ResponseWriter, r *http.Request) {
		content, bodyErr := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if bodyErr != nil {
			WriterError(w, ErrorRequestBody)
			return
		}
		InStruct := reflect.New(fIn0)
		if len(content) > 0 {
			unmarshalErr := json.Unmarshal(content, InStruct.Interface())
			if unmarshalErr != nil {
				WriterError(w, ErrorRequestJsonUnmarshal)
				return
			}
		}
		outValues := fValue.Call([]reflect.Value{InStruct.Elem(), reflect.ValueOf(c)})
		if outValues[1].Interface() != nil {
			e := outValues[1].Interface().(error)
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
