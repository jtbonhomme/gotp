package random

import (
	"encoding/base32"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jtbonhomme/gotp"
)

// keyringSize is the default size of the keyring
const keyringSize int = 5

// Random is a backend used for tests. It generates 5 random keys at startup.
type Random struct {
	keyring string
	keys    []gotp.Secret
}

// New creates a new backend instance.
// It takes as argument the name of the key ring to be created in storage.
func New(kr string) *Random {
	rd := Random{
		keyring: kr,
	}

	for i := 0; i < keyringSize; i++ {
		secret := base32.StdEncoding.EncodeToString([]byte(strings.ToUpper(fmt.Sprintf("value%d", i))))
		key := gotp.Secret{
			Key:   fmt.Sprintf("key%d", i),
			Value: secret,
		}
		rd.keys = append(rd.keys, key)
	}
	fmt.Printf("Populated keyring with %+v\n", rd)
	return &rd
}

// List lists all keys stored in the backend
func (rd *Random) List() (*[]gotp.TOTP, error) {
	var totps []gotp.TOTP
	interval := time.Now().Unix() / gotp.TimeIntervaleSeed
	fmt.Printf("Interval: %d\n", interval)

	for _, key := range rd.keys {
		code, err := gotp.TOTPToken([]byte(key.Value))
		if err != nil {
			return &totps, err
		}
		totp := gotp.TOTP{
			Key:  key.Key,
			Code: code,
		}
		totps = append(totps, totp)
	}
	return &totps, nil
}

// Store adds a new key in the backend
func (rd *Random) Store(key, secret string) error {
	// check if secret is in correct format
	_, err := base32.StdEncoding.DecodeString(strings.ToUpper(secret))
	if err != nil {
		return errors.New("secret is not is correct format: " + secret)
	}
	newKey := gotp.Secret{
		Key:   key,
		Value: secret,
	}
	rd.keys = append(rd.keys, newKey)
	return nil
}
