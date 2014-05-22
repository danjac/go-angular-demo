package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type RequestContext struct {
	Response http.ResponseWriter
	Request  *http.Request
	Vars     map[string]string
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

func (ctx *RequestContext) HandleOK(msg string) {
	ctx.Response.WriteHeader(http.StatusOK)
	ctx.Response.Write([]byte(msg))
}

func (ctx *RequestContext) HandleNotFound(msg string) {
	ctx.Response.WriteHeader(http.StatusNotFound)
	ctx.Response.Write([]byte(msg))
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
