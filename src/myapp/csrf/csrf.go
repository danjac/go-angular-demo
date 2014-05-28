package csrf

import (
	"code.google.com/p/xsrftoken"
	"fmt"
	"myapp/render"
	"net/http"
)

const (
	XsrfCookieName = "csrf_token"
	XsrfHeaderName = "X-CSRF-Token"
)

type CSRF struct {
	SecretKey string
	Handler   http.Handler
}

func NewCSRF(secretKey string, handler http.Handler) *CSRF {
	csrf := &CSRF{secretKey, handler}
	return csrf
}

func (csrf *CSRF) Validate(w http.ResponseWriter, r *http.Request) bool {

	var token string

	cookie, err := r.Cookie(XsrfCookieName)
	if err != nil || cookie.Value == "" {
		token = csrf.Generate()
		csrf.Save(w, token)
	} else {
		token = cookie.Value
	}

	if r.Method == "GET" || r.Method == "OPTIONS" || r.Method == "HEAD" {
		return true
	}

	headerValue := r.Header.Get(XsrfHeaderName)
	return headerValue == token
}

func (csrf *CSRF) Generate() string {
	return xsrftoken.Generate(csrf.SecretKey, "xsrf", "POST")
}

func (csrf *CSRF) Save(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:     XsrfCookieName,
		Path:     "/",
		MaxAge:   0,
		HttpOnly: false,
		Raw:      fmt.Sprintf("%s=%s", XsrfCookieName, token),
		Unparsed: []string{fmt.Sprintf("token=%s", token)},
	}
	http.SetCookie(w, cookie)
}

func (csrf *CSRF) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !csrf.Validate(w, r) {
		render.Status(w, http.StatusForbidden)
		return
	}
	csrf.Handler.ServeHTTP(w, r)
}
