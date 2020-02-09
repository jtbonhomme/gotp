package gotp

// Key describes a secret key to be used to generate TOTP.
// This struct is only used for store or update purpose, it can never be fetched.
type Key struct {
	// URI is the unique resource identifier
	URI string
	// Value is the secret key value for this URI
	Value string
}

// TOTP describes a Time-based One-Time Password
type TOTP struct {
	// URI is the unique resource identifier
	URI string
	// Code is the time-based one-time password for this URI
	Code string
}
