package main

import (
	"code.google.com/p/xsrftoken"
	"fmt"
	"net/http"
)

const (
	XsrfCookieName = "csrf_token"
	XsrfHeaderName = "X-CSRF-Token"
)

type CSRF struct {
	Request  *http.Request
	Response http.ResponseWriter
}

func NewCSRF(w http.ResponseWriter, r *http.Request) *CSRF {
	csrf := &CSRF{r, w}
	return csrf
}

func (csrf *CSRF) Validate() bool {

	var token string

	cookie, err := csrf.Request.Cookie(XsrfCookieName)
	if err != nil || cookie.Value == "" {
		token = csrf.Generate()
		csrf.Save(token)
	} else {
		token = cookie.Value
	}

	if csrf.Request.Method == "GET" {
		return true
	}

	headerValue := csrf.Request.Header.Get(XsrfHeaderName)
	return headerValue == token
}

func (csrf *CSRF) Generate() string {
	// TBD: set up secret key from env
	return xsrftoken.Generate("secret-key", "xsrf", "POST")
}

func (csrf *CSRF) Save(token string) {
	cookie := &http.Cookie{
		Name:     XsrfCookieName,
		Path:     "/",
		MaxAge:   0,
		HttpOnly: false,
		Raw:      fmt.Sprintf("%s=%s", XsrfCookieName, token),
		Unparsed: []string{fmt.Sprintf("token=%s", token)},
	}
	http.SetCookie(csrf.Response, cookie)
}
