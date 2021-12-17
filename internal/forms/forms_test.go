package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("should be valid since no errors")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when reqired fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	form = New(postedData)

	has = form.Has("a")

	if !has {
		t.Error("shows form does not have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("whatever", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	postedData := url.Values{}
	postedData.Add("some_field", "some_value")
	form = New(postedData)

	form.MinLength("some_field", 100)

	if form.Valid() {
		t.Error("form should be invalid for some_field, less than 100 length")
	}

	postedData = url.Values{}
	postedData.Add("another_field", "abc123")
	form = New(postedData)

	form.MinLength("another_field", 1)

	if !form.Valid() {
		t.Error("should pass the test, the vlaue of another field is greater that 1")
	}
}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.IsEmail("x")

	if form.Valid() {
		t.Error("form shows x is email for non-existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should get err msg on field x")
	}

	postedData := url.Values{}
	postedData.Add("email", "a@h.com")
	form = New(postedData)
	form.IsEmail("email")

	if !form.Valid() {
		t.Error("shows invalid when given an valid email")
	}

	isError = form.Errors.Get("email")
	if isError != "" {
		t.Error("should not get err msg on field email but got one")
	}

	postedData = url.Values{}
	postedData.Add("email", "xxxxxx.com")
	form = New(postedData)
	form.IsEmail("email")

	if form.Valid() {
		t.Error("shows valid when given an invalid email")
	}

}
