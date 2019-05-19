package wallet

import (
	"log"

	"github.com/mr-tron/base58"
)

// base58 is basically the same as base64, except it doesn't use 6 characters, namely:
// 0 O l I + /
// They took out the first 4 because 0 and O are similar, so are 1 and I
func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}

	return decode
}
