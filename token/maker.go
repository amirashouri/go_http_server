package token

import "time"

// Maker is an interface to manage tokens
type Maker interface {
	// CreateToken creates a new token
	CreateToken(username string, role string, duration time.Duration) (string, *Payload, error)

	// VerifyToken check if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
