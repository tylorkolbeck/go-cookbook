package model

type Endpoint struct {
	Name   string      `json:"name"`
	Method string      `json:"method"`
	Path   string      `json:"path"`
	Body   interface{} `json:"body"`
}
