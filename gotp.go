package gotp

import (
	"time"

	"github.com/jtbonhomme/gotp/backend"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// DefaultTimeIntervalSeedTimeIntervalSeed is the default time interval seed used to compute the HOTP.
const DefaultTimeIntervalSeed uint = 30

type GOTP struct {
	backend          backend.Backend
	timeIntervalSeed uint
}

// New instanciates a GoTP object with a secured backend
func New(bkd backend.Backend) *GOTP {
	return &GOTP{
		timeIntervalSeed: DefaultTimeIntervalSeed,
		backend:          bkd,
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
	return gotp.backend.List()
}

// Store creates a new key/value pair in the secured key ring
func (gotp *GOTP) Store(key, value string) error {
	return gotp.backend.Store(key, value)
}

// Remove deletes a key in the secured key ring
func (gotp *GOTP) Remove(key string) error {
	return gotp.backend.Remove(key)
}

// Get retrieves a key in the secured key ring
func (gotp *GOTP) Get(key string) (string, error) {
	secret, err := gotp.backend.Read(key)
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
