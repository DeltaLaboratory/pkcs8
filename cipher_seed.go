package pkcs8

import (
	"encoding/asn1"

	"github.com/RyuaNerin/go-krypto/seed"
)

var (
	oidSEED128CBC = asn1.ObjectIdentifier{1, 2, 410, 200004, 1, 4}
)

func init() {
	RegisterCipher(oidSEED128CBC, func() Cipher {
		return SEED128CBC
	})
}

// SEED128CBC is the 128-bit key SEED cipher in CBC mode.
var SEED128CBC = cipherWithBlock{
	ivSize:   16,
	keySize:  16,
	newBlock: seed.NewCipher,
	oid:      oidSEED128CBC,
}
