package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

var store *sessions.CookieStore

func init() {
    // TBD: grab key from environ
    store = sessions.NewCookieStore([]byte("secret-key"))
}

type RequestContext struct {
	Response http.ResponseWriter
	Request  *http.Request
	Vars     map[string]string
}

func (ctx *RequestContext) GetSession(name string) (*sessions.Session, error) {
	return store.Get(ctx.Request, name)
}

func (ctx *RequestContext) SaveSession(session *sessions.Session) error {
	return session.Save(ctx.Request, ctx.Response)
}

func (ctx *RequestContext) Var(name string) string {
	return ctx.Vars[name]
}

func (ctx *RequestContext) DecodeJSON(value interface{}) error {
	return json.NewDecoder(ctx.Request.Body).Decode(value)
}

func (ctx *RequestContext) RenderJSON(status int, value interface{}) {
	ctx.Response.WriteHeader(status)
	ctx.Response.Header().Add("content-type", "application/json")
	json.NewEncoder(ctx.Response).Encode(value)
}

func (ctx *RequestContext) RenderString(status int, msg string) {
	ctx.Response.WriteHeader(status)
	ctx.Response.Write([]byte(msg))
}

func (ctx *RequestContext) HandleOK(msg string) {
	ctx.RenderString(http.StatusOK, msg)
}

func (ctx *RequestContext) HandleNotFound(msg string) {
	ctx.RenderString(http.StatusNotFound, msg)
}

func (ctx *RequestContext) HandleError(err error) {
	http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
}

func NewRequestContext(w http.ResponseWriter, r *http.Request) *RequestContext {
	return &RequestContext{Response: w, Request: r, Vars: mux.Vars(r)}
}

type AppHandler func(ctx *RequestContext)

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn(NewRequestContext(w, r))
}
