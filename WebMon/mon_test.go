package main

import (
	"testing"
)

const (
)

func Test_nginx(t *testing.T) {
	if text, code, err := httpGet(url); err != nil {
		t.Error(err)
	} else {
		t.Log("text :", text)
		t.Log("code:", code)
		status, _ := nginx_status(text)
		t.Log("status:", status)
		data := nginx_data(status)
		t.Log("data:", data)
	}
}
