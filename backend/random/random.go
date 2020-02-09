package random

import (
	"encoding/base32"
	"fmt"
	"strings"

	"github.com/jtbonhomme/gotp"
)

// keyringSize is the default size of the keyring
const keyringSize int = 5

// Random is a backend used for tests. It generates 5 random keys at startup.
type Random struct {
	serviceName string
	keys        []gotp.Key
}

// New creates a new backend instance
func New(name string) *Random {
	rd := Random{
		serviceName: name,
	}

	for i := 0; i < keyringSize; i++ {
		secret := base32.StdEncoding.EncodeToString([]byte(strings.ToUpper(fmt.Sprintf("value%d", i))))
		key := gotp.Key{
			URI:   fmt.Sprintf("uri%d", i),
			Value: secret,
		}
		rd.keys = append(rd.keys, key)
	}
	return &rd
}

// List lists all keys stored in the backend
func (rd *Random) List() (*[]gotp.TOTP, error) {
	var totps []gotp.TOTP
	for _, key := range rd.keys {
		code, err := gotp.TOTPToken(key.Value)
		if err != nil {
			return &totps, err
		}
		totp := gotp.TOTP{
			URI:  key.URI,
			Code: code,
		}
		totps = append(totps, totp)
	}
	return &totps, nil
}

// Store adds a new key in the backend
func (rd *Random) Store(uri, value string) error {
	var err error
	key := gotp.Key{
		URI:   uri,
		Value: value,
	}
	rd.keys = append(rd.keys, key)
	return err
}
