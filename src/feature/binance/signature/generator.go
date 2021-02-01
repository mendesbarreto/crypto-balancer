package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"hash"
)

func Generate(apiSecreteKey string) hash.Hash {
	return hmac.New(sha256.New, []byte(apiSecreteKey))
}
