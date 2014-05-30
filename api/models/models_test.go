package models

import (
	"testing"
)

func TestValidatePostIfContentTooLong(t *testing.T) {

	s := "The number of map elements is called its length. For a map m, it can be discovered using the built-in function len and may change during execution."

	p := Post{Content: s}
	result := p.Validate()

	if result.OK {
		t.Error("Should be invalid")
	}

	msg, _ := result.Errors["content"]
	if msg != "Content must be max 140 characters" {
		t.Error("Should validate content < 140 chars")
	}

}

func TestValidatePostIfContentEmpty(t *testing.T) {
	p := Post{Content: ""}
	result := p.Validate()

	if result.OK {
		t.Error("Should be invalid")
	}

	msg, _ := result.Errors["content"]
	if msg != "Content is missing" {
		t.Error("Should validate missing content")
	}
}
