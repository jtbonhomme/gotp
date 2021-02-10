package gotp

import (
	"github.com/jtbonhomme/gotp/backend/secure"
	"time"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// DefaultTimeIntervalSeedTimeIntervalSeed is the default time interval seed used to compute the HOTP.
const DefaultTimeIntervalSeed int64 = 30

type GOTP struct {
	secRing *secure.Secure
}

func New(secring *secure.Secure) *GOTP {
	// create the key ring
	return &GOTP{
		secRing: secring,
	}
}

// List retrieves all existing keys in the secured key ring
func (gotp *GOTP) List() ([]string, error) {
	return gotp.secRing.List()
}

// Store creates a new key/value pair in the secured key ring
func (gotp *GOTP) Store(key, value string) error {
	return gotp.secRing.Store(key, value)
}

// Remove deletes a key in the secured key ring
func (gotp *GOTP) Remove(key string) error {
	return gotp.secRing.Remove(key)
}

// Get retrieves a key in the secured key ring
func (gotp *GOTP) Get(key string) (string, error) {
	secret, err := gotp.secRing.Read(key)
	if err != nil {
		return "", err
	}
	return totp.GenerateCodeCustom(string(secret), time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
}
