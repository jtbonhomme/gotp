package gotp

import (
	"github.com/jtbonhomme/gotp/backend/secure"
	"time"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// DefaultTimeIntervalSeedTimeIntervalSeed is the default time interval seed used to compute the HOTP.
const DefaultTimeIntervalSeed uint = 30

type GOTP struct {
	secRing *secure.Secure
	timeIntervalSeed uint
}

// New instanciates a GoTP object with a secured backend
func New(secring *secure.Secure) *GOTP {
	return &GOTP{
		timeIntervalSeed: DefaultTimeIntervalSeed,
		secRing: secring,
	}
}

// WithTimeIntervalSeed configures a specific timeIntervalSeed
func (gotp *GOTP) WithTimeIntervalSeed(interval uint) *GOTP {
	g := &GOTP{}
	*g = *gotp

	g.timeIntervalSeed = interval
	return g
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
		Period:    gotp.timeIntervalSeed,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
}
