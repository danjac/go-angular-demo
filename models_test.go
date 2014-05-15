package main

import (
	"github.com/martini-contrib/binding"
	"net/http"
	"testing"
)

func TestValidatePostIfContentTooLong(t *testing.T) {
	var errors binding.Errors
	errors.Fields = make(map[string]string)

	s := "The number of map elements is called its length. For a map m, it can be discovered using the built-in function len and may change during execution."

	p := Post{Content: s}
	p.Validate(&errors, &http.Request{})

	msg, _ := errors.Fields["content"]
	if msg != "Content must be max 140 characters" {
		t.Error("Should validate content < 140 chars")
	}

}

func TestValidatePostIfContentEmpty(t *testing.T) {

	var errors binding.Errors
	errors.Fields = make(map[string]string)

	p := Post{Content: ""}
	p.Validate(&errors, &http.Request{})

	msg, _ := errors.Fields["content"]
	if msg != "Content is missing" {
		t.Error("Should validate missing content")
	}
}
