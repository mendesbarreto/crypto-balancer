package endpoints

type AccountEndpointType string

const (
	Version AccountEndpointType = "v3"
	account                     = "account"
)

type AccountEndpointParamsType string

const (
	timestamp AccountEndpointParamsType = "timestamp"
	signature                           = "signature"
)
