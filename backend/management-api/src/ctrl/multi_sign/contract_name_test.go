package multi_sign

import (
	"encoding/hex"
	"testing"

	"gotest.tools/assert"

	"chainmaker.org/chainmaker/common/v3/evmutils"
)

func Test_contractName(t *testing.T) {
	hexStr := "f7111142e72725dbccdf1d201eeb5cecc82530c6"
	contractName := "token002"
	assert.Equal(t, hexStr, hex.EncodeToString(evmutils.Keccak256([]byte(contractName)))[24:])
}
