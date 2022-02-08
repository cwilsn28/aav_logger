package models

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}
