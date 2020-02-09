package gotp

// TOTP describes a Time-based One-Time Password
type TOTP struct {
	// Key is the unique resource identifier
	Key string
	// Code is the time-based one-time password for this Key
	Code string
}
