package easy_http

import (
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	var res = []byte{}
	req := HttpReq{
		Method:  "GET",
		Url:     "https://example.com",
		ResData: &res,
		UnMarshalFunc: func(buf []byte, dest interface{}) error {
			reflect.ValueOf(dest).Elem().SetBytes(buf)
			return nil
		},
	}
	err := req.Do()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
