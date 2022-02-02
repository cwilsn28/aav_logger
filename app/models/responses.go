package models

type APIResponse struct {
	Status  string      `json:"status"`
	Results interface{} `json:"results"`
}
