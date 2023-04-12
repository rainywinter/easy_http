package easy_http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HttpReq struct {
	Url           string
	Values        map[string]string
	Header        map[string]string
	ReqData       interface{}
	Res           interface{}
	MarshalFunc   func(v interface{}) ([]byte, error)      // default json.Marshal
	UnMarshalFunc func(buf []byte, dest interface{}) error // default json.Marshal
}

func (h HttpReq) Get() error {
	if h.MarshalFunc == nil {
		h.MarshalFunc = json.Marshal
	}
	if h.UnMarshalFunc == nil {
		h.UnMarshalFunc = json.Unmarshal
	}

	vs := url.Values{}
	for k, v := range h.Values {
		vs.Set(k, v)
	}

	body, err := h.MarshalFunc(h.ReqData)
	if err != nil {
		return err
	}

	url := h.Url + "?" + vs.Encode()
	r, err := http.NewRequest("GET", url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	if len(h.Header) != 0 {
		for k, v := range h.Header {
			r.Header.Set(k, v)
		}
	}
	return h.clientDo(r)
}

func (h HttpReq) clientDo(r *http.Request) error {
	c := http.Client{}
	res, err := c.Do(r)
	if err != nil {
		return err
	}
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = h.UnMarshalFunc(buf, h.Res)
	return err
}
