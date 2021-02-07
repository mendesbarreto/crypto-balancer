package network

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type HttpMethod string
type QueryStringMapper func(value string) string

type Request struct {
	Method            string
	QueryValues       url.Values
	BodyValues        url.Values
	Header            http.Header
	Body              io.Reader
	Path              string
	Url               string
	BaseURL           string
	QueryStringMapper QueryStringMapper
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

func (request *Request) QueryString() string {

	queryString := request.QueryValues.Encode()

	if request.QueryStringMapper != nil {
		queryString = request.QueryStringMapper(queryString)
	}

	return queryString
}

func (request *Request) FullUrl() string {
	queryString := request.QueryString()

	if queryString != "" {
		return fmt.Sprintf("%s?%s", request.Url, request.QueryString())
	}

	return request.Url
}
