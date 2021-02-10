package secure

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSecure(t *testing.T) {
	bck := New("workspace")
	serviceType := fmt.Sprintf("%T", bck)
	assert.Equal(t, "*secure.Secure", serviceType)
}

func TestListOK(t *testing.T) {
	s := Secure{
		keyring: MockKeyring{},
	}
	l, err := s.List()
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, l)
}

func TestStoreError(t *testing.T) {
	s := Secure{
		keyring: MockKeyring{},
	}
	err := s.Store("key", "value")
	assert.NotEqual(t, nil, err)
	assert.Equal(t, "secret is not is correct format: value", err.Error())
}

func TestStoreOK(t *testing.T) {
	s := Secure{
		keyring: MockKeyring{},
	}
	err := s.Store("key", "KREESUZAJFJSAQJAKRCVGVA=") // "THIS IS A TEST"
	assert.Equal(t, nil, err)
}

func TestRemoveOK(t *testing.T) {
	s := Secure{
		keyring: MockKeyring{},
	}
	err := s.Remove("key")
	assert.Equal(t, nil, err)
}

func TestReadOK(t *testing.T) {
	s := Secure{
		keyring: MockKeyring{},
	}
	b, err := s.Read("key")
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, b)
	assert.Equal(t, "result", string(b))
}
