package network

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type HttpMethod string
type QueryStringMapper func(value string) string
type Params map[string]interface{}

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

func (request *Request) AddParam(key string, value string) *Request {
	request.QueryValues.Add(key, value)
	return request
}

func (request *Request) SetParam(key string, value string) *Request {
	request.QueryValues.Set(key, value)
	return request
}

func (request *Request) SetParams(params Params) *Request {
	for key, value := range params {
		request.SetParam(key, fmt.Sprintf("%v", value))
	}
	return request
}

func (request *Request) QueryString() string {
	queryString := request.QueryValues.Encode()

	if request.QueryStringMapper != nil {
		queryString = request.QueryStringMapper(queryString)
	}

	return queryString
}

func (request *Request) ToHttpRequest(ctx context.Context) (httpRequest *http.Request, err error) {
	httpUrl := request.Url
	httpRequest, err = http.NewRequest(request.Method, httpUrl, nil)

	if err != nil {
		return nil, err
	}

	httpRequest = httpRequest.WithContext(ctx)
	httpRequest.Header = request.Header
	httpRequest.URL.RawQuery = request.QueryString()

	return httpRequest, err
}
