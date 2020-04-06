package lib

import (
	"encoding/binary"
)

const (
	libraAddressLength = 32
)

// Bytes to libra compability.
func addrToPathAddr(addr []byte, prefix string) []byte {
	zeros := make([]byte, libraAddressLength-len(prefix)-len(addr))

	bytes := make([]byte, 0)
	bytes = append(bytes, []byte(prefix)...)
	bytes = append(bytes, zeros...)
	bytes = append(bytes, addr...)

	return bytes
}

// Get event guid by address, bech32 prefix and counter.
func GetGUID(address []byte, prefix string, counter uint64) []byte {
	bzCounter := make([]byte, 8)

	addr := addrToPathAddr(address, prefix)

	binary.LittleEndian.PutUint64(bzCounter, counter)

	return append(bzCounter, addr...)
}