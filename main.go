package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/types"

	"github.com/borispovod/guid/lib"

	"guid/lib"
)

const (
	prefix = "wallet"
)

// Example how to get GUID for specific address and EventHandler id.
// Launch: go run main.go wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8
func main() {
	address := os.Args[1]
	bzAddr, err := types.GetFromBech32(strings.ToLower(address), prefix)
	if err != nil {
		panic(fmt.Errorf("error during address conversation: %v", err))
	}

	withdrawGuid := lib.GetGUID(bzAddr, prefix, 0)
	depositGuid  := lib.GetGUID(bzAddr, prefix, 1)

	fmt.Printf("Withdraw GUID: 0x%s\n", hex.EncodeToString(withdrawGuid))
	fmt.Printf("Deposit GUID: 0x%s\n", hex.EncodeToString(depositGuid))
}
