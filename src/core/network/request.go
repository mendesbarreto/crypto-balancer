package network

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type HttpMethod string

type Request struct {
	Method      string
	QueryValues url.Values
	BodyValues  url.Values
	Header      http.Header
	Body        io.Reader
	Path        string
	Url         string
	BaseURL     string
}

func (request *Request) AddParam(key string, value interface{}) *Request {
	request.QueryValues.Add(key, fmt.Sprintf("%v", value))
	return request
}

func (request *Request) SetParam(key string, value interface{}) *Request {
	request.QueryValues.Set(key, fmt.Sprintf("%v", value))
	return request
}

func (request *Request) AddFormParam(key string, value interface{}) *Request {
	request.BodyValues.Add(key, fmt.Sprintf("%v", value))
	return request
}

func (request *Request) SetFormParam(key string, value interface{}) *Request {
	request.BodyValues.Set(key, fmt.Sprintf("%v", value))
	return request
}
