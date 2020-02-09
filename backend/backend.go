package backend

import "github.com/jtbonhomme/gotp"

// Backend is an interface that describes how a backend shall store secret keys.
type Backend interface {
	// List lists all keys stored in the backend
	List() (*[]gotp.TOTP, error)
	// Store adds a new key in the backend
	Store(string, string) error
	// Remove deletes a new key in the backend
	Remove(string) error
	// Read retrieves a key stored in the backend
	Read(string) (*gotp.TOTP, error)
}
