package pkcs8

import (
	"encoding/asn1"

	"github.com/RyuaNerin/go-krypto/seed"
)

var (
	oidSEEDCBC = asn1.ObjectIdentifier{1, 2, 410, 200004, 1, 4}
)

func init() {
	RegisterCipher(oidSEEDCBC, func() Cipher {
		return SEEDCBC
	})
}

// SEEDCBC is the 128-bit key SEED cipher in CBC mode.
var SEEDCBC = cipherWithBlock{
	ivSize:   16,
	keySize:  16,
	newBlock: seed.NewCipher,
	oid:      oidSEEDCBC,
}
