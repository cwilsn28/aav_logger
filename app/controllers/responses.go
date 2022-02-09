package controllers

import (
	"net/http"

	"github.com/revel/revel"
)

type Accepted string
type BadRequest string
type Created string
type MethodNotAllowed string
type OK string
type ServerError string

func (r OK) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "application/json")
	resp.GetWriter().Write([]byte(r))
}

func (r Accepted) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusAccepted, "application/json")
	resp.GetWriter().Write([]byte(r))
}

func (r BadRequest) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusBadRequest, "application/json")
	resp.GetWriter().Write([]byte(r))
}

func (r Created) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusCreated, "application/json")
	resp.GetWriter().Write([]byte(r))
}

func (r MethodNotAllowed) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusMethodNotAllowed, "application/json")
	resp.GetWriter().Write([]byte(r))
}

func (r ServerError) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusInternalServerError, "application/json")
	resp.GetWriter().Write([]byte(r))
}
