package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var mh myHandler
	h := NoSurf(&mh)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Error(fmt.Sprintf("type is not http.Handler %T", v))
	}
}

func TestLoadSession(t *testing.T) {
	var mh myHandler
	h := LoadSession(&mh)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Error(fmt.Sprintf("type is not http.Handler %T", v))
	}
}
