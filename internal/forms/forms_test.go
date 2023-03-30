package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/testFormValid", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("Form_Valid the form should not have an error")
	}

	form.Errors.Add("error", "add an error to make test")
	isValid = form.Valid()
	if isValid {
		t.Error("Form_Valid the form should have at list an error")
	}
}

func TestForm_Required(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Form_Required the form should not have an error")
	}

	postedData = url.Values{}
	postedData.Add("a", "value")
	postedData.Add("b", "value")
	postedData.Add("c", "value")

	form = New(postedData)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Form_Required does not have required fields when it does ")
	}
}

func TestForm_Has(t *testing.T) {

	postedData := url.Values{}
	form := New(postedData)
	isHas := form.Has("x")
	if isHas {
		t.Error("Form_Has the field empty should not have data ")
	}

	postedData = url.Values{}
	postedData.Add("withValue", "value")
	form = New(postedData)
	isHas = form.Has("withValue")
	if !isHas {
		t.Error("Form_Has the field withValue should has some value")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("correct", "test@testOK.com")
	form := New(postedData)
	form.IsEmail("correct")
	if !form.Valid() {
		t.Error("Form_IsEmail correct email, but some error happened")
	}

	postedData = url.Values{}
	postedData.Add("wrong", "isNotAnEmail")
	form = New(postedData)
	form.IsEmail("wrong")
	if form.Valid() {
		t.Error("Form_IsEmail wrong email, but has validated as correct")
	}

	postedData = url.Values{}
	form = New(postedData)
	form.IsEmail("x")
	if form.Valid() {
		t.Error("Form_IsEmail invalid field, but has validated as correct")
	}

}

func TestForm_MinLenght(t *testing.T) {

	postedData := url.Values{}
	form := New(postedData)
	form.MinLenght("wrong", 4)
	if form.Valid() {
		t.Error("Form_MinLenght wrong size, but has validated as correct")
	}

	postedData = url.Values{}
	form = New(postedData)
	form.MinLenght("x", 100)
	if form.Valid() {
		t.Error("Form_MinLenght invalid field, but has validated as correct")
	}

	postedData = url.Values{}
	postedData.Add("correct", "four")
	form = New(postedData)
	form.MinLenght("correct", 4)
	if !form.Valid() {
		t.Error("Form_MinLenght correct size, but some error happened")
	}

}

func TestErrors_Get(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	form.MinLenght("x", 100)

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("Form_get should have an error, but did not get one")
	}

	postedData = url.Values{}
	postedData.Add("correct", "four")
	form = New(postedData)
	form.MinLenght("correct", 4)

	isError = form.Errors.Get("correct")
	if isError != "" {
		t.Error("Form_get should not have error, but got one")
	}
}

//go test -coverprofile=covered1
//go tool cover -html=covered1
