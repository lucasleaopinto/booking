package forms

import (
	"net/http"
	"net/url"
)

//Form creates a custom form struct, embeds a url.values Object
type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct
func New(data url.Values) *Form{
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool{
	has := r.Form.Get(field)
	if has == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}

	return true
}

// Valid returns true if ther are no errors, otherwise false
func (f *Form) Valid() bool{
	return len(f.Errors) == 0
}