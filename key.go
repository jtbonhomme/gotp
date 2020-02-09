package gotp

// Secret describes a secret pair {key, value} to be used to generate TOTP.
// This struct is only used for store or update purpose, it can never be fetched.
type Secret struct {
	// Key is the unique resource identifier
	Key string
	// Value is the secret key value for this Key
	Value string
}

// TOTP describes a Time-based One-Time Password
type TOTP struct {
	// Key is the unique resource identifier
	Key string
	// Code is the time-based one-time password for this Key
	Code string
}
