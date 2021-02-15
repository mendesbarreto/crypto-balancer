package endpoints

type AccountEndpointType string

const (
	Version AccountEndpointType = "v3"
	Account                     = "asset"
)

type AccountEndpointParamsType string

const (
	Timestamp AccountEndpointParamsType = "timestamp"
	Signature                           = "signature"
)
