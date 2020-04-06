package lib

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	tmpGUID = "000000000000000077616c6c657400000000000095abf6bf9cd39a391567e4508becb25d0f1b98de"
)

func Test_GetGUID(t *testing.T) {
	address := "wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8"
	prefix := "wallet"
	bzAddr, err := types.GetFromBech32(strings.ToLower(address), prefix)
	if err != nil {
		t.Fatal(err)
	}

	guid := GetGUID(bzAddr, prefix, 0)
	require.Equal(t, tmpGUID, hex.EncodeToString(guid))
}