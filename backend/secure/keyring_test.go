package secure

import (
	"github.com/99designs/keyring"
)

// Keyring provides the uniform interface over the underlying backends
type MockKeyring struct{}

// Returns an Item matching the key or ErrKeyNotFound
func (m MockKeyring) Get(key string) (keyring.Item, error) {
	item := keyring.Item{
		Data: []byte("result"),
	}
	return item, nil
}

// Returns the non-secret parts of an Item
func (m MockKeyring) GetMetadata(key string) (keyring.Metadata, error) {
	return keyring.Metadata{}, nil
}

// Stores an Item on the keyring
func (m MockKeyring) Set(item keyring.Item) error {
	return nil
}

// Removes the item with matching key
func (m MockKeyring) Remove(key string) error {
	return nil
}

// Provides a slice of all keys stored on the keyring
func (m MockKeyring) Keys() ([]string, error) {
	return []string{}, nil
}
