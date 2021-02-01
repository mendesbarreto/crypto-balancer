package signature

import (
	"testing"
)

func TestGenerate(test *testing.T) {
	generateWithEmptyHash := Generate("")
	blockSize := generateWithEmptyHash.BlockSize()
	currentSize := generateWithEmptyHash.Size()

	if blockSize != 64 && currentSize != 32 {
		test.Error("The size of the string is not as expected")
	}
}
