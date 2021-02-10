package secure

import (
	"encoding/base32"
	"errors"
	"fmt"
	"strings"

	"github.com/99designs/keyring"
)

// Secure is a secure backend.
// It uses system keychain to store keys and secrets.
// macOS/OSX Keychain, Windows credential store, or Linux Pass.
type Secure struct {
	keyring keyring.Keyring
}

// New creates a new secure backend instance.
// It takes as argument the name of the keyring to be created or opened.
func New(kr string) *Secure {
	kc := keyring.Config{
		ServiceName:     kr,
		AllowedBackends: []keyring.BackendType{"keychain", "pass"},
		PassPrefix:      kr,
	}
	ring, err := keyring.Open(kc)
	if err != nil {
		_ = fmt.Errorf("error while initialize key ring: %s", err.Error())
		return nil
	}

	sec := Secure{
		keyring: ring,
	}

	return &sec
}

// List gets all keys stored in the keychain backend
func (sec *Secure) List() ([]string, error) {
	return sec.keyring.Keys()
}

// Store adds a new key in the keychain backend
func (sec *Secure) Store(key, secret string) error {
	// check if secret is in correct format
	_, err := base32.StdEncoding.DecodeString(strings.ToUpper(secret))
	if err != nil {
		return errors.New("secret is not is correct format: " + secret)
	}

	return sec.keyring.Set(keyring.Item{
		Key:  key,
		Data: []byte(secret),
	})
}

// Remove deletes a key from the keychain backend
func (sec *Secure) Remove(key string) error {
	return sec.keyring.Remove(key)
}

// Read retrieves a key stored in the backend
func (sec *Secure) Read(key string) ([]byte, error) {
	item, err := sec.keyring.Get(key)
	return item.Data, err
}
