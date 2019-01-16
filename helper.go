package jssdk

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpGet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return content, nil
}

func RandomString() string {
	h := md5.New()
	io.WriteString(h, time.Now().Format(time.RFC3339Nano))
	return fmt.Sprintf("%x", h.Sum(nil))
}
